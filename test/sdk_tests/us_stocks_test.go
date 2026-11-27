//go:build sdk

package sdk_tests

import (
	"math"
	"os"
	"testing"

	"github.com/antihax/optional"
	"github.com/upstox/upstox-go/upstox_client/api"
	"github.com/upstox/upstox-go/upstox_client/models"
)

// usKey returns UPSTOX_TEST_US_INSTRUMENT, defaulting to Coursera.
func usKey() string {
	if v := os.Getenv("UPSTOX_TEST_US_INSTRUMENT"); v != "" {
		return v
	}
	return "US_EQ|COUR"
}

func TestGetUsMarketStatus(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewUSMarketHolidaysAndTimingsApi(c).GetUsMarketStatus(ctx(t))
	check(t, resp, err)
}

func TestGetUsHolidays(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewUSMarketHolidaysAndTimingsApi(c).GetUsHolidays(ctx(t))
	check(t, resp, err)
}

func TestGetUsHoliday(t *testing.T) {
	c := client(t)
	// Thanksgiving 2026: NYSE closed
	resp, _, err := api.NewUSMarketHolidaysAndTimingsApi(c).GetUsHoliday(ctx(t), "2026-11-26")
	check(t, resp, err)
}

func TestGetUsExchangeTimings(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewUSMarketHolidaysAndTimingsApi(c).GetUsExchangeTimings(ctx(t), "2026-11-27")
	check(t, resp, err)
}

func TestGetUsFunds(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewUSStocksFundsApi(c).GetUsFunds(ctx(t))
	check(t, resp, err)
}

func TestGetUsHoldings(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewUSStocksPortfolioApi(c).GetUsHoldings(ctx(t))
	check(t, resp, err)
}

func TestGetUsOrderBook(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewUSStocksOrderApi(c).GetUsOrderBook(ctx(t), nil)
	check(t, resp, err)
}

func TestGetUsOrderBookByTag(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewUSStocksOrderApi(c).GetUsOrderBook(ctx(t), &api.USStocksOrderApiGetUsOrderBookOpts{
		Tag: optional.NewString(testTag),
	})
	check(t, resp, err)
}

func TestGetUsOrderDetailsUnknown(t *testing.T) {
	c := client(t)
	_, httpResp, err := api.NewUSStocksOrderApi(c).GetUsOrderDetails(ctx(t), "sdk-go-no-such-order")
	checkRejected(t, httpResp, err)
}

func TestCancelUsOrderUnknown(t *testing.T) {
	c := client(t)
	_, httpResp, err := api.NewUSStocksOrderApi(c).CancelUsOrder(ctx(t), "sdk-go-no-such-order")
	checkRejected(t, httpResp, err)
}

func TestEstimateUsOrder(t *testing.T) {
	c := client(t)
	body := models.UsOrderEstimateRequest{Quantity: 1, InstrumentKey: usKey(), TransactionType: "BUY", OrderType: "MARKET"}
	resp, _, err := api.NewUSStocksOrderApi(c).EstimateUsOrder(ctx(t), body)
	check(t, resp, err)
}

func TestCalculateUsMargin(t *testing.T) {
	c := client(t)
	body := models.UsMarginRequest{Instruments: []models.UsOrderEstimateRequest{
		{Quantity: 1, InstrumentKey: usKey(), TransactionType: "BUY", OrderType: "MARKET"},
	}}
	resp, _, err := api.NewUSStocksChargesApi(c).CalculateUsMargin(ctx(t), body)
	check(t, resp, err)
}

// TestPlaceAndCancelUsOrder places a 1-share LIMIT BUY at half the estimated price, then cancels it.
func TestPlaceAndCancelUsOrder(t *testing.T) {
	c := client(t)
	requireOrderTests(t)
	o := api.NewUSStocksOrderApi(c)
	est, _, err := o.EstimateUsOrder(ctx(t), models.UsOrderEstimateRequest{
		Quantity: 1, InstrumentKey: usKey(), TransactionType: "BUY", OrderType: "MARKET",
	})
	check(t, est, err)
	if est.Data == nil || est.Data.TradeValue <= 0 {
		t.Skip("estimate returned no trade value to price the order from")
	}
	price := math.Max(0.01, math.Floor(est.Data.TradeValue*50)/100)
	resp, _, err := o.PlaceUsOrder(ctx(t), models.UsPlaceOrderRequest{
		Quantity: 1, Price: price, Validity: "DAY", Tag: testTag,
		InstrumentKey: usKey(), TransactionType: "BUY", OrderType: "LIMIT",
	})
	check(t, resp, err)
	if resp.Data == nil || resp.Data.OrderId == "" {
		t.Fatal("place US order returned no order id")
	}
	id := resp.Data.OrderId
	cancelled := false
	t.Cleanup(func() {
		if cancelled {
			return
		}
		cctx, cancel := cleanupCtx()
		defer cancel()
		if _, _, err := o.CancelUsOrder(cctx, id); err != nil {
			t.Logf("cleanup cancel US order %s: %v %s", id, err, errorBody(err))
		}
	})

	details, _, err := o.GetUsOrderDetails(ctx(t), id)
	check(t, details, err)

	can, _, err := o.CancelUsOrder(ctx(t), id)
	check(t, can, err)
	cancelled = true
}
