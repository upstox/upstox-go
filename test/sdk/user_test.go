//go:build sdk

package sdk

import (
	"testing"

	"github.com/antihax/optional"
	"github.com/upstox/upstox-go/upstox_client/api"
	"github.com/upstox/upstox-go/upstox_client/models"
)

func TestGetProfile(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewUserApi(c).GetProfile(ctx(t))
	check(t, resp, err)
	if resp.Data == nil || resp.Data.UserId == "" {
		t.Fatal("profile has no user id")
	}
}

func TestGetUserFundMargin(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewUserApi(c).GetUserFundMargin(ctx(t), nil)
	check(t, resp, err)
	if resp.Data == nil {
		t.Fatal("no fund data")
	}
}

func TestGetUserFundMarginSegment(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewUserApi(c).GetUserFundMargin(ctx(t), &api.UserApiGetUserFundMarginOpts{
		Segment: optional.NewString("SEC"),
	})
	check(t, resp, err)
}

func TestGetUserFundMarginV3(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewUserApi(c).GetUserFundMarginV3(ctx(t))
	check(t, resp, err)
	if resp.Data == nil {
		t.Fatal("no fund data")
	}
}

func TestGetUserIps(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewUserApi(c).GetUserIps(ctx(t))
	check(t, resp, err)
}

func TestGetKillSwitch(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewUserApi(c).GetKillSwitch(ctx(t))
	check(t, resp, err)
}

// TestUpdateUserIp writes the current primary IP back unchanged, as an IP change can revoke tokens.
func TestUpdateUserIp(t *testing.T) {
	c := client(t)
	requireFlag(t, "UPSTOX_RUN_ACCOUNT_SETTINGS_TESTS")
	u := api.NewUserApi(c)
	cur, _, err := u.GetUserIps(ctx(t))
	check(t, cur, err)
	if cur.Data == nil || cur.Data.PrimaryIp == "" {
		t.Skip("no primary IP registered to write back")
	}
	resp, _, err := u.UpdateUserIp(ctx(t), models.UpdateUserIpRequest{PrimaryIp: cur.Data.PrimaryIp, SecondaryIp: cur.Data.SecondaryIp})
	// UDAPI100068/UDAPI100069: update rejected (unchanged IP or rate limit), matching sanity.py
	checkAllow(t, resp, err, "UDAPI100068", "UDAPI100069")
}

// TestUpdateKillSwitch writes each segment's current state back, so trading access is unchanged.
func TestUpdateKillSwitch(t *testing.T) {
	c := client(t)
	requireFlag(t, "UPSTOX_RUN_ACCOUNT_SETTINGS_TESTS")
	u := api.NewUserApi(c)
	cur, _, err := u.GetKillSwitch(ctx(t))
	check(t, cur, err)
	var body []models.KillSwitchSegmentUpdateRequest
	for _, s := range cur.Data {
		action := "DISABLE"
		if s.KillSwitchEnabled {
			action = "ENABLE"
		}
		body = append(body, models.KillSwitchSegmentUpdateRequest{Segment: s.Segment, Action: action})
	}
	if len(body) == 0 {
		t.Skip("no kill switch segments returned")
	}
	resp, _, err := u.UpdateKillSwitch(ctx(t), body[:1])
	check(t, resp, err)
}
