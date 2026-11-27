//go:build sdk

package sdk_tests

import (
	"testing"

	"github.com/antihax/optional"
	"github.com/upstox/upstox-go/upstox_client/api"
)

func TestGetCompanyProfile(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewFundamentalsApi(c).GetCompanyProfile(ctx(t), relianceISIN)
	check(t, resp, err)
}

func TestGetBalanceSheet(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewFundamentalsApi(c).GetBalanceSheet(ctx(t), relianceISIN, nil)
	check(t, resp, err)
}

func TestGetBalanceSheetStandalone(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewFundamentalsApi(c).GetBalanceSheet(ctx(t), relianceISIN, &api.FundamentalsApiGetBalanceSheetOpts{
		Type_: optional.NewString("standalone"),
	})
	check(t, resp, err)
}

func TestGetCashFlow(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewFundamentalsApi(c).GetCashFlow(ctx(t), relianceISIN, nil)
	check(t, resp, err)
}

func TestGetIncomeStatement(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewFundamentalsApi(c).GetIncomeStatement(ctx(t), relianceISIN, nil)
	check(t, resp, err)
}

func TestGetIncomeStatementQuarterly(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewFundamentalsApi(c).GetIncomeStatement(ctx(t), relianceISIN, &api.FundamentalsApiGetIncomeStatementOpts{
		TimePeriod: optional.NewString("quarterly"),
	})
	check(t, resp, err)
}

func TestGetKeyRatios(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewFundamentalsApi(c).GetKeyRatios(ctx(t), relianceISIN)
	check(t, resp, err)
}

func TestGetShareHoldings(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewFundamentalsApi(c).GetShareHoldings(ctx(t), relianceISIN)
	check(t, resp, err)
}

func TestGetCorporateActions(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewFundamentalsApi(c).GetCorporateActions(ctx(t), relianceISIN)
	check(t, resp, err)
}

func TestGetCompetitors(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewFundamentalsApi(c).GetCompetitors(ctx(t), relianceKey)
	check(t, resp, err)
}
