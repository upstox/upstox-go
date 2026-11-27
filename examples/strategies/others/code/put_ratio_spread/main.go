// Command put_ratio_spread places a Put Ratio Spread options strategy on Nifty 50.
// Buys one ATM put and sells two OTM puts (ATM-1) for the next weekly expiry.
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

	// Step 1: Find the ATM put (long leg)
	longPut := findOption(ctx, instrumentsApi, "PE", "next_week", 0)
	fmt.Printf("Long put  (ATM)   - Trading symbol : %s\n", longPut.TradingSymbol)

	// Step 2: Find the OTM put (short leg — 2x quantity)
	shortPut := findOption(ctx, instrumentsApi, "PE", "next_week", -1)
	fmt.Printf("Short put (ATM-1) - Trading symbol : %s\n", shortPut.TradingSymbol)

	// Step 3: Buy 1x ATM put
	placeOrder(ctx, orderApi, longPut, 1, "BUY", "Long put order placed successfully.")

	// Step 4: Sell 2x OTM put
	placeOrder(ctx, orderApi, shortPut, 2, "SELL", "Short put order placed successfully.")
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
	resp, _, err := orderApi.PlaceOrder(ctx, models.PlaceOrderV3Request{
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
