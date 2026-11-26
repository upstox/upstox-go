//go:build ignore

// Command sandbox_sanity is a port of the Python SDK's test/sdk_tests/sandbox_sanity.py against the Upstox sandbox.
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	data_token "github.com/upstox/upstox-go/test/sdk_tests"
	"github.com/upstox/upstox-go/upstox_client"
	"github.com/upstox/upstox-go/upstox_client/api"
	"github.com/upstox/upstox-go/upstox_client/models"
)

const ideaKey = "NSE_EQ|INE669E01016"

var ctx = context.Background()

// runner counts checks and failures.
type runner struct{ checks, failed int }

func main() {
	token := data_token.SandboxToken()
	if token == "" {
		fmt.Println("no sandbox token: set UPSTOX_SANDBOX_ACCESS_TOKEN, or SandboxAccessToken in test/sdk_tests/data_token.go (never commit it)")
		os.Exit(1)
	}
	cfg := upstox_client.NewSandboxConfiguration()
	cfg.AccessToken = token
	cfg.HTTPClient = &http.Client{Timeout: 30 * time.Second}
	c := upstox_client.NewAPIClient(cfg)
	r := &runner{}

	o := api.NewOrderApi(c)
	placed, _, err := o.PlaceOrder(ctx, models.PlaceOrderRequest{
		Quantity: 1, Product: "D", Validity: "DAY", Price: 12.0, Tag: "string",
		InstrumentToken: "NSE_EQ|INE528G01035", OrderType: "LIMIT", TransactionType: "BUY",
	}, nil)
	r.show(placed, err, "exception in place order")

	_, _, err = o.ModifyOrder(ctx, models.ModifyOrderRequest{
		Quantity: 2, Validity: "DAY", Price: 0, OrderId: "250121112512866", OrderType: "MARKET",
	}, nil)
	r.expect("modify order giving wrong response", err, "UDAPI100010")

	multi, _, err := o.PlaceMultiOrder(ctx, []models.MultiOrderRequest{
		{Quantity: 1, Product: "I", Validity: "DAY", Price: 0, Tag: "kg_python_sdk", InstrumentToken: ideaKey,
			OrderType: "MARKET", TransactionType: "BUY", IsAmo: true, CorrelationId: "1"},
		{Quantity: 1, Product: "D", Validity: "DAY", Price: 8.9, Tag: "kg_python_sdk1", InstrumentToken: ideaKey,
			OrderType: "LIMIT", TransactionType: "BUY", IsAmo: true, CorrelationId: "2"},
	}, nil)
	r.show(multi, err, "Exception when calling OrderApi->place_multi_order")

	_, _, err = o.CancelOrder(ctx, "2501211126016789", nil)
	r.expect("cancel order giving wrong response", err, "UDAPI100010")

	v3 := api.NewOrderV3Api(c)
	placedV3, _, err := v3.PlaceOrderV3(ctx, models.PlaceOrderV3Request{
		Quantity: 1, Product: "D", Validity: "DAY", Price: 9.12, Tag: "string", InstrumentToken: ideaKey,
		OrderType: "LIMIT", TransactionType: "BUY", IsAmo: true, Slice: true,
	}, nil)
	r.show(placedV3, err, "Exception when calling OrderApiV3->place_order")

	_, _, err = v3.ModifyOrderV3(ctx, models.ModifyOrderRequest{
		Quantity: 3, Validity: "DAY", Price: 9.12, OrderId: "250120010373947", OrderType: "LIMIT",
	}, nil)
	r.expect("modify order v3 giving wrong response", err, "UDAPI100010")

	_, _, err = v3.CancelOrderV3(ctx, "25012112309073", nil)
	r.expect("cancel order v3 giving wrong response", err, "UDAPI100010")

	fmt.Printf("%d checks, %d failed\n", r.checks, r.failed)
	if r.failed > 0 {
		os.Exit(1)
	}
}

// show prints the response like Python's print(api_response); an error counts as a failure.
func (r *runner) show(resp interface{}, err error, errMsg string) {
	r.checks++
	if err != nil {
		r.failed++
		fmt.Printf("%s: %s\n", errMsg, describe(err))
		return
	}
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

// expect passes on success or when the error body carries one of codes.
func (r *runner) expect(msg string, err error, codes ...string) {
	r.checks++
	if err == nil {
		return
	}
	for _, got := range errorCodes(err) {
		for _, want := range codes {
			if got == want {
				return
			}
		}
	}
	r.failed++
	fmt.Printf("%s: %s\n", msg, describe(err))
}

// errorCodes extracts errors[].errorCode (or error_code) from a GenericSwaggerError body.
func errorCodes(err error) []string {
	var apiErr upstox_client.GenericSwaggerError
	if !errors.As(err, &apiErr) {
		return nil
	}
	var body struct {
		Errors []struct {
			ErrorCode  string `json:"errorCode"`
			ErrorCode2 string `json:"error_code"`
		} `json:"errors"`
	}
	if json.Unmarshal(apiErr.Body(), &body) != nil {
		return nil
	}
	var codes []string
	for _, e := range body.Errors {
		if e.ErrorCode != "" {
			codes = append(codes, e.ErrorCode)
		}
		if e.ErrorCode2 != "" {
			codes = append(codes, e.ErrorCode2)
		}
	}
	return codes
}

// describe returns the error text plus the API error body, if any.
func describe(err error) string {
	var apiErr upstox_client.GenericSwaggerError
	if errors.As(err, &apiErr) && len(apiErr.Body()) > 0 {
		return err.Error() + " " + string(apiErr.Body())
	}
	return err.Error()
}
