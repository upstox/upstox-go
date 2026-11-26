## Get ltp (last traded price) market quotes

Last traded price of one instrument.

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
	resp, _, err := marketQuoteV3Api.GetLtp(context.Background(), &api.MarketQuoteV3ApiGetLtpOpts{InstrumentKey: optional.NewString("NSE_EQ|INE848E01016")})
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling MarketQuoteV3Api.GetLtp:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling MarketQuoteV3Api.GetLtp:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```

## Get ltp (last traded price) market quotes for multiple instruments keys

Last traded prices for a comma-separated list of instruments.

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
	resp, _, err := marketQuoteV3Api.GetLtp(context.Background(), &api.MarketQuoteV3ApiGetLtpOpts{InstrumentKey: optional.NewString("NSE_EQ|INE848E01016,NSE_EQ|INE669E01016")})
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling MarketQuoteV3Api.GetLtp:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling MarketQuoteV3Api.GetLtp:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```
