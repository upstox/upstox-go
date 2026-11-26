## Get all mutual fund orders

Lists your mutual fund orders.

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
	mutualFundApi := api.NewMutualFundApi(client)

	resp, _, err := mutualFundApi.GetMutualFundOrders(context.Background(), nil)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling MutualFundApi.GetMutualFundOrders:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling MutualFundApi.GetMutualFundOrders:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```

## Get mutual fund orders with filters

Lists only completed BUY orders.

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
	mutualFundApi := api.NewMutualFundApi(client)

	resp, _, err := mutualFundApi.GetMutualFundOrders(context.Background(), &api.MutualFundApiGetMutualFundOrdersOpts{Status: optional.NewString("complete"), TransactionType: optional.NewString("BUY")})
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling MutualFundApi.GetMutualFundOrders:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling MutualFundApi.GetMutualFundOrders:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```

## Get mutual fund orders with pagination

Lists the first page of 10 orders.

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
	mutualFundApi := api.NewMutualFundApi(client)

	resp, _, err := mutualFundApi.GetMutualFundOrders(context.Background(), &api.MutualFundApiGetMutualFundOrdersOpts{PageNumber: optional.NewInt32(1), Records: optional.NewInt32(10)})
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling MutualFundApi.GetMutualFundOrders:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling MutualFundApi.GetMutualFundOrders:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```
