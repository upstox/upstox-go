## Get trades for an order number

Trades executed against one order.

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
	orderApi := api.NewOrderApi(client)

	orderId := "{your_order_id}"
	resp, _, err := orderApi.GetTradesByOrder(context.Background(), orderId)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling OrderApi.GetTradesByOrder:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling OrderApi.GetTradesByOrder:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```
