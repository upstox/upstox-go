//go:build sdk

package sdk_tests

import (
	"testing"

	"github.com/antihax/optional"
	"github.com/upstox/upstox-go/upstox_client/api"
)

func TestGetTradesByDateRange(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewPostTradeApi(c).GetTradesByDateRange(ctx(t), daysAgo(90), daysAgo(1), 1, 100, nil)
	check(t, resp, err)
}

func TestGetTradesByDateRangeSegment(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewPostTradeApi(c).GetTradesByDateRange(ctx(t), daysAgo(90), daysAgo(1), 1, 50, &api.PostTradeApiGetTradesByDateRangeOpts{
		Segment: optional.NewString("EQ"),
	})
	check(t, resp, err)
}
