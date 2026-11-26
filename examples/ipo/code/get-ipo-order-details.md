## Get IPO order details by order id

Fetches one IPO order.

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
	ipoApi := api.NewIPOApi(client)

	// The order id comes from ApplyForIpo or GetIpoOrders
	orderId := "{ipo_order_id}"
	resp, _, err := ipoApi.GetIpoOrderById(context.Background(), orderId)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling IPOApi.GetIpoOrderById:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling IPOApi.GetIpoOrderById:", err)
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

## Read the order and payment status of an IPO order

Prints the status fields and every bid of one IPO order.

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
	ipoApi := api.NewIPOApi(client)

	orderId := "{ipo_order_id}"
	resp, _, err := ipoApi.GetIpoOrderById(context.Background(), orderId)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling IPOApi.GetIpoOrderById:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling IPOApi.GetIpoOrderById:", err)
		}
		return
	}
	order := resp.Data
	if order == nil {
		fmt.Println("no data returned")
		return
	}
	fmt.Println("symbol:         ", order.Symbol)
	fmt.Println("exchange:       ", order.Exchange)
	fmt.Println("order status:   ", order.OrderStatus)
	fmt.Println("payment status: ", order.PaymentStatus)
	fmt.Println("units allotted: ", order.UnitsAllotted)
	for _, bid := range order.Bids {
		fmt.Println("   bid:", bid.Quantity, "@", bid.Price, "=", bid.Amount)
	}
}
```

## Look up the most recent IPO order

Reads the latest order id from the order list, then fetches its details.

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
	ipoApi := api.NewIPOApi(client)

	orders, _, err := ipoApi.GetIpoOrders(context.Background(), &api.IPOApiGetIpoOrdersOpts{PageNumber: optional.NewString("1"), Records: optional.NewString("1")})
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling IPOApi.GetIpoOrders:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling IPOApi.GetIpoOrders:", err)
		}
		return
	}
	if len(orders.Data) == 0 {
		fmt.Println("no IPO orders found")
		return
	}

	resp, _, err := ipoApi.GetIpoOrderById(context.Background(), orders.Data[0].OrderId)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling IPOApi.GetIpoOrderById:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling IPOApi.GetIpoOrderById:", err)
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
