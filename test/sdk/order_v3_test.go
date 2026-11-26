//go:build sdk

package sdk

import (
	"math"
	"testing"

	"github.com/antihax/optional"
	"github.com/upstox/upstox-go/upstox"
	"github.com/upstox/upstox-go/upstox/api"
	"github.com/upstox/upstox-go/upstox/models"
)

func TestModifyOrderV3Unknown(t *testing.T) {
	c := client(t)
	body := models.ModifyOrderRequest{OrderId: missingOrderID, Quantity: 1, Validity: "DAY", Price: 5, OrderType: "LIMIT"}
	resp, _, err := api.NewOrderV3Api(c).ModifyOrderV3(ctx(t), body, &api.OrderV3ApiModifyOrderV3Opts{XAlgoName: algoName()})
	checkAllow(t, resp, err, "UDAPI100010")
}

func TestCancelOrderV3Unknown(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewOrderV3Api(c).CancelOrderV3(ctx(t), missingOrderID, &api.OrderV3ApiCancelOrderV3Opts{XAlgoName: algoName()})
	checkAllow(t, resp, err, "UDAPI100010")
}

func TestPlaceModifyCancelOrderV3(t *testing.T) {
	c := client(t)
	requireOrderTests(t)
	o := api.NewOrderV3Api(c)
	price, amo := testOrderParams(t, c)
	body := models.PlaceOrderV3Request{
		Quantity: 1, Product: "D", Validity: "DAY", Price: price, Tag: testTag, Slice: true,
		InstrumentToken: ideaKey, OrderType: "LIMIT", TransactionType: "BUY", IsAmo: amo,
	}
	placed, _, err := o.PlaceOrderV3(ctx(t), body, &api.OrderV3ApiPlaceOrderV3Opts{XAlgoName: algoName()})
	check(t, placed, err)
	if placed.Data == nil || len(placed.Data.OrderIds) == 0 {
		t.Fatal("place v3 returned no order ids")
	}
	for _, id := range placed.Data.OrderIds {
		cancelOnCleanup(t, c, id)
	}
	id := placed.Data.OrderIds[0]

	st, _, err := api.NewOrderApi(c).GetOrderStatus(ctx(t), &api.OrderApiGetOrderStatusOpts{OrderId: optional.NewString(id)})
	check(t, st, err)
	if st.Data != nil && st.Data.Status == "rejected" {
		t.Skipf("order %s rejected by RMS/exchange (%s); skipping modify", id, st.Data.StatusMessage)
	}

	mod := models.ModifyOrderRequest{OrderId: id, Quantity: 1, Validity: "DAY", Price: farLimitPrice(price, 0.01), OrderType: "LIMIT"}
	modified, _, err := o.ModifyOrderV3(ctx(t), mod, &api.OrderV3ApiModifyOrderV3Opts{XAlgoName: algoName()})
	check(t, modified, err)

	cancelled, _, err := o.CancelOrderV3(ctx(t), id, &api.OrderV3ApiCancelOrderV3Opts{XAlgoName: algoName()})
	check(t, cancelled, err)
}

func TestGetGttOrderDetails(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewOrderV3Api(c).GetGttOrderDetails(ctx(t), nil)
	check(t, resp, err)
}

func TestGetGttOrderDetailsUnknown(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewOrderV3Api(c).GetGttOrderDetails(ctx(t), &api.OrderV3ApiGetGttOrderDetailsOpts{
		GttOrderId: optional.NewString(missingGTTID),
	})
	checkAllow(t, resp, err, "UDAPI100010")
}

func TestModifyGTTOrderUnknown(t *testing.T) {
	c := client(t)
	body := models.GttModifyOrderRequest{
		Type_: "SINGLE", Quantity: 1, GttOrderId: missingGTTID,
		Rules: []models.GttRule{{Strategy: "ENTRY", TriggerType: "BELOW", TriggerPrice: 5}},
	}
	resp, _, err := api.NewOrderV3Api(c).ModifyGTTOrder(ctx(t), body, &api.OrderV3ApiModifyGTTOrderOpts{XAlgoName: algoName()})
	checkAllow(t, resp, err, "UDAPI100010")
}

func TestCancelGTTOrderUnknown(t *testing.T) {
	c := client(t)
	body := models.GttCancelOrderRequest{GttOrderId: missingGTTID}
	resp, _, err := api.NewOrderV3Api(c).CancelGTTOrder(ctx(t), body, &api.OrderV3ApiCancelGTTOrderOpts{XAlgoName: algoName()})
	checkAllow(t, resp, err, "UDAPI100010")
}

