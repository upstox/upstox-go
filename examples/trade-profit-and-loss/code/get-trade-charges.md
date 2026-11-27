## Get trade charges for equity segment

Brokerage, taxes and other charges on equity trades.

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
	resp, _, err := tradeProfitAndLossApi.GetProfitAndLossCharges(context.Background(), "EQ", financialYear, &api.TradeProfitAndLossApiGetProfitAndLossChargesOpts{FromDate: optional.NewString("02-04-2023"), ToDate: optional.NewString("20-03-2024")})
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling TradeProfitAndLossApi.GetProfitAndLossCharges:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling TradeProfitAndLossApi.GetProfitAndLossCharges:", err)
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

## Get trade charges for futures and options segment

Brokerage, taxes and other charges on F&O trades.

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
	resp, _, err := tradeProfitAndLossApi.GetProfitAndLossCharges(context.Background(), "FO", financialYear, &api.TradeProfitAndLossApiGetProfitAndLossChargesOpts{FromDate: optional.NewString("02-04-2023"), ToDate: optional.NewString("20-03-2024")})
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling TradeProfitAndLossApi.GetProfitAndLossCharges:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling TradeProfitAndLossApi.GetProfitAndLossCharges:", err)
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
