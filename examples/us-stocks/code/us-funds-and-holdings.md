## Get US funds

USD balances available to trade and withdraw.

```go
package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/upstox/upstox-go/upstox_client"
	"github.com/upstox/upstox-go/upstox_client/api"
)

func main() {
	cfg := upstox_client.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox_client.NewAPIClient(cfg)
	usStocksFundsApi := api.NewUSStocksFundsApi(client)

	resp, _, err := usStocksFundsApi.GetUsFunds(context.Background())
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling USStocksFundsApi.GetUsFunds:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling USStocksFundsApi.GetUsFunds:", err)
		}
		return
	}
	if resp.Data == nil {
		fmt.Println("no data returned")
		return
	}
	fmt.Println("available to trade (USD):", resp.Data.AvailableToTrade)
	fmt.Println("exchange rate:", resp.Data.ExchangeRate)
}
```

## Get US holdings

Lists your US stock holdings with P&L.

```go
package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/upstox/upstox-go/upstox_client"
	"github.com/upstox/upstox-go/upstox_client/api"
)

func main() {
	cfg := upstox_client.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox_client.NewAPIClient(cfg)
	usStocksPortfolioApi := api.NewUSStocksPortfolioApi(client)

	resp, _, err := usStocksPortfolioApi.GetUsHoldings(context.Background())
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling USStocksPortfolioApi.GetUsHoldings:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling USStocksPortfolioApi.GetUsHoldings:", err)
		}
		return
	}
	for _, holding := range resp.Data {
		fmt.Println(holding.TradingSymbol, holding.Quantity, holding.Pnl)
	}
}
```
