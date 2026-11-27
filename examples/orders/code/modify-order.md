## Modify a delivery order

Changes an open order to a MARKET order for 2 units through the V2 endpoint.

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
	orderApi := api.NewOrderApi(client)

	body := models.ModifyOrderRequest{
		Quantity:          2,
		Validity:          "DAY",
		Price:             0,
		OrderId:           "231222010275930",
		OrderType:         "MARKET",
		DisclosedQuantity: 0,
		TriggerPrice:      0,
	}
	resp, _, err := orderApi.ModifyOrder(context.Background(), body, nil)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling OrderApi.ModifyOrder:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling OrderApi.ModifyOrder:", err)
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
