## Get ltp (last traded price) market quotes

V2 last traded price of one instrument.

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

	resp, _, err := marketQuoteApi.Ltp(context.Background(), &api.MarketQuoteApiLtpOpts{Symbol: optional.NewString("NSE_EQ|INE669E01016")})
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling MarketQuoteApi.Ltp:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling MarketQuoteApi.Ltp:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```

## Get ltp (last traded price) market quotes for multiple instruments keys

V2 last traded prices for a comma-separated list of instruments.

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

	resp, _, err := marketQuoteApi.Ltp(context.Background(), &api.MarketQuoteApiLtpOpts{Symbol: optional.NewString("NSE_EQ|INE669E01016,NSE_EQ|INE848E01016")})
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling MarketQuoteApi.Ltp:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling MarketQuoteApi.Ltp:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```
