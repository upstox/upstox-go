//go:build sdk

package sdk

import (
	"strings"
	"testing"

	"github.com/upstox/upstox-go/upstox/api"
	"github.com/upstox/upstox-go/upstox/models"
)

// Sandbox orders are simulated, so these run with just UPSTOX_SANDBOX_ACCESS_TOKEN.

func TestSandboxPlaceModifyCancelOrder(t *testing.T) {
	c := sandboxClient(t)
	o := api.NewOrderApi(c)
	placed, _, err := o.PlaceOrder(ctx(t), models.PlaceOrderRequest{
		Quantity: 1, Product: "D", Validity: "DAY", Price: 12, Tag: testTag,
		InstrumentToken: "NSE_EQ|INE528G01035", OrderType: "LIMIT", TransactionType: "BUY",
	}, nil)
	check(t, placed, err)
	if placed.Data == nil || placed.Data.OrderId == "" {
		t.Fatal("sandbox place returned no order id")
	}
	id := placed.Data.OrderId

	mod, _, err := o.ModifyOrder(ctx(t), models.ModifyOrderRequest{
		OrderId: id, Quantity: 2, Validity: "DAY", Price: 11, OrderType: "LIMIT",
	}, nil)
	check(t, mod, err)

	can, _, err := o.CancelOrder(ctx(t), id, nil)
	check(t, can, err)
}

func TestSandboxModifyOrderUnknown(t *testing.T) {
	c := sandboxClient(t)
	resp, _, err := api.NewOrderApi(c).ModifyOrder(ctx(t), models.ModifyOrderRequest{
		OrderId: "250121112512866", Quantity: 2, Validity: "DAY", OrderType: "MARKET",
	}, nil)
	checkAllow(t, resp, err, "UDAPI100010")
}

func TestSandboxPlaceMultiOrder(t *testing.T) {
	c := sandboxClient(t)
	body := []models.MultiOrderRequest{
		{Quantity: 1, Product: "I", Validity: "DAY", Tag: testTag, InstrumentToken: ideaKey,
			OrderType: "MARKET", TransactionType: "BUY", CorrelationId: "1"},
		{Quantity: 1, Product: "D", Validity: "DAY", Price: 8.9, Tag: testTag, InstrumentToken: ideaKey,
			OrderType: "LIMIT", TransactionType: "BUY", CorrelationId: "2"},
	}
	resp, _, err := api.NewOrderApi(c).PlaceMultiOrder(ctx(t), body, nil)
	check(t, resp, err)
}

func TestSandboxPlaceModifyCancelOrderV3(t *testing.T) {
	c := sandboxClient(t)
	o := api.NewOrderV3Api(c)
	placed, _, err := o.PlaceOrderV3(ctx(t), models.PlaceOrderV3Request{
		Quantity: 1, Product: "D", Validity: "DAY", Price: 9.12, Tag: testTag, Slice: true,
		InstrumentToken: ideaKey, OrderType: "LIMIT", TransactionType: "BUY",
	}, nil)
	check(t, placed, err)
	if placed.Data == nil || len(placed.Data.OrderIds) == 0 {
		t.Fatal("sandbox place v3 returned no order ids")
	}
	id := placed.Data.OrderIds[0]

	mod, _, err := o.ModifyOrderV3(ctx(t), models.ModifyOrderRequest{
		OrderId: id, Quantity: 1, Validity: "DAY", Price: 9.0, OrderType: "LIMIT",
	}, nil)
	check(t, mod, err)

	can, _, err := o.CancelOrderV3(ctx(t), id, nil)
	check(t, can, err)
}

// Non-order endpoints are refused client-side in sandbox mode, before any request is sent.
func TestSandboxRejectsNonOrderEndpoint(t *testing.T) {
	c := sandboxClient(t)
	_, _, err := api.NewUserApi(c).GetProfile(ctx(t))
	if err == nil || !strings.Contains(err.Error(), "not available in sandbox mode") {
		t.Fatalf("expected sandbox refusal, got %v", err)
	}
}
