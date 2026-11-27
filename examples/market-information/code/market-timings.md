## Get market timings of a date

Exchange trading sessions on one date.

```go
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/upstox/upstox-go/upstox_client"
	"github.com/upstox/upstox-go/upstox_client/api"
)

func main() {
	cfg := upstox_client.NewConfiguration()
	client := upstox_client.NewAPIClient(cfg)
	marketHolidaysAndTimingsApi := api.NewMarketHolidaysAndTimingsApi(client)

	resp, _, err := marketHolidaysAndTimingsApi.GetExchangeTimings(context.Background(), "2024-01-22")
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling MarketHolidaysAndTimingsApi.GetExchangeTimings:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling MarketHolidaysAndTimingsApi.GetExchangeTimings:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```
