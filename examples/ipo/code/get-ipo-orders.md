## Get IPO orders

Lists your IPO applications.

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
	ipoApi := api.NewIPOApi(client)

	resp, _, err := ipoApi.GetIpoOrders(context.Background(), nil)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling IPOApi.GetIpoOrders:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling IPOApi.GetIpoOrders:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```

## Get IPO orders with pagination

Lists the first page of 20 IPO applications and prints the page metadata.

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
	ipoApi := api.NewIPOApi(client)

	resp, _, err := ipoApi.GetIpoOrders(context.Background(), &api.IPOApiGetIpoOrdersOpts{PageNumber: optional.NewString("1"), Records: optional.NewString("20")})
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling IPOApi.GetIpoOrders:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling IPOApi.GetIpoOrders:", err)
		}
		return
	}
	if resp.MetaData != nil && resp.MetaData.Page != nil {
		fmt.Printf("page: %+v\n", *resp.MetaData.Page)
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```

## Iterate over IPO orders and their bids

Prints each IPO order with its bids.

```go
package main

import (
	"context"
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
	ipoApi := api.NewIPOApi(client)

	resp, _, err := ipoApi.GetIpoOrders(context.Background(), &api.IPOApiGetIpoOrdersOpts{PageNumber: optional.NewString("1"), Records: optional.NewString("50")})
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling IPOApi.GetIpoOrders:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling IPOApi.GetIpoOrders:", err)
		}
		return
	}
	for _, order := range resp.Data {
		fmt.Println(order.OrderId, order.Symbol, order.OrderStatus, order.PaymentStatus)
		for _, bid := range order.Bids {
			fmt.Println("   bid:", bid.Quantity, "@", bid.Price, "=", bid.Amount)
		}
	}
}
```
