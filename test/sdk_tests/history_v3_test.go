//go:build sdk

package sdk_tests

import (
	"testing"

	"github.com/upstox/upstox-go/upstox_client/api"
)

func TestGetHistoricalCandleDataV3(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewHistoryV3Api(c).GetHistoricalCandleData(ctx(t), ideaKey, "days", 1, daysAgo(1))
	check(t, resp, err)
}

func TestGetHistoricalCandleDataRangeV3(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewHistoryV3Api(c).GetHistoricalCandleDataRange(ctx(t), ideaKey, "days", 1, daysAgo(1), daysAgo(30))
	check(t, resp, err)
}

func TestGetHistoricalCandleDataRangeV3Minutes(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewHistoryV3Api(c).GetHistoricalCandleDataRange(ctx(t), ideaKey, "minutes", 15, daysAgo(1), daysAgo(5))
	check(t, resp, err)
}

func TestGetHistoricalCandleDataV3Weekly(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewHistoryV3Api(c).GetHistoricalCandleData(ctx(t), niftyKey, "weeks", 1, daysAgo(1))
	check(t, resp, err)
}

func TestGetIntraDayCandleDataV3(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewHistoryV3Api(c).GetIntraDayCandleData(ctx(t), ideaKey, "minutes", 1)
	check(t, resp, err)
}
