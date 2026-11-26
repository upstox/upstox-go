## Get mutual fund order by id

Fetches one mutual fund order.

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
	mutualFundApi := api.NewMutualFundApi(client)

	orderId := "{your_mf_order_id}"
	resp, _, err := mutualFundApi.GetMutualFundOrder(context.Background(), orderId)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling MutualFundApi.GetMutualFundOrder:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling MutualFundApi.GetMutualFundOrder:", err)
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
