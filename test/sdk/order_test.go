//go:build sdk

package sdk

import (
	"testing"

	"github.com/antihax/optional"
	"github.com/upstox/upstox-go/upstox"
	"github.com/upstox/upstox-go/upstox/api"
	"github.com/upstox/upstox-go/upstox/models"
)

func TestGetOrderBook(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewOrderApi(c).GetOrderBook(ctx(t))
	check(t, resp, err)
}

func TestGetTradeHistory(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewOrderApi(c).GetTradeHistory(ctx(t))
	check(t, resp, err)
}

func TestGetOrderDetailsUnknown(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewOrderApi(c).GetOrderDetails(ctx(t), &api.OrderApiGetOrderDetailsOpts{
		OrderId: optional.NewString(missingOrderID),
	})
	checkAllow(t, resp, err, "UDAPI100010")
}

func TestGetOrderStatusUnknown(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewOrderApi(c).GetOrderStatus(ctx(t), &api.OrderApiGetOrderStatusOpts{
		OrderId: optional.NewString(missingOrderID),
	})
	checkAllow(t, resp, err, "UDAPI100010")
}

func TestGetTradesByOrderUnknown(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewOrderApi(c).GetTradesByOrder(ctx(t), missingOrderID)
	checkAllow(t, resp, err, "UDAPI100010")
}

// Modify/cancel against a nonexistent order ID exercise the write path with no side effect.
func TestModifyOrderUnknown(t *testing.T) {
	c := client(t)
	body := models.ModifyOrderRequest{OrderId: missingOrderID, Quantity: 1, Validity: "DAY", Price: 5, OrderType: "LIMIT"}
	resp, _, err := api.NewOrderApi(c).ModifyOrder(ctx(t), body, &api.OrderApiModifyOrderOpts{XAlgoName: algoName()})
	checkAllow(t, resp, err, "UDAPI100010")
}

func TestCancelOrderUnknown(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewOrderApi(c).CancelOrder(ctx(t), missingOrderID, &api.OrderApiCancelOrderOpts{XAlgoName: algoName()})
	checkAllow(t, resp, err, "UDAPI100010")
}

// testOrderParams returns a 1-share BUY LIMIT price 5% under LTP, and whether to send it as AMO.
func testOrderParams(t *testing.T, c *upstox.APIClient) (float64, bool) {
	t.Helper()
	return farLimitPrice(ltp(t, c, ideaKey), 0.05), !marketOpen(t, c)
}

// cancelOnCleanup cancels a placed order when the test ends, whatever the outcome.
func cancelOnCleanup(t *testing.T, c *upstox.APIClient, orderID string) {
	t.Cleanup(func() {
		cctx, cancel := cleanupCtx()
		defer cancel()
		if _, _, err := api.NewOrderApi(c).CancelOrder(cctx, orderID, &api.OrderApiCancelOrderOpts{XAlgoName: algoName()}); err != nil {
			t.Logf("cleanup cancel %s: %v %s", orderID, err, errorBody(err))
		}
	})
}

