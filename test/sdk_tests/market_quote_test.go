//go:build sdk

package sdk_tests

import (
	"testing"

	"github.com/antihax/optional"
	"github.com/upstox/upstox-go/upstox_client/api"
)

func TestGetFullMarketQuote(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewMarketQuoteApi(c).GetFullMarketQuote(ctx(t), &api.MarketQuoteApiGetFullMarketQuoteOpts{
		InstrumentKey: optional.NewString(ideaKey),
	})
	check(t, resp, err)
	if len(resp.Data) == 0 {
		t.Fatal("no quote returned")
	}
}

func TestGetFullMarketQuoteMultiple(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewMarketQuoteApi(c).GetFullMarketQuote(ctx(t), &api.MarketQuoteApiGetFullMarketQuoteOpts{
		InstrumentKey: optional.NewString(ideaKey + "," + relianceKey),
	})
	check(t, resp, err)
	if len(resp.Data) != 2 {
		t.Fatalf("got %d quotes, want 2", len(resp.Data))
	}
}

func TestLtp(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewMarketQuoteApi(c).Ltp(ctx(t), &api.MarketQuoteApiLtpOpts{
		InstrumentKey: optional.NewString(ideaKey),
	})
	check(t, resp, err)
}

func TestGetMarketQuoteOHLC(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewMarketQuoteApi(c).GetMarketQuoteOHLC(ctx(t), "1d", &api.MarketQuoteApiGetMarketQuoteOHLCOpts{
		InstrumentKey: optional.NewString(ideaKey),
	})
	check(t, resp, err)
}
