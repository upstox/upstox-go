## Get smartlist options

Top traded index options, first page of 50.

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
	marketHolidaysAndTimingsApi := api.NewMarketHolidaysAndTimingsApi(client)

	// AssetType: INDEX | STOCK | COMMODITY
	// Category: TOP_TRADED, MOST_ACTIVE, OI_GAINERS, OI_LOSERS, PRICE_GAINERS, PRICE_LOSERS, IV_GAINERS, IV_LOSERS, UNDER_5000, UNDER_10000
	resp, _, err := marketHolidaysAndTimingsApi.GetSmartlistOptions(context.Background(), &api.MarketHolidaysAndTimingsApiGetSmartlistOptionsOpts{AssetType: optional.NewString("INDEX"), Category: optional.NewString("TOP_TRADED"), PageNumber: optional.NewInt32(1), PageSize: optional.NewInt32(50)})
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling MarketHolidaysAndTimingsApi.GetSmartlistOptions:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling MarketHolidaysAndTimingsApi.GetSmartlistOptions:", err)
		}
		return
	}
	if resp.Data == nil {
		fmt.Println("no data returned")
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```

## Get smartlist futures

Most active stock futures, first page of 50.

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
	marketHolidaysAndTimingsApi := api.NewMarketHolidaysAndTimingsApi(client)

	// AssetType: INDEX | STOCK | COMMODITY
	// Category: TOP_TRADED, MOST_ACTIVE, OI_GAINERS, OI_LOSERS, PRICE_GAINERS, PRICE_LOSERS, PREMIUM, DISCOUNT
	resp, _, err := marketHolidaysAndTimingsApi.GetSmartlistFutures(context.Background(), &api.MarketHolidaysAndTimingsApiGetSmartlistFuturesOpts{AssetType: optional.NewString("STOCK"), Category: optional.NewString("MOST_ACTIVE"), PageNumber: optional.NewInt32(1), PageSize: optional.NewInt32(50)})
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling MarketHolidaysAndTimingsApi.GetSmartlistFutures:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling MarketHolidaysAndTimingsApi.GetSmartlistFutures:", err)
		}
		return
	}
	if resp.Data == nil {
		fmt.Println("no data returned")
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```

## Get smartlist MTF

MTF (Margin Trade Funding) stocks enriched with live LTP, first page of 50.

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
	marketHolidaysAndTimingsApi := api.NewMarketHolidaysAndTimingsApi(client)

	resp, _, err := marketHolidaysAndTimingsApi.GetSmartlistMtf(context.Background(), &api.MarketHolidaysAndTimingsApiGetSmartlistMtfOpts{PageNumber: optional.NewInt32(1), PageSize: optional.NewInt32(50)})
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling MarketHolidaysAndTimingsApi.GetSmartlistMtf:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling MarketHolidaysAndTimingsApi.GetSmartlistMtf:", err)
		}
		return
	}
	if resp.Data == nil {
		fmt.Println("no data returned")
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```
