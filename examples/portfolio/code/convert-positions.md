## Convert a position from intraday to delivery

Converts 1 unit of an intraday position to delivery.

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
	portfolioApi := api.NewPortfolioApi(client)

	body := models.ConvertPositionRequest{
		InstrumentToken: "NSE_EQ|INE528G01035",
		Quantity:        1,
		NewProduct:      "D",
		OldProduct:      "I",
		TransactionType: "BUY",
	}
	resp, _, err := portfolioApi.ConvertPositions(context.Background(), body)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling PortfolioApi.ConvertPositions:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling PortfolioApi.ConvertPositions:", err)
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

## Convert a position from delivery to intraday

Converts 1 unit of a delivery position to intraday.

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
	portfolioApi := api.NewPortfolioApi(client)

	body := models.ConvertPositionRequest{
		InstrumentToken: "NSE_EQ|INE528G01035",
		Quantity:        1,
		NewProduct:      "I",
		OldProduct:      "D",
		TransactionType: "BUY",
	}
	resp, _, err := portfolioApi.ConvertPositions(context.Background(), body)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling PortfolioApi.ConvertPositions:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling PortfolioApi.ConvertPositions:", err)
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
