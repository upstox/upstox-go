// Command batman places a Batman (Neutral) options strategy on Nifty 50.
// A double-butterfly strategy combining a call butterfly and a put butterfly around ATM.
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

	// Step 1: Fetch all required instruments
	atmCall := findOption(ctx, instrumentsApi, "CE", "next_week", 0)
	fmt.Printf("ATM call      - Trading symbol : %s\n", atmCall.TradingSymbol)

	otm1Call := findOption(ctx, instrumentsApi, "CE", "next_week", 1)
	fmt.Printf("ATM+1 call    - Trading symbol : %s\n", otm1Call.TradingSymbol)

	otm2Call := findOption(ctx, instrumentsApi, "CE", "next_week", 2)
	fmt.Printf("ATM+2 call    - Trading symbol : %s\n", otm2Call.TradingSymbol)

	atmPut := findOption(ctx, instrumentsApi, "PE", "next_week", 0)
	fmt.Printf("ATM put       - Trading symbol : %s\n", atmPut.TradingSymbol)

	otm1Put := findOption(ctx, instrumentsApi, "PE", "next_week", -1)
	fmt.Printf("ATM-1 put     - Trading symbol : %s\n", otm1Put.TradingSymbol)

	otm2Put := findOption(ctx, instrumentsApi, "PE", "next_week", -2)
	fmt.Printf("ATM-2 put     - Trading symbol : %s\n", otm2Put.TradingSymbol)

	// Step 2: Place all legs
	// Leg 1: Buy ATM call (1x)
	placeOrder(ctx, orderApi, atmCall, 1, "BUY", "ATM call order placed successfully.")

	// Leg 2: Sell ATM+1 call (2x)
	placeOrder(ctx, orderApi, otm1Call, 2, "SELL", "ATM+1 call order placed successfully.")

	// Leg 3: Buy ATM+2 call (1x)
	placeOrder(ctx, orderApi, otm2Call, 1, "BUY", "ATM+2 call order placed successfully.")

	// Leg 4: Buy ATM put (1x)
	placeOrder(ctx, orderApi, atmPut, 1, "BUY", "ATM put order placed successfully.")

	// Leg 5: Sell ATM-1 put (2x)
	placeOrder(ctx, orderApi, otm1Put, 2, "SELL", "ATM-1 put order placed successfully.")

	// Leg 6: Buy ATM-2 put (1x)
	placeOrder(ctx, orderApi, otm2Put, 1, "BUY", "ATM-2 put order placed successfully.")
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
