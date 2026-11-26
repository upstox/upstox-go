//go:build sdk

// Package sdk is a live sanity suite that calls the real Upstox API through the Go SDK.
package sdk_tests

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/antihax/optional"
	"github.com/upstox/upstox-go/upstox_client"
	"github.com/upstox/upstox-go/upstox_client/api"
	"github.com/upstox/upstox-go/upstox_client/models"
)

const (
	// Vodafone Idea: low-priced, liquid NSE equity used for order and quote tests.
	ideaKey = "NSE_EQ|INE669E01016"
	// Reliance Industries, used for fundamentals and margin tests.
	relianceISIN = "INE002A01018"
	relianceKey  = "NSE_EQ|INE002A01018"
	niftyKey     = "NSE_INDEX|Nifty 50"
	// Tag on every order this suite places, so bulk cancels only touch our orders.
	testTag = "sdk-go-test"
	// Syntactically valid IDs that do not exist, for not-found probes.
	missingOrderID = "231221011081579"
	missingGTTID   = "GTT-C250303008840"
)

var ist = time.FixedZone("IST", 5*3600+1800)

// client returns a live client, skipping the test when UPSTOX_ACCESS_TOKEN is unset.
func client(t *testing.T) *upstox_client.APIClient {
	t.Helper()
	token := Token()
	if token == "" {
		t.Skip("no access token: set UPSTOX_ACCESS_TOKEN or data_token.AccessToken")
	}
	cfg := upstox_client.NewConfiguration()
	cfg.AccessToken = token
	return upstox_client.NewAPIClient(cfg)
}

