//go:build sdk

package sdk_tests

import (
	"testing"

	"github.com/antihax/optional"

	"github.com/upstox/upstox-go/upstox_client/api"
)

func TestGetTradeWiseProfitAndLossMetaData(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewTradeProfitAndLossApi(c).GetTradeWiseProfitAndLossMetaData(ctx(t), "EQ", previousFinancialYear(), nil)
	check(t, resp, err)
}

func TestGetTradeWiseProfitAndLossData(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewTradeProfitAndLossApi(c).GetTradeWiseProfitAndLossData(ctx(t), "EQ", previousFinancialYear(), 1, 4, nil)
	check(t, resp, err)
}

func TestGetTradeWiseProfitAndLossDataFO(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewTradeProfitAndLossApi(c).GetTradeWiseProfitAndLossData(ctx(t), "FO", previousFinancialYear(), 1, 4, nil)
	check(t, resp, err)
}

func TestGetProfitAndLossCharges(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewTradeProfitAndLossApi(c).GetProfitAndLossCharges(ctx(t), "EQ", previousFinancialYear(), nil)
	check(t, resp, err)
}

func TestGetProfitAndLossChargesDateRange(t *testing.T) {
	c := client(t)
	fy := previousFinancialYear()
	// Dates are dd-mm-yyyy and must fall inside the financial year
	resp, _, err := api.NewTradeProfitAndLossApi(c).GetProfitAndLossCharges(ctx(t), "EQ", fy, &api.TradeProfitAndLossApiGetProfitAndLossChargesOpts{
		FromDate: optional.NewString("01-04-20" + fy[:2]),
		ToDate:   optional.NewString("31-03-20" + fy[2:]),
	})
	check(t, resp, err)
}
