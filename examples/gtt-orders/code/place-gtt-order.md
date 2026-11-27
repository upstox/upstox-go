## Place Single Leg GTT Order

Places a GTT order with only an entry trigger.

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

	body := models.GttPlaceOrderRequest{
		Type_:           "SINGLE",
		InstrumentToken: "NSE_EQ|INE669E01016",
		Product:         "D",
		Quantity:        1,
		Rules: []models.GttRule{
			{Strategy: "ENTRY", TriggerType: "ABOVE", TriggerPrice: 7},
		},
		TransactionType: "BUY",
	}
	resp, _, err := orderV3Api.PlaceGTTOrder(context.Background(), body, nil)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling OrderV3Api.PlaceGTTOrder:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling OrderV3Api.PlaceGTTOrder:", err)
		}
		return
	}
	if resp.Data == nil {
		fmt.Println("no data returned")
		return
	}
	fmt.Println("GTT order response:")
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```

## Place Multiple Leg GTT Order

Places a GTT order with entry, target and stop-loss legs.

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

	body := models.GttPlaceOrderRequest{
		Type_:           "MULTIPLE",
		InstrumentToken: "NSE_EQ|INE669E01016",
		Product:         "D",
		Quantity:        1,
		Rules: []models.GttRule{
			{Strategy: "ENTRY", TriggerType: "ABOVE", TriggerPrice: 7},
			{Strategy: "TARGET", TriggerType: "IMMEDIATE", TriggerPrice: 9},
			{Strategy: "STOPLOSS", TriggerType: "IMMEDIATE", TriggerPrice: 5},
		},
		TransactionType: "BUY",
	}
	resp, _, err := orderV3Api.PlaceGTTOrder(context.Background(), body, nil)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling OrderV3Api.PlaceGTTOrder:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling OrderV3Api.PlaceGTTOrder:", err)
		}
		return
	}
	if resp.Data == nil {
		fmt.Println("no data returned")
		return
	}
	fmt.Println("GTT order response:")
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```
