## Get profit loss report for futures and options segment

Trade-wise P&L for F&O, first page of 4 rows.

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
	tradeProfitAndLossApi := api.NewTradeProfitAndLossApi(client)

	// Financial year is the last 2 digits of both years, e.g. 2324 for 2023-24
	financialYear := "2324"
	// Dates are dd-mm-yyyy and must fall inside the financial year
	resp, _, err := tradeProfitAndLossApi.GetTradeWiseProfitAndLossData(context.Background(), "FO", financialYear, 1, 4, &api.TradeProfitAndLossApiGetTradeWiseProfitAndLossDataOpts{FromDate: optional.NewString("02-04-2023"), ToDate: optional.NewString("20-03-2024")})
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling TradeProfitAndLossApi.GetTradeWiseProfitAndLossData:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling TradeProfitAndLossApi.GetTradeWiseProfitAndLossData:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```
