## Get full market quote

V2 full quote for one instrument.

```go
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/antihax/optional"
	"github.com/upstox/upstox-go/upstox_client"
	"github.com/upstox/upstox-go/upstox_client/api"
)

func main() {
	cfg := upstox_client.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox_client.NewAPIClient(cfg)
	marketQuoteApi := api.NewMarketQuoteApi(client)

	resp, _, err := marketQuoteApi.GetFullMarketQuote(context.Background(), &api.MarketQuoteApiGetFullMarketQuoteOpts{Symbol: optional.NewString("NSE_EQ|INE669E01016")})
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling MarketQuoteApi.GetFullMarketQuote:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling MarketQuoteApi.GetFullMarketQuote:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```

## Get full market quote for multiple instrument keys

V2 full quotes for a comma-separated list of instruments.

```go
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/antihax/optional"
	"github.com/upstox/upstox-go/upstox_client"
	"github.com/upstox/upstox-go/upstox_client/api"
)

func main() {
	cfg := upstox_client.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox_client.NewAPIClient(cfg)
	marketQuoteApi := api.NewMarketQuoteApi(client)

	resp, _, err := marketQuoteApi.GetFullMarketQuote(context.Background(), &api.MarketQuoteApiGetFullMarketQuoteOpts{Symbol: optional.NewString("NSE_EQ|INE669E01016,NSE_EQ|INE848E01016")})
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling MarketQuoteApi.GetFullMarketQuote:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling MarketQuoteApi.GetFullMarketQuote:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```
