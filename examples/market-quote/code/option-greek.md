## Get Option Greek fields

Delta, gamma, theta, vega and IV of one option contract.

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
	marketQuoteV3Api := api.NewMarketQuoteV3Api(client)

	// For a single instrument
	resp, _, err := marketQuoteV3Api.GetMarketQuoteOptionGreek(context.Background(), &api.MarketQuoteV3ApiGetMarketQuoteOptionGreekOpts{InstrumentKey: optional.NewString("NSE_FO|43885")})
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling MarketQuoteV3Api.GetMarketQuoteOptionGreek:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling MarketQuoteV3Api.GetMarketQuoteOptionGreek:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```

## Get Option Greek fields for multiple instruments keys

Greeks for a comma-separated list of option contracts.

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
	marketQuoteV3Api := api.NewMarketQuoteV3Api(client)

	// For multiple instruments
	resp, _, err := marketQuoteV3Api.GetMarketQuoteOptionGreek(context.Background(), &api.MarketQuoteV3ApiGetMarketQuoteOptionGreekOpts{InstrumentKey: optional.NewString("NSE_FO|38604,NSE_FO|49210")})
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling MarketQuoteV3Api.GetMarketQuoteOptionGreek:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling MarketQuoteV3Api.GetMarketQuoteOptionGreek:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```
