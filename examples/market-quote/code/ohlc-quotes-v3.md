## Get ohlc (open, high, low, close) market quotes

Daily OHLC of one instrument.

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
	marketQuoteV3Api := api.NewMarketQuoteV3Api(client)

	// For a single instrument
	resp, _, err := marketQuoteV3Api.GetMarketQuoteOHLC(context.Background(), "1d", &api.MarketQuoteV3ApiGetMarketQuoteOHLCOpts{InstrumentKey: optional.NewString("NSE_EQ|INE669E01016")})
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling MarketQuoteV3Api.GetMarketQuoteOHLC:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling MarketQuoteV3Api.GetMarketQuoteOHLC:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```

## Get ohlc (open, high, low, close) market quotes for multiple instrument keys

Daily OHLC for a comma-separated list of instruments.

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
	marketQuoteV3Api := api.NewMarketQuoteV3Api(client)

	// For multiple instruments
	resp, _, err := marketQuoteV3Api.GetMarketQuoteOHLC(context.Background(), "1d", &api.MarketQuoteV3ApiGetMarketQuoteOHLCOpts{InstrumentKey: optional.NewString("NSE_EQ|INE669E01016,NSE_EQ|INE848E01016")})
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling MarketQuoteV3Api.GetMarketQuoteOHLC:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling MarketQuoteV3Api.GetMarketQuoteOHLC:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```
