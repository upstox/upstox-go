## Modify Single Leg GTT Order

Changes the entry trigger and quantity of a single-leg GTT order.

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
	orderV3Api := api.NewOrderV3Api(client)

	body := models.GttModifyOrderRequest{
		Type_:      "SINGLE",
		GttOrderId: "GTT-C25270200137952",
		Rules: []models.GttRule{
			{Strategy: "ENTRY", TriggerType: "ABOVE", TriggerPrice: 7.3},
		},
		Quantity: 1,
	}
	resp, _, err := orderV3Api.ModifyGTTOrder(context.Background(), body, nil)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling OrderV3Api.ModifyGTTOrder:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling OrderV3Api.ModifyGTTOrder:", err)
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

## Modify Multiple Leg GTT Order

Changes the entry, target and stop-loss legs of a multi-leg GTT order.

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
	orderV3Api := api.NewOrderV3Api(client)

	body := models.GttModifyOrderRequest{
		Type_:      "MULTIPLE",
		GttOrderId: "GTT-C25280200137522",
		Rules: []models.GttRule{
			{Strategy: "ENTRY", TriggerType: "ABOVE", TriggerPrice: 7.3},
			{Strategy: "TARGET", TriggerType: "IMMEDIATE", TriggerPrice: 7.64},
			{Strategy: "STOPLOSS", TriggerType: "IMMEDIATE", TriggerPrice: 7.1},
		},
		Quantity: 1,
	}
	resp, _, err := orderV3Api.ModifyGTTOrder(context.Background(), body, nil)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling OrderV3Api.ModifyGTTOrder:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling OrderV3Api.ModifyGTTOrder:", err)
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
