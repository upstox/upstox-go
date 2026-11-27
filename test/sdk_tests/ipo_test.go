//go:build sdk

package sdk_tests

import (
	"testing"

	"github.com/antihax/optional"
	"github.com/upstox/upstox-go/upstox_client/api"
	"github.com/upstox/upstox-go/upstox_client/models"
)

func TestGetIpoListing(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewIPOApi(c).GetIpoListing(ctx(t), nil)
	check(t, resp, err)
}

func TestGetIpoListingFiltered(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewIPOApi(c).GetIpoListing(ctx(t), &api.IPOApiGetIpoListingOpts{
		Status:     optional.NewString("open"),
		IssueType:  optional.NewString("regular"),
		PageNumber: optional.NewString("1"),
		Records:    optional.NewString("20"),
	})
	check(t, resp, err)
}

func TestGetIpoDetails(t *testing.T) {
	c := client(t)
	ipo := api.NewIPOApi(c)
	list, _, err := ipo.GetIpoListing(ctx(t), &api.IPOApiGetIpoListingOpts{Status: optional.NewString("listed")})
	check(t, list, err)
	if len(list.Data) == 0 {
		t.Skip("no listed IPOs to look up")
	}
	resp, _, err := ipo.GetIpoDetails(ctx(t), list.Data[0].Id)
	check(t, resp, err)
	if resp.Data == nil || resp.Data.Id != list.Data[0].Id {
		t.Fatalf("details id mismatch for %s", list.Data[0].Id)
	}
}

func TestGetIpoDetailsUnknown(t *testing.T) {
	c := client(t)
	_, httpResp, err := api.NewIPOApi(c).GetIpoDetails(ctx(t), "sdk-go-no-such-ipo")
	checkRejected(t, httpResp, err)
}

func TestGetIpoOrders(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewIPOApi(c).GetIpoOrders(ctx(t), &api.IPOApiGetIpoOrdersOpts{
		PageNumber: optional.NewString("1"),
		Records:    optional.NewString("10"),
	})
	check(t, resp, err)
}

func TestGetIpoOrderById(t *testing.T) {
	c := client(t)
	ipo := api.NewIPOApi(c)
	orders, _, err := ipo.GetIpoOrders(ctx(t), nil)
	check(t, orders, err)
	if len(orders.Data) == 0 {
		_, httpResp, err := ipo.GetIpoOrderById(ctx(t), "sdk-go-no-such-order")
		checkRejected(t, httpResp, err)
		return
	}
	resp, _, err := ipo.GetIpoOrderById(ctx(t), orders.Data[0].OrderId)
	check(t, resp, err)
}

// TestApplyAndCancelIpo blocks funds through a real UPI mandate, so it needs an explicit IPO id and UPI handle.
func TestApplyAndCancelIpo(t *testing.T) {
	c := client(t)
	env := requireEnv(t, "UPSTOX_TEST_IPO_ID", "UPSTOX_TEST_UPI")
	ipo := api.NewIPOApi(c)
	details, _, err := ipo.GetIpoDetails(ctx(t), env[0])
	check(t, details, err)
	d := details.Data
	if d == nil || len(d.Investors) == 0 {
		t.Skip("IPO details carry no investor categories")
	}
	qty := d.MinimumQuantity
	if qty == 0 {
		qty = d.LotSize
	}
	body := models.IpoApplyRequest{
		Id:       env[0],
		Upi:      env[1],
		Category: d.Investors[0].Category,
		Bids:     []models.IpoBidRequest{{Quantity: qty, Price: d.MinimumPrice}},
	}
	resp, _, err := ipo.ApplyForIpo(ctx(t), body)
	check(t, resp, err)
	if resp.Data == nil || resp.Data.OrderId == "" {
		t.Fatal("apply returned no order id")
	}
	cancel, _, err := ipo.CancelIpoOrder(ctx(t), resp.Data.OrderId)
	check(t, cancel, err)
}

func TestCancelIpoOrderUnknown(t *testing.T) {
	c := client(t)
	_, httpResp, err := api.NewIPOApi(c).CancelIpoOrder(ctx(t), "sdk-go-no-such-order")
	checkRejected(t, httpResp, err)
}