func TestPlaceModifyCancelOrder(t *testing.T) {
	c := client(t)
	requireOrderTests(t)
	o := api.NewOrderApi(c)
	price, amo := testOrderParams(t, c)
	body := models.PlaceOrderRequest{
		Quantity: 1, Product: "D", Validity: "DAY", Price: price, Tag: testTag,
		InstrumentToken: ideaKey, OrderType: "LIMIT", TransactionType: "BUY", IsAmo: amo,
	}
	placed, _, err := o.PlaceOrder(ctx(t), body, &api.OrderApiPlaceOrderOpts{XAlgoName: algoName()})
	check(t, placed, err)
	if placed.Data == nil || placed.Data.OrderId == "" {
		t.Fatal("place returned no order id")
	}
	id := placed.Data.OrderId
	cancelOnCleanup(t, c, id)

	details, _, err := o.GetOrderDetails(ctx(t), &api.OrderApiGetOrderDetailsOpts{OrderId: optional.NewString(id)})
	check(t, details, err)
	st, _, err := o.GetOrderStatus(ctx(t), &api.OrderApiGetOrderStatusOpts{OrderId: optional.NewString(id)})
	check(t, st, err)
	if st.Data == nil || st.Data.OrderId != id {
		t.Fatalf("order status id mismatch for %s", id)
	}
	trades, _, err := o.GetTradesByOrder(ctx(t), id)
	check(t, trades, err)
	if st.Data.Status == "rejected" {
		t.Skipf("order %s rejected by RMS/exchange (%s); skipping modify", id, st.Data.StatusMessage)
	}

	mod := models.ModifyOrderRequest{OrderId: id, Quantity: 1, Validity: "DAY", Price: farLimitPrice(price, 0.01), OrderType: "LIMIT"}
	modified, _, err := o.ModifyOrder(ctx(t), mod, &api.OrderApiModifyOrderOpts{XAlgoName: algoName()})
	check(t, modified, err)

	cancelled, _, err := o.CancelOrder(ctx(t), id, &api.OrderApiCancelOrderOpts{XAlgoName: algoName()})
	check(t, cancelled, err)
}

func TestPlaceMultiOrderAndCancelByTag(t *testing.T) {
	c := client(t)
	requireOrderTests(t)
	o := api.NewOrderApi(c)
	price, amo := testOrderParams(t, c)
	leg := func(corr string, product string) models.MultiOrderRequest {
		return models.MultiOrderRequest{
			Quantity: 1, Product: product, Validity: "DAY", Price: price, Tag: testTag,
			InstrumentToken: ideaKey, OrderType: "LIMIT", TransactionType: "BUY", IsAmo: amo, CorrelationId: corr,
		}
	}
	resp, _, err := o.PlaceMultiOrder(ctx(t), []models.MultiOrderRequest{leg("1", "D"), leg("2", "I")},
		&api.OrderApiPlaceMultiOrderOpts{XAlgoName: algoName()})
	// Partial success still returns placed order IDs, which must be cancelled
	for _, d := range resp.Data {
		if d.OrderId != "" {
			cancelOnCleanup(t, c, d.OrderId)
		}
	}
	check(t, resp, err)
	for _, e := range resp.Errors {
		t.Logf("leg %s rejected: %s %s", e.CorrelationId, e.ErrorCode, e.Message)
	}
	if len(resp.Data) == 0 {
		t.Fatal("multi order placed nothing")
	}

	cancel, _, err := o.CancelMultiOrder(ctx(t), &api.OrderApiCancelMultiOrderOpts{
		Tag: optional.NewString(testTag), XAlgoName: algoName(),
	})
	checkAllow(t, cancel, err, "UDAPI1109")
}

// TestCancelMultiOrderByTag only cancels open orders carrying this suite's tag.
func TestCancelMultiOrderByTag(t *testing.T) {
	c := client(t)
	requireOrderTests(t)
	resp, _, err := api.NewOrderApi(c).CancelMultiOrder(ctx(t), &api.OrderApiCancelMultiOrderOpts{
		Tag: optional.NewString(testTag), Segment: optional.NewString("NSE_EQ"), XAlgoName: algoName(),
	})
	checkAllow(t, resp, err, "UDAPI1109")
}

// TestExitPositionsByTag squares off positions opened by orders carrying this suite's tag.
func TestExitPositionsByTag(t *testing.T) {
	c := client(t)
	requireFlag(t, "UPSTOX_RUN_EXIT_POSITIONS")
	resp, _, err := api.NewOrderApi(c).ExitPositions(ctx(t), &api.OrderApiExitPositionsOpts{
		Tag: optional.NewString(testTag), XAlgoName: algoName(),
	})
	checkAllow(t, resp, err, "UDAPI1111", "UDAPI1113")
}
