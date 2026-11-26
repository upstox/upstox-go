// Command bear_butterfly places a Bear Butterfly (Bearish) options strategy on Nifty 50.
// Buys an ATM put, sells two ATM-1 puts, and buys an ATM-2 put for the next weekly expiry.
package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/antihax/optional"
	"github.com/upstox/upstox-go/upstox_client"
	"github.com/upstox/upstox-go/upstox_client/api"
	"github.com/upstox/upstox-go/upstox_client/models"
)

// underlying is the index whose options are traded; e.g. "Nifty Bank" or "SENSEX".
const underlying = "Nifty 50"

func main() {
	cfg := upstox_client.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox_client.NewAPIClient(cfg)
	instrumentsApi := api.NewInstrumentsApi(client)
	orderApi := api.NewOrderV3Api(client)
	ctx := context.Background()

	// Step 1: Find all three strikes
	upperLong := findOption(ctx, instrumentsApi, "PE", "next_week", 0)
	fmt.Printf("Upper long  (ATM)   - Trading symbol : %s\n", upperLong.TradingSymbol)

	middleShort := findOption(ctx, instrumentsApi, "PE", "next_week", -1)
	fmt.Printf("Middle short (ATM-1) - Trading symbol : %s\n", middleShort.TradingSymbol)

	lowerLong := findOption(ctx, instrumentsApi, "PE", "next_week", -2)
	fmt.Printf("Lower long  (ATM-2) - Trading symbol : %s\n", lowerLong.TradingSymbol)

	// Step 2: Buy the upper put (ATM)
	placeOrder(ctx, orderApi, upperLong, 1, "BUY", "Upper long put order placed successfully.")

	// Step 3: Sell two middle puts (ATM-1)
	placeOrder(ctx, orderApi, middleShort, 2, "SELL", "Middle short put order placed successfully.")

	// Step 4: Buy the lower put (ATM-2)
	placeOrder(ctx, orderApi, lowerLong, 1, "BUY", "Lower long put order placed successfully.")
}

// findOption returns the first underlying option matching the type, expiry and ATM offset.
func findOption(ctx context.Context, instrumentsApi *api.InstrumentsApiService, optionType, expiry string, atmOffset int32) models.InstrumentData {
	resp, _, err := instrumentsApi.SearchInstrument(ctx, underlying, &api.InstrumentsApiSearchInstrumentOpts{
		Exchanges:       optional.NewString("NSE"),
		Segments:        optional.NewString("FO"),
		InstrumentTypes: optional.NewString(optionType),
		Expiry:          optional.NewString(expiry),
		AtmOffset:       optional.NewInt32(atmOffset),
	})
	if err != nil {
		exitOnError(err)
	}
	if len(resp.Data) == 0 {
		fmt.Printf("No %s option found for expiry=%s atm_offset=%d\n", optionType, expiry, atmOffset)
		os.Exit(1)
	}
	return resp.Data[0]
}

// placeOrder places a MARKET delivery order for lots × lot size of the instrument and prints the order ids.
func placeOrder(ctx context.Context, orderApi *api.OrderV3ApiService, leg models.InstrumentData, lots int32, transactionType, label string) {
	resp, _, err := orderApi.PlaceOrderV3(ctx, models.PlaceOrderV3Request{
		InstrumentToken:   leg.InstrumentKey,
		Quantity:          int32(leg.LotSize) * lots,
		TransactionType:   transactionType,
		OrderType:         "MARKET",
		Product:           "D",
		Validity:          "DAY",
		Price:             0,
		DisclosedQuantity: 0,
		TriggerPrice:      0,
		MarketProtection:  -1,
		IsAmo:             false,
	}, nil)
	if err != nil {
		exitOnError(err)
	}
	var orderIDs []string
	if resp.Data != nil {
		orderIDs = resp.Data.OrderIds
	}
	fmt.Printf("%s Order ID: %v\n", label, orderIDs)
}

// exitOnError prints the API error body when available and stops, so no further legs are placed.
func exitOnError(err error) {
	var apiErr upstox_client.GenericSwaggerError
	if errors.As(err, &apiErr) {
		fmt.Printf("API error: %v: %s\n", err, apiErr.Body())
	} else {
		fmt.Printf("API error: %v\n", err)
	}
	os.Exit(1)
}
