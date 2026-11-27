## Estimate charges for a market order

Charge estimate for buying 1 AAPL share at market.

```go
package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/upstox/upstox-go/upstox_client"
	"github.com/upstox/upstox-go/upstox_client/api"
	"github.com/upstox/upstox-go/upstox_client/models"
)

func main() {
	cfg := upstox_client.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox_client.NewAPIClient(cfg)
	usStocksOrderApi := api.NewUSStocksOrderApi(client)

	// Nothing is placed
	body := models.UsOrderEstimateRequest{
		Quantity:        1,
		InstrumentKey:   "US_EQ|AAPL",
		TransactionType: "BUY",
		OrderType:       "MARKET",
	}
	resp, _, err := usStocksOrderApi.EstimateUsOrder(context.Background(), body)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling USStocksOrderApi.EstimateUsOrder:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling USStocksOrderApi.EstimateUsOrder:", err)
		}
		return
	}
	if resp.Data == nil {
		fmt.Println("no data returned")
		return
	}
	fmt.Println("total charges:", resp.Data.TotalCharges)
	fmt.Println("estimated total:", resp.Data.EstimatedTotal)
}
```

## Estimate charges for a limit order

Charge estimate for selling 0.5 AAPL shares at 150 USD.

```go
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/upstox/upstox-go/upstox_client"
	"github.com/upstox/upstox-go/upstox_client/api"
	"github.com/upstox/upstox-go/upstox_client/models"
)

func main() {
	cfg := upstox_client.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox_client.NewAPIClient(cfg)
	usStocksOrderApi := api.NewUSStocksOrderApi(client)

	body := models.UsOrderEstimateRequest{
		Quantity:        0.5,
		Price:           150.0,
		InstrumentKey:   "US_EQ|AAPL",
		TransactionType: "SELL",
		OrderType:       "LIMIT",
	}
	resp, _, err := usStocksOrderApi.EstimateUsOrder(context.Background(), body)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling USStocksOrderApi.EstimateUsOrder:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling USStocksOrderApi.EstimateUsOrder:", err)
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

## Calculate margin for multiple orders

Total margin required for a basket of US orders.

```go
package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/upstox/upstox-go/upstox_client"
	"github.com/upstox/upstox-go/upstox_client/api"
	"github.com/upstox/upstox-go/upstox_client/models"
)

func main() {
	cfg := upstox_client.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox_client.NewAPIClient(cfg)
	usStocksChargesApi := api.NewUSStocksChargesApi(client)

	// Up to 20 orders; nothing is placed
	body := models.UsMarginRequest{Instruments: []models.UsOrderEstimateRequest{
		{Quantity: 1, InstrumentKey: "US_EQ|AAPL", TransactionType: "BUY", OrderType: "MARKET"},
		{Quantity: 2, Price: 100.0, InstrumentKey: "US_EQ|MSFT", TransactionType: "BUY", OrderType: "LIMIT"},
	}}
	resp, _, err := usStocksChargesApi.CalculateUsMargin(context.Background(), body)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling USStocksChargesApi.CalculateUsMargin:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling USStocksChargesApi.CalculateUsMargin:", err)
		}
		return
	}
	if resp.Data == nil {
		fmt.Println("no data returned")
		return
	}
	fmt.Println("required margin:", resp.Data.RequiredMargin)
	fmt.Println("sufficient:", resp.Data.Sufficient, "shortfall:", resp.Data.Shortfall)
}
```
