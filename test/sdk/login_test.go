//go:build sdk

package sdk

import (
	"testing"

	"github.com/antihax/optional"
	"github.com/upstox/upstox-go/upstox/api"
	"github.com/upstox/upstox-go/upstox/models"
)

func TestAuthorize(t *testing.T) {
	c := noRedirectClient(t)
	env := requireEnv(t, "UPSTOX_CLIENT_ID", "UPSTOX_REDIRECT_URI")
	httpResp, err := api.NewLoginApi(c).Authorize(ctx(t), env[0], env[1], &api.LoginApiAuthorizeOpts{
		State: optional.NewString("sdk-go-test"),
	})
	checkRedirect(t, httpResp, err)
}

// TestToken exchanges a one-time auth code; each code works once, so supply a fresh one per run.
func TestToken(t *testing.T) {
	c := client(t)
	env := requireEnv(t, "UPSTOX_CLIENT_ID", "UPSTOX_CLIENT_SECRET", "UPSTOX_REDIRECT_URI", "UPSTOX_AUTH_CODE")
	resp, _, err := api.NewLoginApi(c).Token(ctx(t), env[3], env[0], env[1], env[2], "authorization_code", false)
	check(t, resp, err)
	if resp.AccessToken == "" {
		t.Fatal("token exchange returned no access token")
	}
}

func TestTokenInvalidCode(t *testing.T) {
	c := client(t)
	_, httpResp, err := api.NewLoginApi(c).Token(ctx(t), "invalid-code", "invalid-client", "invalid-secret",
		"https://example.com/callback", "authorization_code", false)
	checkRejected(t, httpResp, err)
}

func TestInitTokenRequestInvalidClient(t *testing.T) {
	c := client(t)
	body := models.IndieUserTokenRequest{ClientSecret: "invalid-secret"}
	_, httpResp, err := api.NewLoginApi(c).InitTokenRequestForIndieUser(ctx(t), body, "00000000-0000-0000-0000-000000000000")
	checkRejected(t, httpResp, err)
}

// TestGenerateTokenViaTotpAndPinInvalid sends no X-User-Id, so it cannot count against a real user's PIN attempts.
func TestGenerateTokenViaTotpAndPinInvalid(t *testing.T) {
	c := client(t)
	body := models.TotpPinTokenRequest{Pin: "000000", Totp: "000000"}
	_, httpResp, err := api.NewLoginApi(c).GenerateTokenViaTotpAndPin(ctx(t), body, "invalid-api-key", nil)
	checkRejected(t, httpResp, err)
}

// TestLogout invalidates UPSTOX_ACCESS_TOKEN; run it on its own (-run TestLogout).
func TestLogout(t *testing.T) {
	c := client(t)
	requireFlag(t, "UPSTOX_RUN_LOGOUT_TEST")
	resp, _, err := api.NewLoginApi(c).Logout(ctx(t))
	check(t, resp, err)
}
