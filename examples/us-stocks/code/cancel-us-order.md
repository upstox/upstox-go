## Cancel an open order

Cancels an open US stock order.

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

	resp, _, err := usStocksOrderApi.CancelUsOrder(context.Background(), "{us_order_id}")
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling USStocksOrderApi.CancelUsOrder:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling USStocksOrderApi.CancelUsOrder:", err)
		}
		return
	}
	if resp.Data == nil {
		fmt.Println("no data returned")
		return
	}
	fmt.Println("cancelled order id:", resp.Data.OrderId)
}
```
