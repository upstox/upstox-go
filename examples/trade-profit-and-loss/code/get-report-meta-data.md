## Get report meta data for equity segment

Row count and page size of the equity P&L report.

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
	tradeProfitAndLossApi := api.NewTradeProfitAndLossApi(client)

	// Financial year is the last 2 digits of both years, e.g. 2324 for 2023-24
	financialYear := "2324"
	// Dates are dd-mm-yyyy and must fall inside the financial year
	resp, _, err := tradeProfitAndLossApi.GetTradeWiseProfitAndLossMetaData(context.Background(), "EQ", financialYear, &api.TradeProfitAndLossApiGetTradeWiseProfitAndLossMetaDataOpts{FromDate: optional.NewString("02-04-2023"), ToDate: optional.NewString("20-03-2024")})
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling TradeProfitAndLossApi.GetTradeWiseProfitAndLossMetaData:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling TradeProfitAndLossApi.GetTradeWiseProfitAndLossMetaData:", err)
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
