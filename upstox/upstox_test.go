package upstox_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/antihax/optional"
	"github.com/upstox/upstox-go/upstox"
	"github.com/upstox/upstox-go/upstox/api"
	"github.com/upstox/upstox-go/upstox/models"
)

// recorder captures outgoing requests and answers each with an empty success body.
type recorder struct{ reqs []*http.Request }

func (r *recorder) RoundTrip(req *http.Request) (*http.Response, error) {
	r.reqs = append(r.reqs, req)
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(strings.NewReader(`{"status":"success"}`)),
		Request:    req,
	}, nil
}

func newTestClient(cfg *upstox.Configuration) (*upstox.APIClient, *recorder) {
	rec := &recorder{}
	cfg.HTTPClient = &http.Client{Transport: rec}
	cfg.AccessToken = "test-token"
	return upstox.NewAPIClient(cfg), rec
}

func TestOrderRouting(t *testing.T) {
	client, rec := newTestClient(upstox.NewConfiguration())
	ctx := context.Background()

	api.NewOrderApi(client).PlaceOrder(ctx, models.PlaceOrderRequest{}, nil)
	api.NewOrderV3Api(client).PlaceOrderV3(ctx, models.PlaceOrderV3Request{}, nil)
	api.NewOrderV3Api(client).CancelOrderV3(ctx, "1", nil)
	api.NewUSStocksOrderApi(client).PlaceUsOrder(ctx, models.UsPlaceOrderRequest{})
	api.NewUSStocksOrderApi(client).CancelUsOrder(ctx, "1")
	api.NewOrderApi(client).GetOrderBook(ctx)

	want := []string{
		"api-hft.upstox.com/v2/order/place",
		"api-hft.upstox.com/v3/order/place",
		"api-hft.upstox.com/v3/order/cancel",
		"api.upstox.com/v2/us/order/place",
		"api.upstox.com/v2/us/order/cancel",
		"api.upstox.com/v2/order/retrieve-all",
	}
	if len(rec.reqs) != len(want) {
		t.Fatalf("got %d requests, want %d", len(rec.reqs), len(want))
	}
	for i, req := range rec.reqs {
		if got := req.URL.Host + req.URL.Path; got != want[i] {
			t.Errorf("request %d: got %s, want %s", i, got, want[i])
		}
	}
}

func TestCustomBasePathSkipsOrderHost(t *testing.T) {
	cfg := upstox.NewConfiguration()
	cfg.BasePath = "http://localhost:8080"
	client, rec := newTestClient(cfg)
	api.NewOrderApi(client).PlaceOrder(context.Background(), models.PlaceOrderRequest{}, nil)
	if got := rec.reqs[0].URL.Host; got != "localhost:8080" {
		t.Errorf("got host %s, want localhost:8080", got)
	}
}

func TestSandbox(t *testing.T) {
	client, rec := newTestClient(upstox.NewSandboxConfiguration())
	ctx := context.Background()

	api.NewOrderV3Api(client).PlaceOrderV3(ctx, models.PlaceOrderV3Request{}, nil)
	if len(rec.reqs) != 1 || rec.reqs[0].URL.Host != "api-sandbox.upstox.com" {
		t.Fatalf("sandbox order not sent to sandbox host: %+v", rec.reqs)
	}

	_, _, err := api.NewUserApi(client).GetProfile(ctx)
	if err == nil || !strings.Contains(err.Error(), "not available in sandbox mode") {
		t.Errorf("got err %v, want sandbox restriction", err)
	}
	if len(rec.reqs) != 1 {
		t.Errorf("non-sandbox call was sent")
	}
}

func TestHeadersAndParams(t *testing.T) {
	client, rec := newTestClient(upstox.NewConfiguration())
	ctx := context.Background()

	api.NewUSMarketHolidaysAndTimingsApi(client).GetUsHoliday(ctx, "2026-12-25")
	api.NewUSStocksOrderApi(client).GetUsOrderBook(ctx, &api.USStocksOrderApiGetUsOrderBookOpts{Tag: optional.NewString("sdk-test")})

	req := rec.reqs[0]
	if req.URL.Path != "/v2/us/market/holidays/2026-12-25" {
		t.Errorf("got path %s", req.URL.Path)
	}
	if got := req.Header.Get("Authorization"); got != "Bearer test-token" {
		t.Errorf("got Authorization %q", got)
	}
	if got := req.Header.Get("X-Upstox-SDK-Language"); got != "go" {
		t.Errorf("got SDK language %q", got)
	}
	if got := req.Header.Get("X-Upstox-SDK-Version"); got != upstox.SDKVersion {
		t.Errorf("got SDK version %q", got)
	}
	if got := rec.reqs[1].URL.Query().Get("tag"); got != "sdk-test" {
		t.Errorf("got tag %q", got)
	}
}

func TestUsOrderBodyIsSnakeCase(t *testing.T) {
	client, rec := newTestClient(upstox.NewConfiguration())
	body := models.UsPlaceOrderRequest{InstrumentKey: "US_EQ|COUR", TransactionType: "BUY", OrderType: "LIMIT"}
	api.NewUSStocksOrderApi(client).PlaceUsOrder(context.Background(), body)

	raw, _ := io.ReadAll(rec.reqs[0].Body)
	for _, key := range []string{`"instrument_key"`, `"transaction_type"`, `"order_type"`} {
		if !strings.Contains(string(raw), key) {
			t.Errorf("body %s missing %s", raw, key)
		}
	}
}

func TestAlgoNameHeader(t *testing.T) {
	client, rec := newTestClient(upstox.NewConfiguration())
	ctx := context.Background()
	api.NewOrderV3Api(client).PlaceOrderV3(ctx, models.PlaceOrderV3Request{}, &api.OrderV3ApiPlaceOrderV3Opts{XAlgoName: optional.NewString("my-algo")})
	api.NewOrderApi(client).CancelOrder(ctx, "1", nil)

	if got := rec.reqs[0].Header.Get("X-Algo-Name"); got != "my-algo" {
		t.Errorf("got X-Algo-Name %q, want my-algo", got)
	}
	if got := rec.reqs[1].Header.Get("X-Algo-Name"); got != "" {
		t.Errorf("unset algo name sent as %q", got)
	}
}

func TestMarketProtectionOmittedWhenUnset(t *testing.T) {
	client, rec := newTestClient(upstox.NewConfiguration())
	ctx := context.Background()
	api.NewOrderV3Api(client).PlaceOrderV3(ctx, models.PlaceOrderV3Request{Quantity: 1}, nil)
	api.NewOrderV3Api(client).PlaceOrderV3(ctx, models.PlaceOrderV3Request{Quantity: 1, MarketProtection: 2.5}, nil)

	first, _ := io.ReadAll(rec.reqs[0].Body)
	second, _ := io.ReadAll(rec.reqs[1].Body)
	if strings.Contains(string(first), "market_protection") {
		t.Errorf("unset market_protection sent: %s", first)
	}
	if !strings.Contains(string(second), `"market_protection":2.5`) {
		t.Errorf("market_protection missing: %s", second)
	}
}

func TestPricesKeepFloat64Precision(t *testing.T) {
	client, rec := newTestClient(upstox.NewConfiguration())
	api.NewOrderV3Api(client).PlaceOrderV3(context.Background(), models.PlaceOrderV3Request{Price: 2345.65}, nil)
	raw, _ := io.ReadAll(rec.reqs[0].Body)
	if !strings.Contains(string(raw), `"price":2345.65`) {
		t.Errorf("price lost precision: %s", raw)
	}
}
