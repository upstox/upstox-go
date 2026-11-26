//go:build sdk

package sdk

import (
	"testing"

	"github.com/antihax/optional"
	"github.com/upstox/upstox-go/upstox/api"
)

func TestGetNewsByInstrument(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewNewsApi(c).GetNews(ctx(t), "instrument_keys", &api.NewsApiGetNewsOpts{
		InstrumentKeys: optional.NewString(ideaKey),
	})
	check(t, resp, err)
}

func TestGetNewsPaged(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewNewsApi(c).GetNews(ctx(t), "instrument_keys", &api.NewsApiGetNewsOpts{
		InstrumentKeys: optional.NewString("NSE_EQ|INE040A01034"),
		PageNumber:     optional.NewInt32(1),
		PageSize:       optional.NewInt32(5),
	})
	check(t, resp, err)
}

func TestGetNewsPositions(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewNewsApi(c).GetNews(ctx(t), "positions", nil)
	check(t, resp, err)
}

func TestGetNewsHoldings(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewNewsApi(c).GetNews(ctx(t), "holdings", nil)
	check(t, resp, err)
}
