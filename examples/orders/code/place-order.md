## Place a delivery market order

Delivery BUY at market price.

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

	body := models.PlaceOrderRequest{
		Quantity:          1,
		Product:           "D",
		Validity:          "DAY",
		Price:             0.0,
		Tag:               "string",
		InstrumentToken:   "NSE_EQ|INE528G01035",
		OrderType:         "MARKET",
		TransactionType:   "BUY",
		DisclosedQuantity: 0,
		TriggerPrice:      0.0,
		IsAmo:             false,
	}
	resp, _, err := orderApi.PlaceOrder(context.Background(), body, nil)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling OrderApi.PlaceOrder:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling OrderApi.PlaceOrder:", err)
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

## Place a delivery limit order

Delivery BUY with a limit price of 20.0.

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

	body := models.PlaceOrderRequest{
		Quantity:          1,
		Product:           "D",
		Validity:          "DAY",
		Price:             20.0,
		Tag:               "string",
		InstrumentToken:   "NSE_EQ|INE528G01035",
		OrderType:         "LIMIT",
		TransactionType:   "BUY",
		DisclosedQuantity: 0,
		TriggerPrice:      20.1,
		IsAmo:             false,
	}
	resp, _, err := orderApi.PlaceOrder(context.Background(), body, nil)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling OrderApi.PlaceOrder:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling OrderApi.PlaceOrder:", err)
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

## Place a delivery stop-loss order

Delivery stop-loss limit BUY triggered at 19.5.

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

	body := models.PlaceOrderRequest{
		Quantity:          1,
		Product:           "D",
		Validity:          "DAY",
		Price:             20,
		Tag:               "string",
		InstrumentToken:   "NSE_EQ|INE528G01035",
		OrderType:         "SL",
		TransactionType:   "BUY",
		DisclosedQuantity: 0,
		TriggerPrice:      19.5,
		IsAmo:             false,
	}
	resp, _, err := orderApi.PlaceOrder(context.Background(), body, nil)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling OrderApi.PlaceOrder:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling OrderApi.PlaceOrder:", err)
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

## Place a delivery stop-loss order market

Delivery stop-loss market BUY triggered at 21.5.

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

	body := models.PlaceOrderRequest{
		Quantity:          1,
		Product:           "D",
		Validity:          "DAY",
		Price:             0.0,
		Tag:               "string",
		InstrumentToken:   "NSE_EQ|INE528G01035",
		OrderType:         "SL-M",
		TransactionType:   "BUY",
		DisclosedQuantity: 0,
		TriggerPrice:      21.5,
		IsAmo:             false,
	}
	resp, _, err := orderApi.PlaceOrder(context.Background(), body, nil)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling OrderApi.PlaceOrder:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling OrderApi.PlaceOrder:", err)
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

## Place an intraday market order

Intraday BUY at market price.

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

	body := models.PlaceOrderRequest{
		Quantity:          1,
		Product:           "I",
		Validity:          "DAY",
		Price:             0.0,
		Tag:               "string",
		InstrumentToken:   "NSE_EQ|INE528G01035",
		OrderType:         "MARKET",
		TransactionType:   "BUY",
		DisclosedQuantity: 0,
		TriggerPrice:      0.0,
		IsAmo:             false,
	}
	resp, _, err := orderApi.PlaceOrder(context.Background(), body, nil)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling OrderApi.PlaceOrder:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling OrderApi.PlaceOrder:", err)
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

## Place an intraday limit order

Intraday BUY with a limit price of 20.0.

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

	body := models.PlaceOrderRequest{
		Quantity:          1,
		Product:           "I",
		Validity:          "DAY",
		Price:             20.0,
		Tag:               "string",
		InstrumentToken:   "NSE_EQ|INE528G01035",
		OrderType:         "LIMIT",
		TransactionType:   "BUY",
		DisclosedQuantity: 0,
		TriggerPrice:      20.1,
		IsAmo:             false,
	}
	resp, _, err := orderApi.PlaceOrder(context.Background(), body, nil)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling OrderApi.PlaceOrder:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling OrderApi.PlaceOrder:", err)
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

## Place an intraday stop-loss order

Intraday stop-loss limit BUY triggered at 19.5.

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

	body := models.PlaceOrderRequest{
		Quantity:          1,
		Product:           "I",
		Validity:          "DAY",
		Price:             20,
		Tag:               "string",
		InstrumentToken:   "NSE_EQ|INE528G01035",
		OrderType:         "SL",
		TransactionType:   "BUY",
		DisclosedQuantity: 0,
		TriggerPrice:      19.5,
		IsAmo:             false,
	}
	resp, _, err := orderApi.PlaceOrder(context.Background(), body, nil)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling OrderApi.PlaceOrder:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling OrderApi.PlaceOrder:", err)
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

## Place an intraday stop-loss market order

Intraday stop-loss market BUY triggered at 21.5.

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

	body := models.PlaceOrderRequest{
		Quantity:          1,
		Product:           "I",
		Validity:          "DAY",
		Price:             0.0,
		Tag:               "string",
		InstrumentToken:   "NSE_EQ|INE528G01035",
		OrderType:         "SL-M",
		TransactionType:   "BUY",
		DisclosedQuantity: 0,
		TriggerPrice:      21.5,
		IsAmo:             false,
	}
	resp, _, err := orderApi.PlaceOrder(context.Background(), body, nil)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling OrderApi.PlaceOrder:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling OrderApi.PlaceOrder:", err)
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

## Place a delivery market amo (after market order)

Delivery market BUY queued as an after-market order.

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

	body := models.PlaceOrderRequest{
		Quantity:          1,
		Product:           "D",
		Validity:          "DAY",
		Price:             0.0,
		Tag:               "string",
		InstrumentToken:   "NSE_EQ|INE528G01035",
		OrderType:         "MARKET",
		TransactionType:   "BUY",
		DisclosedQuantity: 0,
		TriggerPrice:      0.0,
		IsAmo:             true,
	}
	resp, _, err := orderApi.PlaceOrder(context.Background(), body, nil)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling OrderApi.PlaceOrder:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling OrderApi.PlaceOrder:", err)
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
