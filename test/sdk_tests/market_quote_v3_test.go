//go:build sdk

package sdk_tests

import (
	"testing"

	"github.com/antihax/optional"
	"github.com/upstox/upstox-go/upstox_client/api"
)

func TestGetLtpV3(t *testing.T) {
	c := client(t)
	if p := ltp(t, c, ideaKey); p <= 0 {
		t.Fatalf("ltp = %v", p)
	}
}

func TestGetLtpV3Multiple(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewMarketQuoteV3Api(c).GetLtp(ctx(t), &api.MarketQuoteV3ApiGetLtpOpts{
		InstrumentKey: optional.NewString(ideaKey + "," + niftyKey),
	})
	check(t, resp, err)
	if len(resp.Data) != 2 {
		t.Fatalf("got %d prices, want 2", len(resp.Data))
	}
}

func TestGetFullMarketQuoteV3(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewMarketQuoteV3Api(c).GetFullMarketQuoteV3(ctx(t), &api.MarketQuoteV3ApiGetFullMarketQuoteV3Opts{
		InstrumentKey: optional.NewString(ideaKey),
	})
	check(t, resp, err)
}

func TestGetMarketQuoteOHLCV3(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewMarketQuoteV3Api(c).GetMarketQuoteOHLCV3(ctx(t), "1d", &api.MarketQuoteV3ApiGetMarketQuoteOHLCV3Opts{
		InstrumentKey: optional.NewString(ideaKey),
	})
	check(t, resp, err)
}

func TestGetMarketQuoteOHLCV3Intraday(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewMarketQuoteV3Api(c).GetMarketQuoteOHLCV3(ctx(t), "I1", &api.MarketQuoteV3ApiGetMarketQuoteOHLCV3Opts{
		InstrumentKey: optional.NewString(ideaKey),
	})
	check(t, resp, err)
}

func TestGetMarketQuoteOptionGreek(t *testing.T) {
	c := client(t)
	opt := niftyOption(t, c)
	resp, _, err := api.NewMarketQuoteV3Api(c).GetMarketQuoteOptionGreek(ctx(t), &api.MarketQuoteV3ApiGetMarketQuoteOptionGreekOpts{
		InstrumentKey: optional.NewString(opt.InstrumentKey),
	})
	check(t, resp, err)
}
