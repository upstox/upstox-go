## Cancel all the open orders

Cancels every open order.

```go
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/upstox/upstox-go/upstox_client"
	"github.com/upstox/upstox-go/upstox_client/api"
)

func main() {
	cfg := upstox_client.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox_client.NewAPIClient(cfg)
	orderApi := api.NewOrderApi(client)

	resp, _, err := orderApi.CancelMultiOrder(context.Background(), nil)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling OrderApi.CancelMultiOrder:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling OrderApi.CancelMultiOrder:", err)
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

## Cancel all the open orders for a given segment

Cancels open orders in the NSE_FO segment only.

```go
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/antihax/optional"
	"github.com/upstox/upstox-go/upstox_client"
	"github.com/upstox/upstox-go/upstox_client/api"
)

func main() {
	cfg := upstox_client.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox_client.NewAPIClient(cfg)
	orderApi := api.NewOrderApi(client)

	resp, _, err := orderApi.CancelMultiOrder(context.Background(), &api.OrderApiCancelMultiOrderOpts{Segment: optional.NewString("NSE_FO")})
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling OrderApi.CancelMultiOrder:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling OrderApi.CancelMultiOrder:", err)
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

## Cancel all the open orders for a given tag

Cancels open orders placed with tag `xyz`.

```go
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/antihax/optional"
	"github.com/upstox/upstox-go/upstox_client"
	"github.com/upstox/upstox-go/upstox_client/api"
)

func main() {
	cfg := upstox_client.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox_client.NewAPIClient(cfg)
	orderApi := api.NewOrderApi(client)

	resp, _, err := orderApi.CancelMultiOrder(context.Background(), &api.OrderApiCancelMultiOrderOpts{Tag: optional.NewString("xyz")})
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling OrderApi.CancelMultiOrder:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling OrderApi.CancelMultiOrder:", err)
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
