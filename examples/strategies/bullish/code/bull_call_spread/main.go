// Command bull_call_spread places a Bull Call Spread (Bullish) options strategy on Nifty 50.
// Buys an ATM Nifty 50 call option and sells an OTM call option (ATM+1 strike) for the next weekly expiry.
package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/antihax/optional"
	"github.com/upstox/upstox-go/upstox"
	"github.com/upstox/upstox-go/upstox/api"
	"github.com/upstox/upstox-go/upstox/models"
)

// underlying is the index whose options are traded; e.g. "Nifty Bank" or "SENSEX".
const underlying = "Nifty 50"

func main() {
	cfg := upstox.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox.NewAPIClient(cfg)
	instrumentsApi := api.NewInstrumentsApi(client)
	orderApi := api.NewOrderV3Api(client)
	ctx := context.Background()

	// Step 1: Find the ATM call option (long leg)
	longCall := findOption(ctx, instrumentsApi, "CE", "next_week", 0)
	fmt.Printf("Long call  - Trading symbol : %s\n", longCall.TradingSymbol)
	fmt.Printf("Long call  - Instrument key : %s\n", longCall.InstrumentKey)

	// Step 2: Find the OTM call option one strike above ATM (short leg)
	shortCall := findOption(ctx, instrumentsApi, "CE", "next_week", 1)
	fmt.Printf("Short call - Trading symbol : %s\n", shortCall.TradingSymbol)
	fmt.Printf("Short call - Instrument key : %s\n", shortCall.InstrumentKey)

	// Step 3: Buy the ATM call
	placeOrder(ctx, orderApi, longCall, 1, "BUY", "Long call order placed successfully.")

	// Step 4: Sell the OTM call
	placeOrder(ctx, orderApi, shortCall, 1, "SELL", "Short call order placed successfully.")
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
	var apiErr upstox.GenericSwaggerError
	if errors.As(err, &apiErr) {
		fmt.Printf("API error: %v: %s\n", err, apiErr.Body())
	} else {
		fmt.Printf("API error: %v\n", err)
	}
	os.Exit(1)
}
