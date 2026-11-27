## Cancel GTT Order

Cancels a GTT order by its GTT order id.

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

	body := models.GttCancelOrderRequest{GttOrderId: "GTT-C250303008840"}
	resp, _, err := orderV3Api.CancelGTTOrder(context.Background(), body, nil)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling OrderV3Api.CancelGTTOrder:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling OrderV3Api.CancelGTTOrder:", err)
		}
		return
	}
	if resp.Data == nil {
		fmt.Println("no data returned")
		return
	}
	fmt.Println("GTT order canceled:")
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```
