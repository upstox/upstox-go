## Place a market order by quantity

Buys 1.5 shares of AAPL at market.

```go
package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/upstox/upstox-go/upstox"
	"github.com/upstox/upstox-go/upstox/api"
	"github.com/upstox/upstox-go/upstox/models"
)

func main() {
	cfg := upstox.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox.NewAPIClient(cfg)
	usStocksOrderApi := api.NewUSStocksOrderApi(client)

	// Fractional quantities are allowed only on fractionable stocks
	body := models.UsPlaceOrderRequest{
		Quantity:        1.5,
		Validity:        "DAY",
		InstrumentKey:   "US_EQ|AAPL",
		TransactionType: "BUY",
		OrderType:       "MARKET",
	}
	resp, _, err := usStocksOrderApi.PlaceUsOrder(context.Background(), body)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling USStocksOrderApi.PlaceUsOrder:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling USStocksOrderApi.PlaceUsOrder:", err)
		}
		return
	}
	if resp.Data == nil {
		fmt.Println("no data returned")
		return
	}
	fmt.Println("order id:", resp.Data.OrderId)
}
```

## Place a market order by USD amount

Buys 50 USD worth of AAPL at market.

```go
package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/upstox/upstox-go/upstox"
	"github.com/upstox/upstox-go/upstox/api"
	"github.com/upstox/upstox-go/upstox/models"
)

func main() {
	cfg := upstox.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox.NewAPIClient(cfg)
	usStocksOrderApi := api.NewUSStocksOrderApi(client)

	// Size the order by USD amount instead of quantity
	body := models.UsPlaceOrderRequest{
		Amount:          50.0,
		Validity:        "DAY",
		InstrumentKey:   "US_EQ|AAPL",
		TransactionType: "BUY",
		OrderType:       "MARKET",
	}
	resp, _, err := usStocksOrderApi.PlaceUsOrder(context.Background(), body)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling USStocksOrderApi.PlaceUsOrder:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling USStocksOrderApi.PlaceUsOrder:", err)
		}
		return
	}
	if resp.Data == nil {
		fmt.Println("no data returned")
		return
	}
	fmt.Println("order id:", resp.Data.OrderId)
}
```

## Place a limit order

Buys 2 shares of AAPL with a limit price of 150.25 USD.

```go
package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/upstox/upstox-go/upstox"
	"github.com/upstox/upstox-go/upstox/api"
	"github.com/upstox/upstox-go/upstox/models"
)

func main() {
	cfg := upstox.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox.NewAPIClient(cfg)
	usStocksOrderApi := api.NewUSStocksOrderApi(client)

	body := models.UsPlaceOrderRequest{
		Quantity:        2,
		Price:           150.25,
		Validity:        "DAY",
		Tag:             "my-us-order",
		InstrumentKey:   "US_EQ|AAPL",
		TransactionType: "BUY",
		OrderType:       "LIMIT",
	}
	resp, _, err := usStocksOrderApi.PlaceUsOrder(context.Background(), body)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling USStocksOrderApi.PlaceUsOrder:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling USStocksOrderApi.PlaceUsOrder:", err)
		}
		return
	}
	if resp.Data == nil {
		fmt.Println("no data returned")
		return
	}
	fmt.Println("order id:", resp.Data.OrderId)
}
```
