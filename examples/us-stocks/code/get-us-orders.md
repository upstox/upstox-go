## Get US order book

Lists today's US orders in every state.

```go
package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/upstox/upstox-go/upstox"
	"github.com/upstox/upstox-go/upstox/api"
)

func main() {
	cfg := upstox.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox.NewAPIClient(cfg)
	usStocksOrderApi := api.NewUSStocksOrderApi(client)

	// Today's orders in every state
	resp, _, err := usStocksOrderApi.GetUsOrderBook(context.Background(), nil)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling USStocksOrderApi.GetUsOrderBook:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling USStocksOrderApi.GetUsOrderBook:", err)
		}
		return
	}
	for _, order := range resp.Data {
		fmt.Println(order.OrderId, order.TradingSymbol, order.Status, order.FilledQuantity)
	}
}
```

## Get US order book filtered by tag

Lists today's US orders placed with a given tag.

```go
package main

import (
	"context"
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
	usStocksOrderApi := api.NewUSStocksOrderApi(client)

	// Exact, case-sensitive match on the tag set when placing the order
	resp, _, err := usStocksOrderApi.GetUsOrderBook(context.Background(), &api.USStocksOrderApiGetUsOrderBookOpts{Tag: optional.NewString("my-us-order")})
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling USStocksOrderApi.GetUsOrderBook:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling USStocksOrderApi.GetUsOrderBook:", err)
		}
		return
	}
	for _, order := range resp.Data {
		fmt.Println(order.OrderId, order.Tag, order.Status)
	}
}
```

## Get US order details by order id

Fetches the status of one US order.

```go
package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/upstox/upstox-go/upstox"
	"github.com/upstox/upstox-go/upstox/api"
)

func main() {
	cfg := upstox.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox.NewAPIClient(cfg)
	usStocksOrderApi := api.NewUSStocksOrderApi(client)

	resp, _, err := usStocksOrderApi.GetUsOrderDetails(context.Background(), "{us_order_id}")
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling USStocksOrderApi.GetUsOrderDetails:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling USStocksOrderApi.GetUsOrderDetails:", err)
		}
		return
	}
	if resp.Data == nil {
		fmt.Println("no data returned")
		return
	}
	fmt.Println(resp.Data.Status, resp.Data.StatusMessage)
}
```