// placeGTT places a GTT and registers its cancellation, returning the GTT order ID.
func placeGTT(t *testing.T, c *upstox.APIClient, body models.GttPlaceOrderRequest) string {
	t.Helper()
	o := api.NewOrderV3Api(c)
	resp, _, err := o.PlaceGTTOrder(ctx(t), body, &api.OrderV3ApiPlaceGTTOrderOpts{XAlgoName: algoName()})
	check(t, resp, err)
	if resp.Data == nil || len(resp.Data.GttOrderIds) == 0 {
		t.Fatal("place GTT returned no id")
	}
	id := resp.Data.GttOrderIds[0]
	t.Cleanup(func() {
		cctx, cancel := cleanupCtx()
		defer cancel()
		if _, _, err := o.CancelGTTOrder(cctx, models.GttCancelOrderRequest{GttOrderId: id}, &api.OrderV3ApiCancelGTTOrderOpts{XAlgoName: algoName()}); err != nil {
			t.Logf("cleanup cancel GTT %s: %v %s", id, err, errorBody(err))
		}
	})
	return id
}

// TestPlaceModifyCancelGTTOrder buys only if the price falls 20%, so it will not trigger during the test.
func TestPlaceModifyCancelGTTOrder(t *testing.T) {
	c := client(t)
	requireOrderTests(t)
	o := api.NewOrderV3Api(c)
	last := ltp(t, c, ideaKey)
	id := placeGTT(t, c, models.GttPlaceOrderRequest{
		Type_: "SINGLE", Quantity: 1, Product: "D", InstrumentToken: ideaKey, TransactionType: "BUY",
		Rules: []models.GttRule{{Strategy: "ENTRY", TriggerType: "BELOW", TriggerPrice: farLimitPrice(last, 0.20)}},
	})

	details, _, err := o.GetGttOrderDetails(ctx(t), &api.OrderV3ApiGetGttOrderDetailsOpts{GttOrderId: optional.NewString(id)})
	check(t, details, err)
	if len(details.Data) == 0 || details.Data[0].GttOrderId != id {
		t.Fatalf("GTT details missing %s", id)
	}

	mod := models.GttModifyOrderRequest{
		Type_: "SINGLE", Quantity: 1, GttOrderId: id,
		Rules: []models.GttRule{{Strategy: "ENTRY", TriggerType: "BELOW", TriggerPrice: farLimitPrice(last, 0.25)}},
	}
	modified, _, err := o.ModifyGTTOrder(ctx(t), mod, &api.OrderV3ApiModifyGTTOrderOpts{XAlgoName: algoName()})
	check(t, modified, err)

	cancelled, _, err := o.CancelGTTOrder(ctx(t), models.GttCancelOrderRequest{GttOrderId: id}, &api.OrderV3ApiCancelGTTOrderOpts{XAlgoName: algoName()})
	check(t, cancelled, err)
}

// TestPlaceGTTMultiLegTrailingStop mirrors place_gtt_tsl_order.py: entry, target and a trailing stop-loss.
func TestPlaceGTTMultiLegTrailingStop(t *testing.T) {
	c := client(t)
	requireOrderTests(t)
	last := ltp(t, c, ideaKey)
	entry := farLimitPrice(last, 0.20)
	gap := math.Max(0.05, math.Floor(entry*0.02*20)/20)
	rules := []models.GttRule{
		{Strategy: "ENTRY", TriggerType: "BELOW", TriggerPrice: entry},
		{Strategy: "TARGET", TriggerType: "IMMEDIATE", TriggerPrice: farLimitPrice(entry, -0.10)},
		{Strategy: "STOPLOSS", TriggerType: "IMMEDIATE", TriggerPrice: farLimitPrice(entry, 0.05), TrailingGap: gap},
	}
	id := placeGTT(t, c, models.GttPlaceOrderRequest{
		Type_: "MULTIPLE", Quantity: 1, Product: "D", InstrumentToken: ideaKey, TransactionType: "BUY", Rules: rules,
	})
	mod := models.GttModifyOrderRequest{Type_: "MULTIPLE", Quantity: 2, GttOrderId: id, Rules: rules}
	modified, _, err := api.NewOrderV3Api(c).ModifyGTTOrder(ctx(t), mod, &api.OrderV3ApiModifyGTTOrderOpts{XAlgoName: algoName()})
	check(t, modified, err)
}
