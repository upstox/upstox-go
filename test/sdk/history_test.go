//go:build sdk

package sdk

import (
	"testing"

	"github.com/upstox/upstox-go/upstox_client/api"
)

func TestGetHistoricalCandleData(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewHistoryApi(c).GetHistoricalCandleData(ctx(t), ideaKey, "day", daysAgo(1))
	check(t, resp, err)
	if resp.Data == nil || len(resp.Data.Candles) == 0 {
		t.Fatal("no candles returned")
	}
}

func TestGetHistoricalCandleDataRange(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewHistoryApi(c).GetHistoricalCandleDataRange(ctx(t), ideaKey, "day", daysAgo(1), daysAgo(30))
	check(t, resp, err)
}

func TestGetHistoricalCandleDataRangeMinute(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewHistoryApi(c).GetHistoricalCandleDataRange(ctx(t), ideaKey, "1minute", daysAgo(1), daysAgo(3))
	check(t, resp, err)
}

func TestGetIntraDayCandleData(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewHistoryApi(c).GetIntraDayCandleData(ctx(t), ideaKey, "1minute")
	check(t, resp, err)
}
