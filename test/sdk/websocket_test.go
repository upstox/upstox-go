//go:build sdk

package sdk

import (
	"strings"
	"testing"

	"github.com/antihax/optional"
	"github.com/upstox/upstox-go/upstox/api"
)

func checkWSURI(t *testing.T, uri string) {
	t.Helper()
	if !strings.HasPrefix(uri, "wss://") {
		t.Fatalf("authorized redirect URI %q is not wss://", redactQuery(uri))
	}
}

func TestGetMarketDataFeedAuthorize(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewWebsocketApi(c).GetMarketDataFeedAuthorize(ctx(t))
	check(t, resp, err)
	if resp.Data == nil {
		t.Fatal("no authorize data")
	}
	checkWSURI(t, resp.Data.AuthorizedRedirectUri)
}

func TestGetMarketDataFeedAuthorizeV3(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewWebsocketApi(c).GetMarketDataFeedAuthorizeV3(ctx(t))
	check(t, resp, err)
	if resp.Data == nil {
		t.Fatal("no authorize data")
	}
	checkWSURI(t, resp.Data.AuthorizedRedirectUri)
}

func TestGetPortfolioStreamFeedAuthorize(t *testing.T) {
	c := client(t)
	resp, _, err := api.NewWebsocketApi(c).GetPortfolioStreamFeedAuthorize(ctx(t), &api.WebsocketApiGetPortfolioStreamFeedAuthorizeOpts{
		UpdateTypes: optional.NewString("order,gtt_order,position,holding"),
	})
	check(t, resp, err)
	if resp.Data == nil {
		t.Fatal("no authorize data")
	}
	checkWSURI(t, resp.Data.AuthorizedRedirectUri)
}

// The raw feed endpoints redirect to the websocket URL; the client must not follow it over HTTP.
func TestGetMarketDataFeed(t *testing.T) {
	c := noRedirectClient(t)
	httpResp, err := api.NewWebsocketApi(c).GetMarketDataFeed(ctx(t))
	checkRedirect(t, httpResp, err)
}

func TestGetMarketDataFeedV3(t *testing.T) {
	c := noRedirectClient(t)
	httpResp, err := api.NewWebsocketApi(c).GetMarketDataFeedV3(ctx(t))
	checkRedirect(t, httpResp, err)
}

func TestGetPortfolioStreamFeed(t *testing.T) {
	c := noRedirectClient(t)
	httpResp, err := api.NewWebsocketApi(c).GetPortfolioStreamFeed(ctx(t), &api.WebsocketApiGetPortfolioStreamFeedOpts{
		UpdateTypes: optional.NewString("order,position"),
	})
	checkRedirect(t, httpResp, err)
}
