//go:build sdk

package sdk_tests

import (
	"testing"

	"github.com/antihax/optional"
	"github.com/upstox/upstox-go/upstox_client/api"
)

func TestGetHolidays(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewMarketHolidaysAndTimingsApi(c).GetHolidays(ctx(t))
	check(t, resp, err)
	if len(resp.Data) == 0 {
		t.Fatal("no holidays returned")
	}
}

func TestGetHoliday(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewMarketHolidaysAndTimingsApi(c).GetHoliday(ctx(t), "2024-01-22")
	check(t, resp, err)
}

func TestGetExchangeTimings(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewMarketHolidaysAndTimingsApi(c).GetExchangeTimings(ctx(t), today())
	check(t, resp, err)
}

func TestGetMarketStatus(t *testing.T) {
	c := client(t)
	for _, ex := range []string{"NSE", "BSE", "MCX"} {
		resp, _, err := api.NewMarketHolidaysAndTimingsApi(c).GetMarketStatus(ctx(t), ex)
		check(t, resp, err)
		if resp.Data == nil || resp.Data.Status == "" {
			t.Fatalf("%s: empty market status", ex)
		}
	}
}

func TestGetOiData(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewMarketHolidaysAndTimingsApi(c).GetOiData(ctx(t), niftyKey, niftyExpiry(t, c), lastWeekday())
	check(t, resp, err)
}

func TestGetChangeOiData(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewMarketHolidaysAndTimingsApi(c).GetChangeOiData(ctx(t), niftyKey, niftyExpiry(t, c), lastWeekday(), 5)
	check(t, resp, err)
}

func TestGetPcrData(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewMarketHolidaysAndTimingsApi(c).GetPcrData(ctx(t), niftyKey, niftyExpiry(t, c), lastWeekday(), 30)
	check(t, resp, err)
}

func TestGetMaxPainData(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewMarketHolidaysAndTimingsApi(c).GetMaxPainData(ctx(t), niftyKey, niftyExpiry(t, c), lastWeekday(), 30)
	check(t, resp, err)
}

func TestGetFiiData(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewMarketHolidaysAndTimingsApi(c).GetFiiData(ctx(t), "NSE_FO|INDEX_FUTURES", "1D", nil)
	check(t, resp, err)
}

func TestGetFiiDataFrom(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewMarketHolidaysAndTimingsApi(c).GetFiiData(ctx(t), "NSE_EQ|CASH", "1D", &api.MarketHolidaysAndTimingsApiGetFiiDataOpts{
		From: optional.NewString(daysAgo(14)),
	})
	check(t, resp, err)
}

func TestGetDiiData(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewMarketHolidaysAndTimingsApi(c).GetDiiData(ctx(t), "NSE_EQ|CASH", "1D", nil)
	check(t, resp, err)
}

func TestGetSmartlistOptions(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewMarketHolidaysAndTimingsApi(c).GetSmartlistOptions(ctx(t), &api.MarketHolidaysAndTimingsApiGetSmartlistOptionsOpts{
		AssetType:  optional.NewString("INDEX"),
		Category:   optional.NewString("TOP_TRADED"),
		PageNumber: optional.NewInt32(1),
		PageSize:   optional.NewInt32(50),
	})
	check(t, resp, err)
}

func TestGetSmartlistFutures(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewMarketHolidaysAndTimingsApi(c).GetSmartlistFutures(ctx(t), &api.MarketHolidaysAndTimingsApiGetSmartlistFuturesOpts{
		AssetType:  optional.NewString("STOCK"),
		Category:   optional.NewString("MOST_ACTIVE"),
		PageNumber: optional.NewInt32(1),
		PageSize:   optional.NewInt32(50),
	})
	check(t, resp, err)
}

func TestGetSmartlistMtf(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewMarketHolidaysAndTimingsApi(c).GetSmartlistMtf(ctx(t), &api.MarketHolidaysAndTimingsApiGetSmartlistMtfOpts{
		PageNumber: optional.NewInt32(1),
		PageSize:   optional.NewInt32(50),
	})
	check(t, resp, err)
}
