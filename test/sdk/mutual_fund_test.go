//go:build sdk

package sdk

import (
	"testing"

	"github.com/antihax/optional"
	"github.com/upstox/upstox-go/upstox/api"
)

func TestGetMutualFundHoldings(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewMutualFundApi(c).GetMutualFundHoldings(ctx(t))
	check(t, resp, err)
}

func TestGetMutualFundOrders(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewMutualFundApi(c).GetMutualFundOrders(ctx(t), nil)
	check(t, resp, err)
}

func TestGetMutualFundOrdersFiltered(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewMutualFundApi(c).GetMutualFundOrders(ctx(t), &api.MutualFundApiGetMutualFundOrdersOpts{
		Status:          optional.NewString("COMPLETED"),
		TransactionType: optional.NewString("BUY"),
		PageNumber:      optional.NewInt32(1),
		Records:         optional.NewInt32(10),
	})
	check(t, resp, err)
}

func TestGetMutualFundOrderUnknown(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewMutualFundApi(c).GetMutualFundOrder(ctx(t), "MF-ORDER-001")
	checkAllow(t, resp, err, "UDAPI100010", "UDAPI100050")
}

func TestGetMutualFundSips(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewMutualFundApi(c).GetMutualFundSips(ctx(t), nil)
	check(t, resp, err)
}

func TestGetMutualFundSipsPaged(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewMutualFundApi(c).GetMutualFundSips(ctx(t), &api.MutualFundApiGetMutualFundSipsOpts{
		PageNumber: optional.NewInt32(1),
		Records:    optional.NewInt32(10),
	})
	check(t, resp, err)
}
