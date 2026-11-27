//go:build sdk

package sdk_tests

import (
	"testing"
	"time"

	"github.com/upstox/upstox-go/upstox_client"
	"github.com/upstox/upstox-go/upstox_client/api"
)

// pastExpiries returns NIFTY expiries before today, newest first.
func pastExpiries(t *testing.T, c *upstox_client.APIClient) []string {
	t.Helper()
	resp, _, err := api.NewExpiredInstrumentApi(c).GetExpiries(ctx(t), niftyKey)
	check(t, resp, err)
	var out []string
	for i := len(resp.Data) - 1; i >= 0; i-- {
		if resp.Data[i] < today() {
			out = append(out, resp.Data[i])
		}
	}
	if len(out) == 0 {
		t.Skip("no past NIFTY expiries returned")
	}
	return out
}

func TestGetExpiries(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewExpiredInstrumentApi(c).GetExpiries(ctx(t), niftyKey)
	check(t, resp, err)
	if len(resp.Data) == 0 {
		t.Fatal("no expiries returned")
	}
}

func TestGetExpiredOptionContracts(t *testing.T) {
	c := client(t)
	expiry := pastExpiries(t, c)[0]
	resp, _, err := api.NewExpiredInstrumentApi(c).GetExpiredOptionContracts(ctx(t), niftyKey, expiry)
	check(t, resp, err)
}

func TestGetExpiredFutureContracts(t *testing.T) {
	c := client(t)
	expiry := pastExpiries(t, c)[0]
	resp, _, err := api.NewExpiredInstrumentApi(c).GetExpiredFutureContracts(ctx(t), niftyKey, expiry)
	check(t, resp, err)
}

func TestGetExpiredHistoricalCandleData(t *testing.T) {
	c := client(t)
	e := api.NewExpiredInstrumentApi(c)
	// Only monthly expiries carry futures, so walk back until one has contracts
	for i, expiry := range pastExpiries(t, c) {
		if i >= 6 {
			break
		}
		fut, _, err := e.GetExpiredFutureContracts(ctx(t), niftyKey, expiry)
		check(t, fut, err)
		if len(fut.Data) == 0 || fut.Data[0].InstrumentKey == "" {
			continue
		}
		to, _ := time.Parse("2006-01-02", expiry)
		from := to.AddDate(0, 0, -7).Format("2006-01-02")
		resp, _, err := e.GetExpiredHistoricalCandleData(ctx(t), fut.Data[0].InstrumentKey, "day", expiry, from)
		check(t, resp, err)
		return
	}
	t.Skip("no expired NIFTY future found in recent expiries")
}
