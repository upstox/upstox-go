//go:build sdk

package sdk

import (
	"testing"

	"github.com/antihax/optional"
	"github.com/upstox/upstox-go/upstox/api"
)

func TestSearchInstrument(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewInstrumentsApi(c).SearchInstrument(ctx(t), "Nifty 50", nil)
	check(t, resp, err)
	if len(resp.Data) == 0 {
		t.Fatal("no instruments returned")
	}
}

func TestSearchInstrumentByExchange(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewInstrumentsApi(c).SearchInstrument(ctx(t), "Reliance", &api.InstrumentsApiSearchInstrumentOpts{
		Exchanges: optional.NewString("NSE"),
	})
	check(t, resp, err)
}

func TestSearchInstrumentBySegment(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewInstrumentsApi(c).SearchInstrument(ctx(t), "TCS", &api.InstrumentsApiSearchInstrumentOpts{
		Segments: optional.NewString("EQ"),
	})
	check(t, resp, err)
}

func TestSearchInstrumentByType(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewInstrumentsApi(c).SearchInstrument(ctx(t), "Nifty", &api.InstrumentsApiSearchInstrumentOpts{
		InstrumentTypes: optional.NewString("INDEX"),
	})
	check(t, resp, err)
}

func TestSearchInstrumentPagination(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewInstrumentsApi(c).SearchInstrument(ctx(t), "HDFC", &api.InstrumentsApiSearchInstrumentOpts{
		PageNumber: optional.NewInt32(1),
		Records:    optional.NewInt32(5),
	})
	check(t, resp, err)
	if len(resp.Data) > 5 {
		t.Fatalf("got %d records, want at most 5", len(resp.Data))
	}
}

func TestSearchInstrumentATMOption(t *testing.T) {
	c := client(t)
	opt := niftyOption(t, c)
	if opt.InstrumentType != "CE" {
		t.Fatalf("instrument type = %q, want CE", opt.InstrumentType)
	}
}
