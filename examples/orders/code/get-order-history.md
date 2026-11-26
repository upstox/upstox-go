## Get order history for an order number

Every state transition of one order.

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
	orderApi := api.NewOrderApi(client)

	orderId := "240112010371054"
	resp, _, err := orderApi.GetOrderDetails(context.Background(), &api.OrderApiGetOrderDetailsOpts{OrderId: optional.NewString(orderId)})
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling OrderApi.GetOrderDetails:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling OrderApi.GetOrderDetails:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```
