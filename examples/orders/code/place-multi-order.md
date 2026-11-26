## Place a multi order

Places a single order through the multi-order endpoint.

```go
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/upstox/upstox-go/upstox"
	"github.com/upstox/upstox-go/upstox/api"
	"github.com/upstox/upstox-go/upstox/models"
)

func main() {
	cfg := upstox.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox.NewAPIClient(cfg)
	orderApi := api.NewOrderApi(client)

	body := []models.MultiOrderRequest{
		{
			Quantity:          25,
			Product:           "D",
			Validity:          "DAY",
			Price:             0,
			Tag:               "string",
			Slice:             false,
			InstrumentToken:   "NSE_FO|44166",
			OrderType:         "MARKET",
			TransactionType:   "BUY",
			DisclosedQuantity: 0,
			TriggerPrice:      0,
			IsAmo:             true,
			CorrelationId:     "1",
		},
	}
	resp, _, err := orderApi.PlaceMultiOrder(context.Background(), body, nil)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling OrderApi.PlaceMultiOrder:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling OrderApi.PlaceMultiOrder:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```

## Place Multiple BUY and SELL Orders

Places a BUY and a SELL leg in one request; each leg needs a unique correlation id.

```go
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/upstox/upstox-go/upstox"
	"github.com/upstox/upstox-go/upstox/api"
	"github.com/upstox/upstox-go/upstox/models"
)

func main() {
	cfg := upstox.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox.NewAPIClient(cfg)
	orderApi := api.NewOrderApi(client)

	body := []models.MultiOrderRequest{
		{
			Quantity:          25,
			Product:           "D",
			Validity:          "DAY",
			Price:             0,
			Tag:               "string",
			Slice:             false,
			InstrumentToken:   "NSE_FO|44166",
			OrderType:         "MARKET",
			TransactionType:   "BUY",
			DisclosedQuantity: 0,
			TriggerPrice:      0,
			IsAmo:             true,
			CorrelationId:     "1",
		},
		{
			Quantity:          25,
			Product:           "D",
			Validity:          "DAY",
			Price:             0,
			Tag:               "string",
			Slice:             false,
			InstrumentToken:   "NSE_FO|44122",
			OrderType:         "MARKET",
			TransactionType:   "SELL",
			DisclosedQuantity: 0,
			TriggerPrice:      0,
			IsAmo:             true,
			CorrelationId:     "2",
		},
	}
	resp, _, err := orderApi.PlaceMultiOrder(context.Background(), body, nil)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling OrderApi.PlaceMultiOrder:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling OrderApi.PlaceMultiOrder:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```

## Place Multiple Orders with Auto Slicing enabled

Lets the exchange freeze-quantity limit split the order automatically.

```go
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/upstox/upstox-go/upstox"
	"github.com/upstox/upstox-go/upstox/api"
	"github.com/upstox/upstox-go/upstox/models"
)

func main() {
	cfg := upstox.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox.NewAPIClient(cfg)
	orderApi := api.NewOrderApi(client)

	body := []models.MultiOrderRequest{
		{
			Quantity:          25,
			Product:           "D",
			Validity:          "DAY",
			Price:             0,
			Tag:               "string",
			Slice:             true,
			InstrumentToken:   "NSE_FO|44166",
			OrderType:         "MARKET",
			TransactionType:   "BUY",
			DisclosedQuantity: 0,
			TriggerPrice:      0,
			IsAmo:             true,
			CorrelationId:     "1",
		},
	}
	resp, _, err := orderApi.PlaceMultiOrder(context.Background(), body, nil)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling OrderApi.PlaceMultiOrder:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling OrderApi.PlaceMultiOrder:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```
