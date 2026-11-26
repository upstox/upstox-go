//go:build sdk

package sdk

import (
	"os"
	"testing"

	"github.com/upstox/upstox-go/upstox/api"
	"github.com/upstox/upstox-go/upstox/models"
)

const missingPayoutID = "sdk-go-no-such-transaction"

func TestGetPayoutModes(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewUserApi(c).GetPayoutModes(ctx(t))
	check(t, resp, err)
}

func TestGetPayoutHistory(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewUserApi(c).GetPayoutHistory(ctx(t))
	check(t, resp, err)
}

func TestGetPayinHistory(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewUserApi(c).GetPayinHistory(ctx(t))
	check(t, resp, err)
}

func TestModifyPayoutUnknown(t *testing.T) {
	c := client(t)
	_, httpResp, err := api.NewUserApi(c).ModifyPayout(ctx(t), models.ModifyPayoutRequest{Amount: 2}, missingPayoutID)
	checkRejected(t, httpResp, err)
}

func TestCancelPayoutUnknown(t *testing.T) {
	c := client(t)
	_, httpResp, err := api.NewUserApi(c).CancelPayout(ctx(t), missingPayoutID)
	checkRejected(t, httpResp, err)
}

// TestInitiateModifyCancelPayout requests a real ₹1 withdrawal, then modifies and cancels it.
func TestInitiateModifyCancelPayout(t *testing.T) {
	c := client(t)
	requireFlag(t, "UPSTOX_RUN_PAYOUT_TESTS")
	mode := os.Getenv("UPSTOX_PAYOUT_MODE")
	if mode == "" {
		mode = "IMPS"
	}
	u := api.NewUserApi(c)
	resp, _, err := u.InitiatePayout(ctx(t), models.InitiatePayoutRequest{Mode: mode, Amount: 1})
	check(t, resp, err)
	if resp.Data == nil || resp.Data.TransactionId == "" {
		t.Fatal("initiate payout returned no transaction id")
	}
	id := resp.Data.TransactionId
	cancelled := false
	t.Cleanup(func() {
		if cancelled {
			return
		}
		cctx, cancel := cleanupCtx()
		defer cancel()
		if _, _, err := u.CancelPayout(cctx, id); err != nil {
			t.Logf("cleanup cancel payout %s: %v %s", id, err, errorBody(err))
		}
	})

	mod, _, err := u.ModifyPayout(ctx(t), models.ModifyPayoutRequest{Amount: 2}, id)
	check(t, mod, err)

	can, _, err := u.CancelPayout(ctx(t), id)
	check(t, can, err)
	cancelled = true
}
