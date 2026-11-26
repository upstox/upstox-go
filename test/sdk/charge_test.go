//go:build sdk

package sdk

import (
	"testing"

	"github.com/upstox/upstox-go/upstox_client/api"
	"github.com/upstox/upstox-go/upstox_client/models"
)

func TestGetBrokerage(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewChargeApi(c).GetBrokerage(ctx(t), ideaKey, 10, "D", "BUY", 13.4)
	check(t, resp, err)
	if resp.Data == nil || resp.Data.Charges == nil {
		t.Fatal("brokerage response has no charges")
	}
}

func TestGetBrokerageIntradaySell(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewChargeApi(c).GetBrokerage(ctx(t), relianceKey, 1, "I", "SELL", 1300)
	check(t, resp, err)
}

func TestPostMargin(t *testing.T) {
	c := client(t)
	body := models.MarginRequest{Instruments: []models.Instrument{
		{InstrumentKey: "NSE_EQ|INE528G01035", Quantity: 12, Product: "D", TransactionType: "BUY"},
		{InstrumentKey: relianceKey, Quantity: 1, Product: "D", TransactionType: "BUY"},
	}}
	resp, _, err := api.NewChargeApi(c).PostMargin(ctx(t), body)
	check(t, resp, err)
}

func TestPostMarginOption(t *testing.T) {
	c := client(t)
	opt := niftyOption(t, c)
	body := models.MarginRequest{Instruments: []models.Instrument{
		{InstrumentKey: opt.InstrumentKey, Quantity: int32(opt.LotSize), Product: "D", TransactionType: "SELL"},
	}}
	resp, _, err := api.NewChargeApi(c).PostMargin(ctx(t), body)
	check(t, resp, err)
}
