## Get ohlc (open, high, low, close) market quotes

V2 daily OHLC of one instrument.

```go
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/antihax/optional"
	"github.com/upstox/upstox-go/upstox"
	"github.com/upstox/upstox-go/upstox/api"
)

func main() {
	cfg := upstox.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox.NewAPIClient(cfg)
	marketQuoteApi := api.NewMarketQuoteApi(client)

	interval := "1d"
	resp, _, err := marketQuoteApi.GetMarketQuoteOHLC(context.Background(), interval, &api.MarketQuoteApiGetMarketQuoteOHLCOpts{Symbol: optional.NewString("NSE_EQ|INE669E01016")})
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling MarketQuoteApi.GetMarketQuoteOHLC:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling MarketQuoteApi.GetMarketQuoteOHLC:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```

## Get ohlc (open, high, low, close) market quotes for multiple instrument keys

V2 daily OHLC for a comma-separated list of instruments.

```go
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/antihax/optional"
	"github.com/upstox/upstox-go/upstox"
	"github.com/upstox/upstox-go/upstox/api"
)

func main() {
	cfg := upstox.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox.NewAPIClient(cfg)
	marketQuoteApi := api.NewMarketQuoteApi(client)

	interval := "1d"
	resp, _, err := marketQuoteApi.GetMarketQuoteOHLC(context.Background(), interval, &api.MarketQuoteApiGetMarketQuoteOHLCOpts{Symbol: optional.NewString("NSE_EQ|INE669E01016,NSE_EQ|INE848E01016")})
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling MarketQuoteApi.GetMarketQuoteOHLC:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling MarketQuoteApi.GetMarketQuoteOHLC:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```
