## Get GTT Order Details

Fetches the details of one GTT order.

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
	orderV3Api := api.NewOrderV3Api(client)

	resp, _, err := orderV3Api.GetGttOrderDetails(context.Background(), &api.OrderV3ApiGetGttOrderDetailsOpts{GttOrderId: optional.NewString("GTT-C25030300128840")})
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling OrderV3Api.GetGttOrderDetails:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling OrderV3Api.GetGttOrderDetails:", err)
		}
		return
	}
	fmt.Println("GTT order details:")
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```
