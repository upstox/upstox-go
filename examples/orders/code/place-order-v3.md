## Place an order with slicing enabled

A large F&O market order that the V3 endpoint slices into freeze-limit chunks.

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

	body := models.PlaceOrderV3Request{
		Quantity:          4000,
		Product:           "D",
		Validity:          "DAY",
		Price:             0,
		Tag:               "string",
		InstrumentToken:   "NSE_FO|43919",
		OrderType:         "MARKET",
		TransactionType:   "BUY",
		DisclosedQuantity: 0,
		TriggerPrice:      0.0,
		IsAmo:             false,
		Slice:             true,
	}
	resp, _, err := orderV3Api.PlaceOrderV3(context.Background(), body, nil)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling OrderV3Api.PlaceOrderV3:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling OrderV3Api.PlaceOrderV3:", err)
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

## Place an order with slicing disabled

An F&O market order placed as a single order.

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

	body := models.PlaceOrderV3Request{
		Quantity:          75,
		Product:           "D",
		Validity:          "DAY",
		Price:             0,
		Tag:               "string",
		InstrumentToken:   "NSE_FO|43919",
		OrderType:         "MARKET",
		TransactionType:   "BUY",
		DisclosedQuantity: 0,
		TriggerPrice:      0.0,
		IsAmo:             false,
		Slice:             false,
	}
	resp, _, err := orderV3Api.PlaceOrderV3(context.Background(), body, nil)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling OrderV3Api.PlaceOrderV3:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling OrderV3Api.PlaceOrderV3:", err)
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