// noRedirectClient is client(t) that surfaces 3xx responses instead of following them.
func noRedirectClient(t *testing.T) *upstox_client.APIClient {
	t.Helper()
	token := Token()
	if token == "" {
		t.Skip("no access token: set UPSTOX_ACCESS_TOKEN or data_token.AccessToken")
	}
	cfg := upstox_client.NewConfiguration()
	cfg.AccessToken = token
	cfg.HTTPClient = &http.Client{
		Timeout: 30 * time.Second,
		CheckRedirect: func(*http.Request, []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	return upstox_client.NewAPIClient(cfg)
}

// sandboxClient returns a sandbox client, skipping when UPSTOX_SANDBOX_ACCESS_TOKEN is unset.
func sandboxClient(t *testing.T) *upstox_client.APIClient {
	t.Helper()
	token := SandboxToken()
	if token == "" {
		t.Skip("no sandbox token: set UPSTOX_SANDBOX_ACCESS_TOKEN or data_token.SandboxAccessToken")
	}
	cfg := upstox_client.NewSandboxConfiguration()
	cfg.AccessToken = token
	return upstox_client.NewAPIClient(cfg)
}

// ctx returns a per-test context with a 30s timeout.
func ctx(t *testing.T) context.Context {
	t.Helper()
	c, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	t.Cleanup(cancel)
	return c
}

// requireOrderTests skips unless UPSTOX_RUN_ORDER_TESTS=1.
func requireOrderTests(t *testing.T) {
	t.Helper()
	requireFlag(t, "UPSTOX_RUN_ORDER_TESTS")
}

// requireFlag skips unless the named env var is "1".
func requireFlag(t *testing.T, name string) {
	t.Helper()
	if os.Getenv(name) != "1" {
		t.Skipf("%s=1 not set", name)
	}
}

// requireEnv skips unless every named env var is set, and returns their values in order.
func requireEnv(t *testing.T, names ...string) []string {
	t.Helper()
	vals := make([]string, len(names))
	for i, n := range names {
		vals[i] = os.Getenv(n)
		if vals[i] == "" {
			t.Skipf("%s not set", n)
		}
	}
	return vals
}

// check fails the test on any error or a non-"success" status.
func check(t *testing.T, resp interface{}, err error) {
	t.Helper()
	if err != nil {
		skipIfNotEntitled(t, err)
		t.Fatalf("API error: %v %s", err, errorBody(err))
	}
	if s, ok := status(resp); ok && s != "success" {
		t.Fatalf("status = %q, want success", s)
	}
}

// checkAllow is check, but also passes when the API returns one of the given business error codes.
// It reports whether the call succeeded, so callers can skip follow-up assertions.
func checkAllow(t *testing.T, resp interface{}, err error, allowed ...string) bool {
	t.Helper()
	if err != nil {
		codes := errorCodes(err)
		for _, c := range codes {
			for _, a := range allowed {
				if c == a {
					t.Logf("accepted business error %s: %s", c, errorBody(err))
					return false
				}
			}
		}
		skipIfNotEntitled(t, err)
		t.Fatalf("API error: %v codes=%v %s", err, codes, errorBody(err))
	}
	check(t, resp, nil)
	return true
}

// checkRejected passes when the API answered a deliberately invalid call with a structured 4xx error.
func checkRejected(t *testing.T, httpResp *http.Response, err error) {
	t.Helper()
	if err == nil {
		t.Fatalf("expected the API to reject the request, got %v", statusCode(httpResp))
	}
	if httpResp == nil || httpResp.StatusCode < 400 || httpResp.StatusCode >= 500 {
		t.Fatalf("expected a 4xx error, got %v: %v %s", statusCode(httpResp), err, errorBody(err))
	}
	t.Logf("rejected as expected (%d) codes=%v", httpResp.StatusCode, errorCodes(err))
}

// checkRedirect passes when a raw endpoint answered with a redirect or success.
// skipIfNotEntitled skips tests for APIs the account's plan does not include.
func skipIfNotEntitled(t *testing.T, err error) {
	t.Helper()
	for _, c := range errorCodes(err) {
		if c == "UDAPI1149" {
			t.Skipf("account lacks the required plan (UDAPI1149): %s", errorBody(err))
		}
	}
}

// checkDiscontinued expects a retired endpoint to answer 410 Gone.
func checkDiscontinued(t *testing.T, httpResp *http.Response, err error) {
	t.Helper()
	if httpResp == nil || httpResp.StatusCode != http.StatusGone {
		t.Fatalf("want 410 Gone for a discontinued endpoint, got %v: %v %s", statusCode(httpResp), err, errorBody(err))
	}
}

func checkRedirect(t *testing.T, httpResp *http.Response, err error) {
	t.Helper()
	if httpResp == nil {
		t.Fatalf("no response: %v", err)
	}
	if httpResp.StatusCode >= 400 {
		t.Fatalf("HTTP %d: %v %s", httpResp.StatusCode, err, errorBody(err))
	}
	t.Logf("HTTP %d location=%q", httpResp.StatusCode, redactQuery(httpResp.Header.Get("Location")))
}

func statusCode(r *http.Response) interface{} {
	if r == nil {
		return "no response"
	}
	return r.StatusCode
}

// redactQuery drops the query string, which can carry one-time auth codes.
func redactQuery(u string) string {
	if i := strings.IndexByte(u, '?'); i >= 0 {
		return u[:i] + "?..."
	}
	return u
}

// status reads the Status field by reflection, since generated models have no getters.
func status(resp interface{}) (string, bool) {
	v := reflect.ValueOf(resp)
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		if v.IsNil() {
			return "", false
		}
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return "", false
	}
	f := v.FieldByName("Status")
	if !f.IsValid() || f.Kind() != reflect.String {
		return "", false
	}
	return f.String(), true
}

// errorBody returns the raw body of a GenericSwaggerError, or "".
func errorBody(err error) string {
	var apiErr upstox_client.GenericSwaggerError
	if errors.As(err, &apiErr) {
		return string(apiErr.Body())
	}
	return ""
}

// errorCodes extracts error codes from an Upstox error body; it accepts both errorCode and error_code.
func errorCodes(err error) []string {
	var apiErr upstox_client.GenericSwaggerError
	if !errors.As(err, &apiErr) {
		return nil
	}
	var body struct {
		Errors []struct {
			ErrorCode  string `json:"errorCode"`
			ErrorCode2 string `json:"error_code"`
		} `json:"errors"`
	}
	if json.Unmarshal(apiErr.Body(), &body) != nil {
		return nil
	}
	var codes []string
	for _, e := range body.Errors {
		if e.ErrorCode != "" {
			codes = append(codes, e.ErrorCode)
		}
		if e.ErrorCode2 != "" {
			codes = append(codes, e.ErrorCode2)
		}
	}
	return codes
}

// today returns the current IST date as yyyy-MM-dd.
func today() string {
	return time.Now().In(ist).Format("2006-01-02")
}

// daysAgo returns the IST date n days back as yyyy-MM-dd.
func daysAgo(n int) string {
	return time.Now().In(ist).AddDate(0, 0, -n).Format("2006-01-02")
}

// lastWeekday returns today, or the preceding Friday on a weekend, as yyyy-MM-dd.
func lastWeekday() string {
	d := time.Now().In(ist)
	for d.Weekday() == time.Saturday || d.Weekday() == time.Sunday {
		d = d.AddDate(0, 0, -1)
	}
	return d.Format("2006-01-02")
}

// previousFinancialYear returns the last completed Indian FY in Upstox's "2425" form.
func previousFinancialYear() string {
	now := time.Now().In(ist)
	start := now.Year() - 1
	if now.Month() < time.April {
		start--
	}
	return fmt.Sprintf("%02d%02d", start%100, (start+1)%100)
}

// ltp fetches the last traded price for an instrument key.
func ltp(t *testing.T, c *upstox_client.APIClient, key string) float64 {
	t.Helper()
	resp, _, err := api.NewMarketQuoteV3Api(c).GetLtp(ctx(t), &api.MarketQuoteV3ApiGetLtpOpts{InstrumentKey: optional.NewString(key)})
	check(t, resp, err)
	// Response keys are exchange:symbol, not the instrument key
	for _, q := range resp.Data {
		if q.LastPrice > 0 {
			return q.LastPrice
		}
	}
	t.Fatalf("no LTP for %s", key)
	return 0
}

// farLimitPrice returns a price frac below ltp, rounded down to a 0.05 tick.
func farLimitPrice(ltp, frac float64) float64 {
	p := math.Floor(ltp*(1-frac)*20) / 20
	if p < 0.05 {
		p = 0.05
	}
	return math.Round(p*100) / 100
}

// marketOpen reports whether NSE is in its normal trading session; AMO is used otherwise.
func marketOpen(t *testing.T, c *upstox_client.APIClient) bool {
	t.Helper()
	resp, _, err := api.NewMarketHolidaysAndTimingsApi(c).GetMarketStatus(ctx(t), "NSE")
	check(t, resp, err)
	return resp.Data != nil && resp.Data.Status == "NORMAL_OPEN"
}

// niftyOption finds the at-the-money NIFTY call for the current month via instrument search.
func niftyOption(t *testing.T, c *upstox_client.APIClient) models.InstrumentData {
	t.Helper()
	resp, _, err := api.NewInstrumentsApi(c).SearchInstrument(ctx(t), "Nifty 50", &api.InstrumentsApiSearchInstrumentOpts{
		Exchanges:       optional.NewString("NSE"),
		Segments:        optional.NewString("FO"),
		InstrumentTypes: optional.NewString("CE"),
		Expiry:          optional.NewString("current_month"),
		AtmOffset:       optional.NewInt32(0),
		Records:         optional.NewInt32(1),
	})
	check(t, resp, err)
	for _, d := range resp.Data {
		if d.InstrumentKey != "" && d.Expiry != "" {
			return d
		}
	}
	t.Skip("no NIFTY option found via instrument search")
	return models.InstrumentData{}
}

// niftyExpiry returns the yyyy-MM-dd expiry of the current-month NIFTY option.
func niftyExpiry(t *testing.T, c *upstox_client.APIClient) string {
	t.Helper()
	e := niftyOption(t, c).Expiry
	if len(e) > 10 {
		e = e[:10]
	}
	return e
}

// cleanupCtx returns a fresh context for t.Cleanup funcs, which run after ctx(t) is cancelled.
func cleanupCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 30*time.Second)
}

// algoName is sent as X-Algo-Name on order writes.
func algoName() optional.String {
	if v := os.Getenv("UPSTOX_ALGO_NAME"); v != "" {
		return optional.NewString(v)
	}
	return optional.NewString("sdk-go-test")
}
