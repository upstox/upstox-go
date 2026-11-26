## Get full market quote

Full quote (OHLC, depth, OI, circuit limits) for one instrument.

```go
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/antihax/optional"
	"github.com/upstox/upstox-go/upstox"
	"github.com/upstox/upstox-go/upstox/api"
)

func main() {
	cfg := upstox.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox.NewAPIClient(cfg)
	marketQuoteV3Api := api.NewMarketQuoteV3Api(client)

	// For a single instrument
	resp, _, err := marketQuoteV3Api.GetFullMarketQuoteV3(context.Background(), &api.MarketQuoteV3ApiGetFullMarketQuoteV3Opts{InstrumentKey: optional.NewString("NSE_EQ|INE848E01016")})
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling MarketQuoteV3Api.GetFullMarketQuoteV3:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling MarketQuoteV3Api.GetFullMarketQuoteV3:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```

## Get full market quote for multiple instrument keys

Pass a comma-separated list of instrument keys.

```go
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/antihax/optional"
	"github.com/upstox/upstox-go/upstox"
	"github.com/upstox/upstox-go/upstox/api"
)

func main() {
	cfg := upstox.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox.NewAPIClient(cfg)
	marketQuoteV3Api := api.NewMarketQuoteV3Api(client)

	// For multiple instruments (up to 500 in one call)
	resp, _, err := marketQuoteV3Api.GetFullMarketQuoteV3(context.Background(), &api.MarketQuoteV3ApiGetFullMarketQuoteV3Opts{InstrumentKey: optional.NewString("NSE_EQ|INE848E01016,NSE_EQ|INE669E01016")})
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling MarketQuoteV3Api.GetFullMarketQuoteV3:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling MarketQuoteV3Api.GetFullMarketQuoteV3:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```

## Read fields from the full market quote response

Walks the quote map and prints the main fields, OHLC and market depth.

```go
package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/antihax/optional"
	"github.com/upstox/upstox-go/upstox"
	"github.com/upstox/upstox-go/upstox/api"
)

func main() {
	cfg := upstox.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox.NewAPIClient(cfg)
	marketQuoteV3Api := api.NewMarketQuoteV3Api(client)

	resp, _, err := marketQuoteV3Api.GetFullMarketQuoteV3(context.Background(), &api.MarketQuoteV3ApiGetFullMarketQuoteV3Opts{InstrumentKey: optional.NewString("NSE_EQ|INE848E01016")})
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling MarketQuoteV3Api.GetFullMarketQuoteV3:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling MarketQuoteV3Api.GetFullMarketQuoteV3:", err)
		}
		return
	}
	// Data is keyed by "<exchange>:<trading_symbol>"
	for key, quote := range resp.Data {
		fmt.Println(key)
		fmt.Println("  symbol                :", quote.Symbol)
		fmt.Println("  instrument_token      :", quote.InstrumentToken)
		fmt.Println("  last_price            :", quote.LastPrice)
		fmt.Println("  volume                :", quote.Volume)
		fmt.Println("  average_price         :", quote.AveragePrice)
		fmt.Println("  net_change            :", quote.NetChange)
		fmt.Println("  prev_close_price      :", quote.PrevClosePrice)
		fmt.Println("  lower_circuit_limit   :", quote.LowerCircuitLimit)
		fmt.Println("  upper_circuit_limit   :", quote.UpperCircuitLimit)
		fmt.Println("  year_high / year_low  :", quote.YearHigh, "/", quote.YearLow)
		fmt.Println("  oi / previous_oi      :", quote.Oi, "/", quote.PreviousOi)
		fmt.Println("  cas_eligible          :", quote.CasEligible)
		// Nested OHLC snapshot
		if quote.Ohlc != nil {
			fmt.Println("  ohlc                  :", quote.Ohlc.Open, quote.Ohlc.High, quote.Ohlc.Low, quote.Ohlc.Close)
		}
		// Nested market depth (top 5 bids and asks)
		if quote.Depth != nil {
			for _, level := range quote.Depth.Buy {
				fmt.Println("  bid:", level.Price, level.Quantity, level.Orders)
			}
			for _, level := range quote.Depth.Sell {
				fmt.Println("  ask:", level.Price, level.Quantity, level.Orders)
			}
		}
	}
}
```
