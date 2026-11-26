## Get market status for a particular exchange

Current open/close status of NSE.

```go
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/upstox/upstox-go/upstox"
	"github.com/upstox/upstox-go/upstox/api"
)

func main() {
	cfg := upstox.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox.NewAPIClient(cfg)
	marketHolidaysAndTimingsApi := api.NewMarketHolidaysAndTimingsApi(client)

	resp, _, err := marketHolidaysAndTimingsApi.GetMarketStatus(context.Background(), "NSE")
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling MarketHolidaysAndTimingsApi.GetMarketStatus:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling MarketHolidaysAndTimingsApi.GetMarketStatus:", err)
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
