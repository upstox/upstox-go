## Get margin details for equity delivery orders

Required margin for a delivery BUY of 5 units of a equity instrument.

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
	chargeApi := api.NewChargeApi(client)

	body := models.MarginRequest{
		Instruments: []models.Instrument{
			{InstrumentKey: "NSE_EQ|INE528G01035", Quantity: 5, Product: "D", TransactionType: "BUY"},
		},
	}
	resp, _, err := chargeApi.PostMargin(context.Background(), body)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling ChargeApi.PostMargin:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling ChargeApi.PostMargin:", err)
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

## Get margin details for future delivery orders

Required margin for a delivery BUY of 5 units of a futures instrument.

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
	chargeApi := api.NewChargeApi(client)

	body := models.MarginRequest{
		Instruments: []models.Instrument{
			{InstrumentKey: "NSE_FO|35000", Quantity: 5, Product: "D", TransactionType: "BUY"},
		},
	}
	resp, _, err := chargeApi.PostMargin(context.Background(), body)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling ChargeApi.PostMargin:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling ChargeApi.PostMargin:", err)
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

## Get margin details for option delivery orders

Required margin for a delivery BUY of 5 units of a options instrument.

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
	chargeApi := api.NewChargeApi(client)

	body := models.MarginRequest{
		Instruments: []models.Instrument{
			{InstrumentKey: "NSE_FO|54524", Quantity: 5, Product: "D", TransactionType: "BUY"},
		},
	}
	resp, _, err := chargeApi.PostMargin(context.Background(), body)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling ChargeApi.PostMargin:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling ChargeApi.PostMargin:", err)
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

## Get margin details for MCX delivery orders

Required margin for a delivery BUY of 5 units of a MCX instrument.

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
	chargeApi := api.NewChargeApi(client)

	body := models.MarginRequest{
		Instruments: []models.Instrument{
			{InstrumentKey: "MCX_FO|435356", Quantity: 5, Product: "D", TransactionType: "BUY"},
		},
	}
	resp, _, err := chargeApi.PostMargin(context.Background(), body)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling ChargeApi.PostMargin:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling ChargeApi.PostMargin:", err)
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

## Get margin details for currency delivery orders

Required margin for a delivery BUY of 5 units of a currency instrument.

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
	chargeApi := api.NewChargeApi(client)

	body := models.MarginRequest{
		Instruments: []models.Instrument{
			{InstrumentKey: "NCD_FO|15758", Quantity: 5, Product: "D", TransactionType: "BUY"},
		},
	}
	resp, _, err := chargeApi.PostMargin(context.Background(), body)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling ChargeApi.PostMargin:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling ChargeApi.PostMargin:", err)
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
