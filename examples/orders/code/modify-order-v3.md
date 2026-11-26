## Modify a delivery order

Changes an open order to a LIMIT at 9.12 through the V3 (HFT) endpoint.

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
	orderV3Api := api.NewOrderV3Api(client)

	body := models.ModifyOrderRequest{
		Quantity:          1,
		Validity:          "DAY",
		Price:             9.12,
		OrderId:           "25030310405859",
		OrderType:         "LIMIT",
		DisclosedQuantity: 0,
		TriggerPrice:      0,
	}
	resp, _, err := orderV3Api.ModifyOrderV3(context.Background(), body, nil)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling OrderV3Api.ModifyOrderV3:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling OrderV3Api.ModifyOrderV3:", err)
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
