## Cancel an open order

Cancels an open US stock order.

```go
package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/upstox/upstox-go/upstox_client"
	"github.com/upstox/upstox-go/upstox_client/api"
)

func main() {
	cfg := upstox_client.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox_client.NewAPIClient(cfg)
	usStocksOrderApi := api.NewUSStocksOrderApi(client)

	resp, _, err := usStocksOrderApi.CancelUsOrder(context.Background(), "{us_order_id}")
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
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
