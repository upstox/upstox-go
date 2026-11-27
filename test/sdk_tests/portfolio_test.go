//go:build sdk

package sdk_tests

import (
	"testing"

	"github.com/upstox/upstox-go/upstox_client/api"
	"github.com/upstox/upstox-go/upstox_client/models"
)

func TestGetHoldings(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewPortfolioApi(c).GetHoldings(ctx(t))
	check(t, resp, err)
}

func TestGetPositions(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewPortfolioApi(c).GetPositions(ctx(t))
	check(t, resp, err)
}

func TestGetMtfPositions(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewPortfolioApi(c).GetMtfPositions(ctx(t))
	check(t, resp, err)
}

// TestConvertPositions converts an intraday position if one exists, so it is gated with the order tests.
func TestConvertPositions(t *testing.T) {
	c := client(t)
	requireOrderTests(t)
	body := models.ConvertPositionRequest{
		InstrumentToken: "NSE_EQ|INE528G01035", NewProduct: "D", OldProduct: "I", TransactionType: "BUY", Quantity: 1,
	}
	resp, _, err := api.NewPortfolioApi(c).ConvertPositions(ctx(t), body)
	// UDAPI1035: no matching position to convert
	checkAllow(t, resp, err, "UDAPI1035")
}
