// Command long_calendar_call places a Long Calendar with Calls (Bullish) options strategy on Nifty 50.
// Sells a current-week ATM call and buys a next-week ATM call at the same strike.
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

	// Step 1: Find the near-term ATM call (current week — short leg)
	nearCall := findOption(ctx, instrumentsApi, "CE", "current_week", 0)
	fmt.Printf("Near-term call (current week) - Trading symbol : %s\n", nearCall.TradingSymbol)
	fmt.Printf("Near-term call (current week) - Instrument key : %s\n", nearCall.InstrumentKey)

	// Step 2: Find the far-term ATM call (next week — long leg)
	farCall := findOption(ctx, instrumentsApi, "CE", "next_week", 0)
	fmt.Printf("Far-term call  (next week)    - Trading symbol : %s\n", farCall.TradingSymbol)
	fmt.Printf("Far-term call  (next week)    - Instrument key : %s\n", farCall.InstrumentKey)

	// Step 3: Sell the near-term call
	placeOrder(ctx, orderApi, nearCall, 1, "SELL", "Near-term call order placed successfully.")

	// Step 4: Buy the far-term call
	placeOrder(ctx, orderApi, farCall, 1, "BUY", "Far-term call order placed successfully.")
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
