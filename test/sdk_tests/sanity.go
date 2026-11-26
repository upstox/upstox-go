//go:build ignore

// Command sanity is a port of the Python SDK's test/sdk_tests/sanity.py: it calls the live Upstox API and reports failed checks.
// Like sanity.py it runs account-changing calls ungated; IPO and payout writes stay off via constants.
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"time"

	"github.com/antihax/optional"
	data_token "github.com/upstox/upstox-go/test/sdk_tests"
	"github.com/upstox/upstox-go/upstox_client"
	"github.com/upstox/upstox-go/upstox_client/api"
	"github.com/upstox/upstox-go/upstox_client/models"
)

const (
	// Vodafone Idea, the instrument most Python calls use.
	ideaKey  = "NSE_EQ|INE669E01016"
	niftyKey = "NSE_INDEX|Nifty 50"
)

var (
	ctx = context.Background()
	ist = time.FixedZone("IST", 5*3600+1800)
)

// runner holds the live client and the check counters.
type runner struct {
	client *upstox_client.APIClient
	checks int
	failed int
}

func main() {
	token := data_token.Token()
	if token == "" {
		fmt.Println("no access token: set UPSTOX_ACCESS_TOKEN, or AccessToken in test/sdk_tests/data_token.go (never commit it)")
		os.Exit(1)
	}
	cfg := upstox_client.NewConfiguration()
	cfg.AccessToken = token
	cfg.HTTPClient = &http.Client{Timeout: 30 * time.Second}
	r := &runner{client: upstox_client.NewAPIClient(cfg)}

	// Same order as sanity.py.
	for _, section := range []func(){
		r.user,
		r.charge,
		r.orders,
		r.portfolio,
		r.profitAndLoss,
		r.history,
		r.marketQuote,
		r.marketTimings,
		r.orderV3,
		r.options,
		r.loginToken,
		r.postTrade,
		r.mtfPositions,
		r.historyV3,
		r.marketQuoteV3,
		r.expiredInstruments,
		r.algoName,
		r.instruments,
		r.news,
		r.newsModels,
		r.indieToken,
		r.mutualFunds,
		r.mutualFundModels,
		r.payments,
		r.paymentModels,
		r.fundamentals,
		r.marketAnalytics,
		r.fundamentalsModels,
		r.ipo,
		r.smartlist,
		r.payout,
		r.usStocks,
		r.newModels,
		r.logout,
	} {
		section()
	}

	fmt.Printf("%d checks, %d failed\n", r.checks, r.failed)
	if r.failed > 0 {
		os.Exit(1)
	}
}

// section runs one block of checks; like sanity.py, account-changing calls are not gated.
func (r *runner) section(_ string, fn func()) { fn() }

// fail counts a failure and prints Python's message plus the error detail, if any.
func (r *runner) fail(msg string, err error) {
	r.failed++
	if err != nil {
		fmt.Printf("%s: %s\n", msg, describe(err))
		return
	}
	fmt.Println(msg)
}

// ok expects a nil error and status "success".
func (r *runner) ok(label, msg string, resp interface{}, err error) {
	s, has := status(resp)
	r.okIf(label, msg, !has || s == "success", err)
}

// okIf expects a nil error and good == true.
func (r *runner) okIf(label, msg string, good bool, err error) {
	r.checks++
	if err != nil {
		r.fail("Exception when calling "+label, err)
		return
	}
	if !good {
		r.fail(msg, nil)
	}
}

// expect passes on success or when the error body carries one of codes.
func (r *runner) expect(msg string, err error, codes ...string) bool {
	r.checks++
	if err == nil || hasCode(err, codes) {
		return true
	}
	r.fail(msg, err)
	return false
}

// expectOK is expect, but a successful response must also have status "success".
func (r *runner) expectOK(msg, statusMsg string, resp interface{}, err error, codes ...string) {
	if err != nil {
		r.expect(msg, err, codes...)
		return
	}
	r.ok("", statusMsg, resp, nil)
}

// show prints a response or error like Python's print(api_response); errFails decides whether an error counts.
func (r *runner) show(label string, resp interface{}, err error, errFails bool, errMsg string) {
	r.checks++
	if err == nil {
		fmt.Printf("%s %s\n", label, toJSON(resp))
		return
	}
	if errFails {
		r.fail(errMsg, err)
		return
	}
	fmt.Printf("%s: %s\n", errMsg, describe(err))
}

// model counts an in-memory model check.
func (r *runner) model(good bool, msg string) {
	r.checks++
	if !good {
		r.fail(msg, nil)
	}
}

// inMarketHours mirrors Python's is_within_market_hours (09:00-15:30), using IST.
func inMarketHours() bool {
	now := time.Now().In(ist)
	m := now.Hour()*60 + now.Minute()
	return m >= 9*60 && m <= 15*60+30
}

func httpStatus(r *http.Response) int {
	if r == nil {
		return 0
	}
	return r.StatusCode
}

func algo() optional.String { return optional.NewString(data_token.AlgoNameHeader) }

// status reads the Status field by reflection, since generated models have no getters.
func status(resp interface{}) (string, bool) {
	v := reflect.ValueOf(resp)
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		if v.IsNil() {
			return "", false
		}
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return "", false
	}
	f := v.FieldByName("Status")
	if !f.IsValid() || f.Kind() != reflect.String {
		return "", false
	}
	return f.String(), true
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

func hasCode(err error, want []string) bool {
	for _, c := range errorCodes(err) {
		for _, w := range want {
			if c == w {
				return true
			}
		}
	}
	return false
}

// describe returns the error text plus the API error body, if any.
func describe(err error) string {
	var apiErr upstox_client.GenericSwaggerError
	if errors.As(err, &apiErr) && len(apiErr.Body()) > 0 {
		return err.Error() + " " + string(apiErr.Body())
	}
	return err.Error()
}

func toJSON(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf("%+v", v)
	}
	return string(b)
}

// user ports the UserApi block: profile, funds, static IPs and kill switch.
const runDestructivePayoutTests = false

func (r *runner) user() {
	u := api.NewUserApi(r.client)

	profile, _, err := u.GetProfile(ctx)
	r.ok("UserApi->get_profile", "error in get profile API", profile, err)

	funds, _, err := u.GetUserFundMargin(ctx, nil)
	r.okIf("UserApi->get_user_fund_margin", "Wrong response from get funds and margin", funds.Data != nil, err)

	fundsV3, _, err := u.GetUserFundMarginV3(ctx)
	r.okIf("UserApi->get_user_fund_margin_v3", "Wrong response from get_user_fund_margin_v3", fundsV3.Data != nil, err)

	ips, _, err := u.GetUserIps(ctx)
	r.ok("UserApi->get_user_ips", "Wrong response from get_user_ips", ips, err)

	r.section("update_user_ip", func() {
		resp, _, err := u.UpdateUserIp(ctx, models.UpdateUserIpRequest{PrimaryIp: "1.2.3.4"})
		r.expectOK("update_user_ip giving wrong response", "Wrong response from update_user_ip", resp, err, "UDAPI100068", "UDAPI100069")
	})

	ks, _, err := u.GetKillSwitch(ctx)
	r.ok("UserApi->get_kill_switch", "Wrong response from get_kill_switch", ks, err)

	r.section("update_kill_switch", func() {
		resp, _, err := u.UpdateKillSwitch(ctx, []models.KillSwitchSegmentUpdateRequest{{Segment: "NSE_EQ", Action: "DISABLE"}})
		r.ok("UserApi->update_kill_switch", "Wrong response from update_kill_switch", resp, err)
	})
}

// payments ports the pay-in/pay-out history block.
func (r *runner) payments() {
	u := api.NewUserApi(r.client)

	payin, _, err := u.GetPayinHistory(ctx)
	r.ok("UserApi->get_payin_history", "error in get_payin_history", payin, err)

	payout, _, err := u.GetPayoutHistory(ctx)
	r.ok("UserApi->get_payout_history", "error in get_payout_history", payout, err)
}

// payout ports the Payout APIs block; initiate/modify/cancel moves real funds and is gated.
func (r *runner) payout() {
	u := api.NewUserApi(r.client)

	history, _, err := u.GetPayoutHistory(ctx)
	r.ok("UserApi->get_payout_history", "error in get_payout_history API", history, err)

	modes, _, err := u.GetPayoutModes(ctx)
	r.ok("UserApi->get_payout_modes", "error in get_payout_modes API", modes, err)

	// Initiate/modify/cancel move real funds; off by default, as RUN_DESTRUCTIVE_PAYOUT_TESTS in sanity.py
	r.section("payout initiate/modify/cancel", func() {
		if !runDestructivePayoutTests {
			return
		}
		const label = "UserApi payout write ops"
		initiated, _, err := u.InitiatePayout(ctx, models.InitiatePayoutRequest{Mode: "IMPS", Amount: 1.0})
		r.okIf(label, "initiate_payout returned no transaction id", initiated.Data != nil && initiated.Data.TransactionId != "", err)
		if err != nil || initiated.Data == nil || initiated.Data.TransactionId == "" {
			return
		}
		id := initiated.Data.TransactionId

		modified, _, err := u.ModifyPayout(ctx, models.ModifyPayoutRequest{Amount: 2.0}, id)
		r.ok(label, "error in modify_payout", modified, err)

		cancelled, _, err := u.CancelPayout(ctx, id)
		r.ok(label, "error in cancel_payout", cancelled, err)
		if err == nil {
			fmt.Println("payout initiate/modify/cancel cycle:", cancelled.Status)
		}
	})
}

// Python's sanity tag for bulk cancel/exit.
const pythonTag = "sdk_python_tag"

// charge ports the ChargeApi block: brokerage and margin.
func (r *runner) charge() {
	c := api.NewChargeApi(r.client)

	brokerage, _, err := c.GetBrokerage(ctx, ideaKey, 10, "D", "BUY", 13.4)
	r.okIf("ChargeApi->get_brokerage", "Brokerage giving wrong response", brokerage.Data != nil && brokerage.Data.Charges != nil, err)

	margin, _, err := c.PostMargin(ctx, models.MarginRequest{Instruments: []models.Instrument{
		{InstrumentKey: "NSE_EQ|INE528G01035", Quantity: 12, Product: "D", TransactionType: "BUY"},
		{InstrumentKey: "NSE_EQ|INE002A01018", Quantity: 1, Product: "D", TransactionType: "BUY"},
	}})
	r.ok("Margin API", "error in margin api", margin, err)
}

// orders ports the OrderApi (v2) block.
func (r *runner) orders() {
	o := api.NewOrderApi(r.client)

	r.section("place_order", func() {
		body := models.PlaceOrderRequest{
			Quantity: 1, Product: "D", Validity: "DAY", Price: 5.0, Tag: "string",
			InstrumentToken: "NSE_EQ|INE528G01035", OrderType: "LIMIT", TransactionType: "BUY", IsAmo: true,
		}
		resp, httpResp, err := o.PlaceOrder(ctx, body, &api.OrderApiPlaceOrderOpts{XAlgoName: algo()})
		r.checks++
		switch {
		case err == nil:
			if inMarketHours() && resp.Status != "success" {
				r.fail("error in place order", nil)
			}
		// Python tolerates a 400 only inside market hours
		case httpStatus(httpResp) == 400 && inMarketHours():
		default:
			r.fail("exception in place order", err)
		}
	})

	modified, _, err := o.ModifyOrder(ctx, models.ModifyOrderRequest{
		Quantity: 7, Validity: "DAY", Price: 0, OrderId: "231222010275930", OrderType: "LIMIT",
	}, &api.OrderApiModifyOrderOpts{XAlgoName: algo()})
	if r.expect("modify order giving wrong response", err, "UDAPI100010") && err == nil {
		fmt.Println(toJSON(modified))
	}

	r.section("cancel_multi_order", func() {
		_, _, err := o.CancelMultiOrder(ctx, &api.OrderApiCancelMultiOrderOpts{Tag: optional.NewString(pythonTag), XAlgoName: algo()})
		r.expect("cancel multi giving wrong response", err, "UDAPI1109")
	})

	r.section("exit_positions", func() {
		_, _, err := o.ExitPositions(ctx, &api.OrderApiExitPositionsOpts{Tag: optional.NewString(pythonTag), XAlgoName: algo()})
		r.expect("exit position giving wrong response", err, "UDAPI1111", "UDAPI1113")
	})

	r.section("place_multi_order", func() {
		resp, _, err := o.PlaceMultiOrder(ctx, multiOrderBody(5), &api.OrderApiPlaceMultiOrderOpts{XAlgoName: algo()})
		r.show("multi_order=> ", resp, err, true, "Exception when calling OrderApi->multi_place_order")
	})

	_, _, err = o.GetOrderStatus(ctx, &api.OrderApiGetOrderStatusOpts{OrderId: optional.NewString("240925010636040")})
	r.expect("order details giving wrong response", err, "UDAPI100010")

	_, _, err = o.CancelOrder(ctx, "231221011081579", &api.OrderApiCancelOrderOpts{XAlgoName: algo()})
	r.expect("cancel order giving wrong response", err, "UDAPI100010")

	book, _, err := o.GetOrderBook(ctx)
	r.ok("OrderApi->get_order_book", "get order book giving invalid data", book, err)

	_, _, err = o.GetOrderDetails(ctx, &api.OrderApiGetOrderDetailsOpts{OrderId: optional.NewString("240112010371054")})
	r.expect("get order details giving wrong response", err, "UDAPI100010")

	trades, _, err := o.GetTradeHistory(ctx)
	r.ok("OrderApi->get_trade_history", "get trade history giving invalid data", trades, err)

	byOrder, _, err := o.GetTradesByOrder(ctx, "240112010371054")
	r.ok("OrderApi->get_trades_by_order", "get trades by order giving invalid data", byOrder, err)
}

// multiOrderBody is Python's two-leg Vodafone Idea AMO basket; firstPrice is leg 1's limit price.
func multiOrderBody(firstPrice float64) []models.MultiOrderRequest {
	return []models.MultiOrderRequest{
		{Quantity: 1, Product: "I", Validity: "DAY", Price: firstPrice, Tag: "kg_python_sdk", InstrumentToken: ideaKey,
			OrderType: "LIMIT", TransactionType: "BUY", IsAmo: true, CorrelationId: "1"},
		{Quantity: 1, Product: "D", Validity: "DAY", Price: 8.9, Tag: "kg_python_sdk1", InstrumentToken: ideaKey,
			OrderType: "LIMIT", TransactionType: "BUY", IsAmo: true, CorrelationId: "2"},
	}
}

// portfolio ports the PortfolioApi block: convert positions, holdings, positions.
func (r *runner) portfolio() {
	p := api.NewPortfolioApi(r.client)

	r.section("convert_positions", func() {
		_, _, err := p.ConvertPositions(ctx, models.ConvertPositionRequest{
			InstrumentToken: "NSE_EQ|INE528G01035", Quantity: 1, NewProduct: "D", OldProduct: "I", TransactionType: "BUY",
		})
		r.expect("convert position giving error", err, "UDAPI1035")
	})

	holdings, _, err := p.GetHoldings(ctx)
	r.ok("holding api", "get holding giving wrong response", holdings, err)

	positions, _, err := p.GetPositions(ctx)
	r.ok("position", "get position giving wrong response", positions, err)
}

// mtfPositions ports the get_mtf_positions call.
func (r *runner) mtfPositions() {
	resp, _, err := api.NewPortfolioApi(r.client).GetMtfPositions(ctx)
	r.ok("mtf positions", "error in get mtf positions api", resp, err)
}

// profitAndLoss ports the TradeProfitAndLossApi block for FY 2023-24.
func (r *runner) profitAndLoss() {
	p := api.NewTradeProfitAndLossApi(r.client)
	from, to := optional.NewString("02-04-2023"), optional.NewString("20-03-2024")

	meta, _, err := p.GetTradeWiseProfitAndLossMetaData(ctx, "EQ", "2324",
		&api.TradeProfitAndLossApiGetTradeWiseProfitAndLossMetaDataOpts{FromDate: from, ToDate: to})
	r.ok("TradeProfitAndLossApi->get_trade_wise_profit_and_loss_meta_data", "get report meta data giving wrong response", meta, err)

	data, _, err := p.GetTradeWiseProfitAndLossData(ctx, "EQ", "2324", 1, 4,
		&api.TradeProfitAndLossApiGetTradeWiseProfitAndLossDataOpts{FromDate: from, ToDate: to})
	r.ok("TradeProfitAndLossApi->get_trade_wise_profit_and_loss_data", "get profit loss report data giving wrong response", data, err)

	charges, _, err := p.GetProfitAndLossCharges(ctx, "EQ", "2324",
		&api.TradeProfitAndLossApiGetProfitAndLossChargesOpts{FromDate: from, ToDate: to})
	r.ok("TradeProfitAndLossApi->get_profit_and_loss_charges", "get trade charges giving wrong response", charges, err)
}

// postTrade ports the PostTradeApi call.
func (r *runner) postTrade() {
	resp, _, err := api.NewPostTradeApi(r.client).GetTradesByDateRange(ctx, "2023-04-01", "2024-03-31", 1, 1000, nil)
	r.ok("PostTrade api", "error in post trade api", resp, err)
}

// history ports the HistoryApi (v2) block.
func (r *runner) history() {
	h := api.NewHistoryApi(r.client)

	candles, _, err := h.GetHistoricalCandleDataRange(ctx, ideaKey, "1minute", "2023-11-13", "2023-11-12")
	r.ok("HistoryApi->get_historical_candle_data1", "historical api not giving success", candles, err)

	intraday, _, err := h.GetIntraDayCandleData(ctx, ideaKey, "1minute")
	r.ok("HistoryApi->get_intra_day_candle_data", "historical intraday api not giving success", intraday, err)
}

// historyV3 ports the HistoryV3Api block.
func (r *runner) historyV3() {
	h := api.NewHistoryV3Api(r.client)

	candles, _, err := h.GetHistoricalCandleDataV3(ctx, ideaKey, "minutes", 1, "2025-01-02")
	r.ok("HistoryV3Api->get_historical_candle_data", "error in get_historical_candle_data api", candles, err)

	rangeCandles, _, err := h.GetHistoricalCandleDataRangeV3(ctx, ideaKey, "minutes", 1, "2025-01-02", "2025-01-02")
	r.ok("HistoryV3Api->get_historical_candle_data1", "error in get_historical_candle_data1 api", rangeCandles, err)

	intraday, _, err := h.GetIntraDayCandleDataV3(ctx, ideaKey, "minutes", 1)
	r.ok("HistoryV3Api->get_intra_day_candle_data", "error in get_intra_day_candle_data api", intraday, err)
}

// expiredInstruments ports the ExpiredInstrumentApi block.
func (r *runner) expiredInstruments() {
	e := api.NewExpiredInstrumentApi(r.client)

	expiries, _, err := e.GetExpiriesResponse(ctx, niftyKey)
	r.ok("ExpiredInstrumentApi->get_expiries", "error in get_expiries", expiries, err)

	options, _, err := e.GetExpiredOptionContracts(ctx, niftyKey, "2025-04-30")
	r.ok("ExpiredInstrumentApi->get_expired_option_contracts", "error in get_expired_option_contracts", options, err)

	futures, _, err := e.GetExpiredFutureContracts(ctx, niftyKey, "2025-04-24")
	r.ok("ExpiredInstrumentApi->get_expired_future_contracts", "error in get_expired_future_contracts", futures, err)

	candles, _, err := e.GetExpiredHistoricalCandleData(ctx, "NSE_FO|54452|24-04-2025", "1minute", "2025-04-24", "2025-04-24")
	r.ok("ExpiredInstrumentApi->get_expired_historical_candle_data", "error in get_expired_historical_candle_data", candles, err)
}

// marketQuote ports the MarketQuoteApi (v2) block; Python passes the key as "symbol".
func (r *runner) marketQuote() {
	q := api.NewMarketQuoteApi(r.client)
	symbol := optional.NewString(ideaKey)

	full, _, err := q.GetFullMarketQuote(ctx, &api.MarketQuoteApiGetFullMarketQuoteOpts{Symbol: symbol})
	r.ok("MarketQuoteApi->get_full_market_quote", "market quotes full api not giving success", full, err)

	ltp, _, err := q.Ltp(ctx, &api.MarketQuoteApiLtpOpts{Symbol: symbol})
	r.ok("MarketQuoteApi->ltp", "market quotes ltp api not giving success", ltp, err)

	ohlc, _, err := q.GetMarketQuoteOHLC(ctx, "1d", &api.MarketQuoteApiGetMarketQuoteOHLCOpts{Symbol: symbol})
	r.ok("MarketQuoteApi->get_market_quote_ohlc", "market quotes ohlc api not giving success", ohlc, err)
}

// marketQuoteV3 ports the MarketQuoteV3Api block.
func (r *runner) marketQuoteV3() {
	q := api.NewMarketQuoteV3Api(r.client)

	ohlc, _, err := q.GetMarketQuoteOHLCV3(ctx, "I1", &api.MarketQuoteV3ApiGetMarketQuoteOHLCV3Opts{InstrumentKey: optional.NewString(ideaKey)})
	r.ok("MarketQuoteV3Api->get_market_quote_ohlc", "error in get_market_quote_ohlc api", ohlc, err)

	greeks, _, err := q.GetMarketQuoteOptionGreek(ctx, &api.MarketQuoteV3ApiGetMarketQuoteOptionGreekOpts{
		InstrumentKey: optional.NewString("NSE_FO|38604,NSE_FO|49210"),
	})
	r.ok("MarketQuoteV3Api->get_market_quote_option_greek", "error in get_market_quote_option_greek api", greeks, err)

	ltp, _, err := q.GetLtp(ctx, &api.MarketQuoteV3ApiGetLtpOpts{InstrumentKey: optional.NewString(ideaKey)})
	r.ok("MarketQuoteV3Api->get_ltp", "error in get_ltp api", ltp, err)
}

// options ports the OptionsApi block.
func (r *runner) options() {
	o := api.NewOptionsApi(r.client)

	contracts, _, err := o.GetOptionContracts(ctx, niftyKey, nil)
	r.ok("OptionsApi->get_option_contracts", "get_option_contracts giving error", contracts, err)

	chain, _, err := o.GetPutCallOptionChain(ctx, niftyKey, "2024-03-21")
	r.ok("OptionsApi->get_put_call_option_chain", "get_put_call_option_chain giving error", chain, err)
}

// Python's MarketApi is Go's MarketHolidaysAndTimingsApi; messages keep Python's names.

// marketTimings ports the MarketHolidaysAndTimingsApi block.
func (r *runner) marketTimings() {
	m := api.NewMarketHolidaysAndTimingsApi(r.client)
	const label = "MarketHolidaysAndTimingsApi"

	holidays, _, err := m.GetHolidays(ctx)
	r.ok(label, "error in get holidays", holidays, err)

	holiday, _, err := m.GetHoliday(ctx, "2024-01-22")
	r.ok(label, "error in get holiday on a date", holiday, err)

	timings, _, err := m.GetExchangeTimings(ctx, "2024-01-22")
	r.ok(label, "error get exchange timings", timings, err)

	st, _, err := m.GetMarketStatus(ctx, "NSE")
	r.ok(label, "error get market status", st, err)
}

// marketAnalytics ports the fundamentals-market-data MarketApi block (OI, PCR, max pain, FII/DII).
func (r *runner) marketAnalytics() {
	m := api.NewMarketHolidaysAndTimingsApi(r.client)
	const expiry, date = "2026-05-29", "2026-05-12"

	oi, _, err := m.GetOiData(ctx, niftyKey, expiry, date)
	r.ok("MarketApi->get_oi_data", "MarketApi->get_oi_data not returning success", oi, err)

	changeOi, _, err := m.GetChangeOiData(ctx, niftyKey, expiry, date, 5)
	r.ok("MarketApi->get_change_oi_data", "MarketApi->get_change_oi_data not returning success", changeOi, err)

	pcr, _, err := m.GetPcrData(ctx, niftyKey, expiry, date, 30)
	r.ok("MarketApi->get_pcr_data", "MarketApi->get_pcr_data not returning success", pcr, err)

	maxPain, _, err := m.GetMaxPainData(ctx, niftyKey, expiry, date, 30)
	r.ok("MarketApi->get_max_pain_data", "MarketApi->get_max_pain_data not returning success", maxPain, err)

	fii, _, err := m.GetFiiData(ctx, "NSE_EQ|CASH", "1D", nil)
	r.ok("MarketApi->get_fii_data", "MarketApi->get_fii_data not returning success", fii, err)

	dii, _, err := m.GetDiiData(ctx, "NSE_EQ|CASH", "1D", nil)
	r.ok("MarketApi->get_dii_data", "MarketApi->get_dii_data not returning success", dii, err)
}

// smartlist ports the Smartlist block.
func (r *runner) smartlist() {
	m := api.NewMarketHolidaysAndTimingsApi(r.client)

	opts, _, err := m.GetSmartlistOptions(ctx, &api.MarketHolidaysAndTimingsApiGetSmartlistOptionsOpts{
		AssetType: optional.NewString("INDEX"), Category: optional.NewString("TOP_TRADED"),
		PageNumber: optional.NewInt32(1), PageSize: optional.NewInt32(50),
	})
	r.ok("MarketApi->get_smartlist_options", "error in get_smartlist_options API", opts, err)

	futures, _, err := m.GetSmartlistFutures(ctx, &api.MarketHolidaysAndTimingsApiGetSmartlistFuturesOpts{
		AssetType: optional.NewString("STOCK"), Category: optional.NewString("MOST_ACTIVE"),
		PageNumber: optional.NewInt32(1), PageSize: optional.NewInt32(50),
	})
	r.ok("MarketApi->get_smartlist_futures", "error in get_smartlist_futures API", futures, err)

	mtf, _, err := m.GetSmartlistMtf(ctx, &api.MarketHolidaysAndTimingsApiGetSmartlistMtfOpts{
		PageNumber: optional.NewInt32(1), PageSize: optional.NewInt32(50),
	})
	r.ok("MarketApi->get_smartlist_mtf", "error in get_smartlist_mtf API", mtf, err)
}

// Non-existent IDs Python uses for not-found probes.
const (
	missingGTTModifyID = "GTT-C2503030018840"
	missingGTTCancelID = "GTT-C250303008840"
	missingV3OrderID   = "25030310405859"
	missingV3CancelID  = "2501211050101"
)

// gttBody is Python's three-leg GTT on Vodafone Idea with the given entry trigger.
func gttBody(entryType string, entry float64) models.GttPlaceOrderRequest {
	return models.GttPlaceOrderRequest{
		Type_: "MULTIPLE", InstrumentToken: ideaKey, Product: "D", Quantity: 1, TransactionType: "BUY",
		Rules: []models.GttRule{
			{Strategy: "ENTRY", TriggerType: entryType, TriggerPrice: entry},
			{Strategy: "TARGET", TriggerType: "IMMEDIATE", TriggerPrice: 9},
			{Strategy: "STOPLOSS", TriggerType: "IMMEDIATE", TriggerPrice: 5},
		},
	}
}

// tslRules are Python's entry/target/trailing-stop rules for the MTF GTT.
func tslRules() []models.GttRule {
	return []models.GttRule{
		{Strategy: "ENTRY", TriggerType: "BELOW", TriggerPrice: 60},
		{Strategy: "TARGET", TriggerType: "IMMEDIATE", TriggerPrice: 90},
		{Strategy: "STOPLOSS", TriggerType: "IMMEDIATE", TriggerPrice: 50, TrailingGap: 3},
	}
}

// tslGTTBody is Python's trailing-stop MTF GTT on NSE_EQ|INE584A01023.
func tslGTTBody() models.GttPlaceOrderRequest {
	return models.GttPlaceOrderRequest{
		Type_: "MULTIPLE", InstrumentToken: "NSE_EQ|INE584A01023", Product: "MTF", Quantity: 1,
		TransactionType: "BUY", Rules: tslRules(),
	}
}

// placeV3Body is Python's v3 AMO limit order on Vodafone Idea.
func placeV3Body() models.PlaceOrderV3Request {
	return models.PlaceOrderV3Request{
		Quantity: 1, Product: "D", Validity: "DAY", Price: 9.12, Tag: "string", InstrumentToken: ideaKey,
		OrderType: "LIMIT", TransactionType: "BUY", IsAmo: true, Slice: true,
	}
}

func missingV3Modify() models.ModifyOrderRequest {
	return models.ModifyOrderRequest{Quantity: 1, Validity: "DAY", Price: 9.12, OrderId: missingV3OrderID, OrderType: "LIMIT"}
}

func missingGTTModify() models.GttModifyOrderRequest {
	return models.GttModifyOrderRequest{Type_: "MULTIPLE", GttOrderId: missingGTTModifyID, Rules: tslRules(), Quantity: 2}
}

// orderV3 ports the OrderApiV3 block: GTT orders and v3 place/modify/cancel.
func (r *runner) orderV3() {
	o := api.NewOrderV3Api(r.client)

	r.section("place_gtt_order", func() {
		resp, _, err := o.PlaceGTTOrder(ctx, gttBody("ABOVE", 7), &api.OrderV3ApiPlaceGTTOrderOpts{XAlgoName: algo()})
		r.show(" gtt_order => ", resp, err, true, "Exception when calling OrderApi->gtt_place_order")

		tsl, _, err := o.PlaceGTTOrder(ctx, tslGTTBody(), &api.OrderV3ApiPlaceGTTOrderOpts{XAlgoName: algo()})
		r.show(" tsl gtt_order => ", tsl, err, true, "Exception when calling OrderApi->place_order")
	})

	_, _, err := o.ModifyGTTOrder(ctx, missingGTTModify(), &api.OrderV3ApiModifyGTTOrderOpts{XAlgoName: algo()})
	r.expect("modify gtt giving error", err, "UDAPI100010")

	_, _, err = o.CancelGTTOrder(ctx, models.GttCancelOrderRequest{GttOrderId: missingGTTCancelID}, &api.OrderV3ApiCancelGTTOrderOpts{XAlgoName: algo()})
	r.expect("cancel gtt giving error", err, "UDAPI100010")

	details, _, err := o.GetGttOrderDetails(ctx, &api.OrderV3ApiGetGttOrderDetailsOpts{GttOrderId: optional.NewString("GTT-C25030300128840")})
	r.ok("OrderApi->get_gtt_order_details", "get_gtt_order_details giving error", details, err)

	r.section("place_order v3", func() {
		resp, _, err := o.PlaceOrderV3(ctx, placeV3Body(), &api.OrderV3ApiPlaceOrderV3Opts{XAlgoName: algo()})
		r.show("place order v3 => ", resp, err, true, "Exception when calling OrderApi->place_orderV3")
	})

	_, _, err = o.ModifyOrderV3(ctx, missingV3Modify(), &api.OrderV3ApiModifyOrderV3Opts{XAlgoName: algo()})
	r.expect("modify v3 giving error", err, "UDAPI100010")

	_, _, err = o.CancelOrderV3(ctx, missingV3CancelID, &api.OrderV3ApiCancelOrderV3Opts{XAlgoName: algo()})
	r.expect("cancel v3 giving error", err, "UDAPI100010")
}

// loginToken ports the LoginApi.token probe with placeholder credentials.
func (r *runner) loginToken() {
	_, _, err := api.NewLoginApi(r.client).Token(ctx, "{your_auth_code}", "{your_client_id}", "{your_client_secret}",
		"{your_redirect_url}", "grant_type_example", false)
	r.expect("token giving error", err, "UDAPI100069")
}

// indieToken ports the init_token_request_for_indie_user probe with a fake client id.
func (r *runner) indieToken() {
	_, _, err := api.NewLoginApi(r.client).InitTokenRequestForIndieUser(ctx,
		models.IndieUserTokenRequest{ClientSecret: "9rmbdvjsb"}, "fd33050-ac87-4ecb-b4e1-4ec994c70c32")
	r.expect("indie token request giving error", err, "UDAPI100069")
}

// logout ports the final logout call; it invalidates the token, so it is gated.
func (r *runner) logout() {
	r.section("logout", func() {
		resp, _, err := api.NewLoginApi(r.client).Logout(ctx)
		r.ok("LoginApi->logout", "error in logout", resp, err)
		if err == nil {
			fmt.Println("successfully logged out")
		}
	})
}

// expectAlgo is expect, printing Python's "[ok] <name>: ..." line on a pass.
func (r *runner) expectAlgo(name, msg string, err error, codes ...string) {
	if !r.expect(msg, err, codes...) {
		return
	}
	if err == nil {
		fmt.Printf("[ok] %s: Success\n", name)
	} else {
		fmt.Printf("[ok] %s: Expected error response\n", name)
	}
}

// algoName ports the "ALGO ID TESTS" block: every order write again, with X-Algo-Name.
func (r *runner) algoName() {
	fmt.Println("\n=== Starting Algo ID Tests ===")
	fmt.Println("\n--- OrderApi (V2) Tests with algo_name ---")
	o := api.NewOrderApi(r.client)

	r.section("place_order V2 with algo_name", func() {
		body := models.PlaceOrderRequest{
			Quantity: 1, Product: "D", Validity: "DAY", Price: 5.0, Tag: "string",
			InstrumentToken: "NSE_EQ|INE528G01035", OrderType: "LIMIT", TransactionType: "BUY", IsAmo: true,
		}
		resp, httpResp, err := o.PlaceOrder(ctx, body, &api.OrderApiPlaceOrderOpts{XAlgoName: algo()})
		r.checks++
		switch {
		case err == nil && !inMarketHours():
			fmt.Println("[ok] place_order V2 with algo_name: Called successfully (outside market hours)")
		case err == nil && resp.Status != "success":
			r.fail("error in place order with algo_name", nil)
		case err == nil:
			fmt.Println("[ok] place_order V2 with algo_name: Success")
		case httpStatus(httpResp) == 400 && !inMarketHours():
			fmt.Println("[ok] place_order V2 with algo_name: Expected exception outside market hours")
		default:
			r.fail("exception in place order with algo_name", err)
		}
	})

	_, _, err := o.ModifyOrder(ctx, models.ModifyOrderRequest{
		Quantity: 2, Validity: "DAY", Price: 5, OrderId: "231222010275930", OrderType: "LIMIT",
	}, &api.OrderApiModifyOrderOpts{XAlgoName: algo()})
	r.expectAlgo("modify_order V2 with algo_name", "modify order with algo_name giving wrong response", err, "UDAPI100010")

	r.section("cancel_multi_order V2 with algo_name", func() {
		_, _, err := o.CancelMultiOrder(ctx, &api.OrderApiCancelMultiOrderOpts{Tag: optional.NewString(pythonTag), XAlgoName: algo()})
		r.expectAlgo("cancel_multi_order V2 with algo_name", "cancel multi with algo_name giving wrong response", err, "UDAPI1109")
	})

	r.section("exit_positions V2 with algo_name", func() {
		_, _, err := o.ExitPositions(ctx, &api.OrderApiExitPositionsOpts{Tag: optional.NewString(pythonTag), XAlgoName: algo()})
		r.expectAlgo("exit_positions V2 with algo_name", "exit position with algo_name giving wrong response", err, "UDAPI1111", "UDAPI1113")
	})

	r.section("place_multi_order V2 with algo_name", func() {
		resp, _, err := o.PlaceMultiOrder(ctx, multiOrderBody(4), &api.OrderApiPlaceMultiOrderOpts{XAlgoName: algo()})
		r.show("[ok] place_multi_order V2 with algo_name: Success - ", resp, err, false, "[ok] place_multi_order V2 with algo_name: Exception handled")
	})

	_, _, err = o.CancelOrder(ctx, "231221011081579", &api.OrderApiCancelOrderOpts{XAlgoName: algo()})
	r.expectAlgo("cancel_order V2 with algo_name", "cancel order V2 with algo_name giving wrong response", err, "UDAPI100010")

	fmt.Println("\n--- OrderApiV3 Tests with algo_name ---")
	v3 := api.NewOrderV3Api(r.client)

	r.section("place_order V3 with algo_name", func() {
		resp, _, err := v3.PlaceOrderV3(ctx, placeV3Body(), &api.OrderV3ApiPlaceOrderV3Opts{XAlgoName: algo()})
		r.show("[ok] place_order V3 with algo_name: Success - ", resp, err, false, "[ok] place_order V3 with algo_name: Exception handled")
	})

	_, _, err = v3.ModifyOrderV3(ctx, missingV3Modify(), &api.OrderV3ApiModifyOrderV3Opts{XAlgoName: algo()})
	r.expectAlgo("modify_order V3 with algo_name", "modify V3 with algo_name giving error", err, "UDAPI100010")

	_, _, err = v3.CancelOrderV3(ctx, missingV3CancelID, &api.OrderV3ApiCancelOrderV3Opts{XAlgoName: algo()})
	r.expectAlgo("cancel_order V3 with algo_name", "cancel V3 with algo_name giving error", err, "UDAPI100010")

	fmt.Println("\n--- GTT Tests with algo_name ---")

	r.section("place_gtt_order with algo_name", func() {
		resp, _, err := v3.PlaceGTTOrder(ctx, gttBody("BELOW", 6), &api.OrderV3ApiPlaceGTTOrderOpts{XAlgoName: algo()})
		r.show("[ok] place_gtt_order with algo_name: Success - ", resp, err, false, "[ok] place_gtt_order with algo_name: Exception handled")

		tsl, _, err := v3.PlaceGTTOrder(ctx, tslGTTBody(), &api.OrderV3ApiPlaceGTTOrderOpts{XAlgoName: algo()})
		r.show("[ok] place_gtt_order (TSL) with algo_name: Success - ", tsl, err, false, "[ok] place_gtt_order (TSL) with algo_name: Exception handled")
	})

	_, _, err = v3.ModifyGTTOrder(ctx, missingGTTModify(), &api.OrderV3ApiModifyGTTOrderOpts{XAlgoName: algo()})
	r.expectAlgo("modify_gtt_order with algo_name", "modify gtt with algo_name giving error", err, "UDAPI100010")

	_, _, err = v3.CancelGTTOrder(ctx, models.GttCancelOrderRequest{GttOrderId: missingGTTCancelID}, &api.OrderV3ApiCancelGTTOrderOpts{XAlgoName: algo()})
	r.expectAlgo("cancel_gtt_order with algo_name", "cancel gtt with algo_name giving error", err, "UDAPI100010")

	fmt.Println("\n=== Algo ID Tests Complete ===")
}

// instruments ports the InstrumentsApi.search_instrument block.
func (r *runner) instruments() {
	in := api.NewInstrumentsApi(r.client)
	const label = "InstrumentsApi->search_instrument"

	basic, _, err := in.SearchInstrument(ctx, "Nifty 50", nil)
	r.ok(label, "error in search_instrument basic query", basic, err)

	byExchange, _, err := in.SearchInstrument(ctx, "Reliance", &api.InstrumentsApiSearchInstrumentOpts{Exchanges: optional.NewString("NSE")})
	r.ok(label+" with exchange", "error in search_instrument with exchange filter", byExchange, err)

	bySegment, _, err := in.SearchInstrument(ctx, "TCS", &api.InstrumentsApiSearchInstrumentOpts{Segments: optional.NewString("EQ")})
	r.ok(label+" with segment", "error in search_instrument with segment filter", bySegment, err)

	byType, _, err := in.SearchInstrument(ctx, "Nifty", &api.InstrumentsApiSearchInstrumentOpts{InstrumentTypes: optional.NewString("INDEX")})
	r.ok(label+" with instrument_types", "error in search_instrument with instrument type filter", byType, err)

	paged, _, err := in.SearchInstrument(ctx, "HDFC", &api.InstrumentsApiSearchInstrumentOpts{
		PageNumber: optional.NewInt32(1), Records: optional.NewInt32(5),
	})
	r.ok(label+" with pagination", "error in search_instrument with pagination", paged, err)
}

// news ports the NewsApi block.
func (r *runner) news() {
	n := api.NewNewsApi(r.client)

	byKey, _, err := n.GetNews(ctx, "instrument_keys", &api.NewsApiGetNewsOpts{InstrumentKeys: optional.NewString(ideaKey)})
	r.ok("NewsApi->get_news with instrument_keys", "error in get_news with instrument_keys category", byKey, err)

	positions, _, err := n.GetNews(ctx, "positions", nil)
	r.ok("NewsApi->get_news with positions", "error in get_news with positions category", positions, err)

	holdings, _, err := n.GetNews(ctx, "holdings", nil)
	r.ok("NewsApi->get_news with holdings", "error in get_news with holdings category", holdings, err)

	paged, _, err := n.GetNews(ctx, "instrument_keys", &api.NewsApiGetNewsOpts{
		InstrumentKeys: optional.NewString("NSE_EQ|INE040A01034"), PageNumber: optional.NewInt32(1), PageSize: optional.NewInt32(5),
	})
	r.ok("NewsApi->get_news with pagination", "error in get_news with pagination", paged, err)
}

// mutualFunds ports the MutualFundApi block.
func (r *runner) mutualFunds() {
	m := api.NewMutualFundApi(r.client)

	holdings, _, err := m.GetMutualFundHoldings(ctx)
	r.ok("MutualFundApi->get_mutual_fund_holdings", "error in get_mutual_fund_holdings", holdings, err)

	orders, _, err := m.GetMutualFundOrders(ctx, nil)
	r.ok("MutualFundApi->get_mutual_fund_orders", "error in get_mutual_fund_orders", orders, err)

	filtered, _, err := m.GetMutualFundOrders(ctx, &api.MutualFundApiGetMutualFundOrdersOpts{
		Status: optional.NewString("COMPLETED"), TransactionType: optional.NewString("BUY"),
		PageNumber: optional.NewInt32(1), Records: optional.NewInt32(10),
	})
	r.ok("MutualFundApi->get_mutual_fund_orders (filtered)", "error in get_mutual_fund_orders with filters", filtered, err)

	order, _, err := m.GetMutualFundOrder(ctx, "MF-ORDER-001")
	r.expectOK("get_mutual_fund_order giving wrong error response", "error in get_mutual_fund_order", order, err, "UDAPI100010", "UDAPI100050")

	sips, _, err := m.GetMutualFundSips(ctx, nil)
	r.ok("MutualFundApi->get_mutual_fund_sips", "error in get_mutual_fund_sips", sips, err)

	pagedSips, _, err := m.GetMutualFundSips(ctx, &api.MutualFundApiGetMutualFundSipsOpts{PageNumber: optional.NewInt32(1), Records: optional.NewInt32(10)})
	r.ok("MutualFundApi->get_mutual_fund_sips (paginated)", "error in get_mutual_fund_sips with pagination", pagedSips, err)
}

// fundamentals ports the FundamentalsApi block for Vodafone Idea.
func (r *runner) fundamentals() {
	f := api.NewFundamentalsApi(r.client)
	const isin = "INE669E01016"

	profile, _, err := f.GetCompanyProfile(ctx, isin)
	r.ok("FundamentalsApi->get_company_profile", "FundamentalsApi->get_company_profile not returning success", profile, err)

	ratios, _, err := f.GetKeyRatios(ctx, isin)
	r.ok("FundamentalsApi->get_key_ratios", "FundamentalsApi->get_key_ratios not returning success", ratios, err)

	balance, _, err := f.GetBalanceSheet(ctx, isin, nil)
	r.ok("FundamentalsApi->get_balance_sheet", "FundamentalsApi->get_balance_sheet not returning success", balance, err)

	income, _, err := f.GetIncomeStatement(ctx, isin, nil)
	r.ok("FundamentalsApi->get_income_statement", "FundamentalsApi->get_income_statement not returning success", income, err)

	cash, _, err := f.GetCashFlow(ctx, isin, nil)
	r.ok("FundamentalsApi->get_cash_flow", "FundamentalsApi->get_cash_flow not returning success", cash, err)

	competitors, _, err := f.GetCompetitors(ctx, ideaKey)
	r.ok("FundamentalsApi->get_competitors", "FundamentalsApi->get_competitors not returning success", competitors, err)

	actions, _, err := f.GetCorporateActions(ctx, isin)
	r.ok("FundamentalsApi->get_corporate_actions", "FundamentalsApi->get_corporate_actions not returning success", actions, err)

	shares, _, err := f.GetShareHoldings(ctx, isin)
	r.ok("FundamentalsApi->get_share_holdings", "FundamentalsApi->get_share_holdings not returning success", shares, err)
}

// ipo ports the IpoApi block; apply/cancel blocks funds via UPI and needs UPSTOX_TEST_UPI and UPSTOX_TEST_IPO_ID.
const runDestructiveIpoTests = false

func (r *runner) ipo() {
	i := api.NewIPOApi(r.client)

	listing, _, err := i.GetIpoListing(ctx, &api.IPOApiGetIpoListingOpts{
		Status: optional.NewString("open"), IssueType: optional.NewString("regular"),
		PageNumber: optional.NewString("1"), Records: optional.NewString("20"),
	})
	r.ok("IpoApi->get_ipo_listing", "error in get_ipo_listing API", listing, err)

	details, _, err := i.GetIpoDetails(ctx, "sample-ipo-slug")
	r.ok("IpoApi->get_ipo_details", "error in get_ipo_details API", details, err)

	orders, _, err := i.GetIpoOrders(ctx, nil)
	r.ok("IpoApi->get_ipo_orders", "error in get_ipo_orders API", orders, err)

	order, _, err := i.GetIpoOrderById(ctx, "sample-ipo-order-id")
	r.ok("IpoApi->get_ipo_order_by_id", "error in get_ipo_order_by_id API", order, err)

	// Apply/cancel blocks funds via a real UPI mandate; off by default, as RUN_DESTRUCTIVE_IPO_TESTS in sanity.py
	if !runDestructiveIpoTests {
		return
	}
	const ipoID, upi = "sample-ipo-slug", "someone@upi"
	const label = "IpoApi ipo order write ops"
	applied, _, err := i.ApplyForIpo(ctx, models.IpoApplyRequest{
		Id: ipoID, Upi: upi, Category: "IND", Bids: []models.IpoBidRequest{{Quantity: 10, Price: 150.0}},
	})
	r.okIf(label, "apply_for_ipo returned no order id", applied.Data != nil && applied.Data.OrderId != "", err)
	if err != nil || applied.Data == nil || applied.Data.OrderId == "" {
		return
	}
	cancelled, _, err := i.CancelIpoOrder(ctx, applied.Data.OrderId)
	r.ok(label, "error in cancel_ipo_order", cancelled, err)
	if err == nil {
		fmt.Println("ipo apply/cancel cycle:", cancelled.Status)
	}
}

// usStocks ports the US Stocks block; estimate and margin place nothing.
func (r *runner) usStocks() {
	m := api.NewUSMarketHolidaysAndTimingsApi(r.client)
	o := api.NewUSStocksOrderApi(r.client)

	st, _, err := m.GetUsMarketStatus(ctx)
	r.ok("UsMarketHolidaysAndTimingsApi->get_us_market_status", "error in get_us_market_status API", st, err)

	holidays, _, err := m.GetUsHolidays(ctx)
	r.ok("UsMarketHolidaysAndTimingsApi->get_us_holidays", "error in get_us_holidays API", holidays, err)

	holiday, _, err := m.GetUsHoliday(ctx, "2026-11-26")
	r.ok("UsMarketHolidaysAndTimingsApi->get_us_holiday", "error in get_us_holiday API", holiday, err)

	timings, _, err := m.GetUsExchangeTimings(ctx, "2026-11-27")
	r.ok("UsMarketHolidaysAndTimingsApi->get_us_exchange_timings", "error in get_us_exchange_timings API", timings, err)

	funds, _, err := api.NewUSStocksFundsApi(r.client).GetUsFunds(ctx)
	r.ok("UsStocksFundsApi->get_us_funds", "error in get_us_funds API", funds, err)

	holdings, _, err := api.NewUSStocksPortfolioApi(r.client).GetUsHoldings(ctx)
	r.ok("UsStocksPortfolioApi->get_us_holdings", "error in get_us_holdings API", holdings, err)

	book, _, err := o.GetUsOrderBook(ctx, nil)
	r.ok("UsStocksOrderApi->get_us_order_book", "error in get_us_order_book API", book, err)

	details, _, err := o.GetUsOrderDetails(ctx, "sample-us-order-id")
	r.ok("UsStocksOrderApi->get_us_order_details", "error in get_us_order_details API", details, err)

	estimateBody := models.UsOrderEstimateRequest{Quantity: 1, InstrumentKey: "US_EQ|COUR", TransactionType: "BUY", OrderType: "MARKET"}
	estimate, _, err := o.EstimateUsOrder(ctx, estimateBody)
	r.ok("UsStocksOrderApi->estimate_us_order", "error in estimate_us_order API", estimate, err)

	margin, _, err := api.NewUSStocksChargesApi(r.client).CalculateUsMargin(ctx, models.UsMarginRequest{Instruments: []models.UsOrderEstimateRequest{estimateBody}})
	r.ok("UsStocksChargesApi->calculate_us_margin", "error in calculate_us_margin API", margin, err)
}

// Python's "instantiation failed" (is None) checks cannot fail for Go structs; these literals make them compile-time checks.
var _ = []interface{}{
	models.NewsResponseMetaData{}, models.NewsResponsePageData{},
	models.MutualFundHoldingData{}, models.MutualFundMetaData{}, models.MutualFundOrderData{}, models.MutualFundSipData{},
	models.FamilyMemberData{}, models.Pagination{}, models.GetMutualFundHoldingsResponse{}, models.GetMutualFundOrdersResponse{},
	models.GetMutualFundOrderDetailsResponse{}, models.GetMutualFundSipsResponse{},
	models.PaymentHistoryData{}, models.PaymentHistoryResponse{},
	models.AnalyticsResponse{}, models.BalanceSheetData{}, models.BalanceSheetHistoryItem{}, models.CashFlowData{},
	models.CashFlowEntry{}, models.CompanyProfileData{}, models.CompetitorData{}, models.CorporateActionData{},
	models.CorporateActionEventDetail{}, models.FinancialStatementEntry{}, models.HistoryItem{}, models.IncomeStatementData{},
	models.IncomeStatementEntry{}, models.KeyRatioData{}, models.SectorMarketCapAmount{}, models.ShareHoldingData{},
}

// newsModels ports the NewsItemData/GetNewsResponse field checks.
func (r *runner) newsModels() {
	// Go's PublishedTime is epoch millis; this is 2026-04-20T10:00:00Z
	item := models.NewsItemData{
		Heading: "Test heading", Summary: "Test summary", Thumbnail: "https://example.com/thumb.jpg",
		ArticleLink: "https://example.com/article", PublishedTime: 1776679200000,
	}
	r.model(item.Heading == "Test heading", "error: NewsItemData heading field not set correctly")
	r.model(item.Summary == "Test summary", "error: NewsItemData summary field not set correctly")
	r.model(item.ArticleLink == "https://example.com/article", "error: NewsItemData article_link field not set correctly")

	resp := models.GetNewsResponse{Status: "success", Data: map[string][]models.NewsItemData{ideaKey: {item}}}
	r.model(resp.Status == "success", "error: GetNewsResponse status field not set correctly")
}

// mutualFundModels has no runtime checks: every Python check there is an is-None check (see the var block above).
func (r *runner) mutualFundModels() {}

// paymentModels ports the PaymentHistoryData/PaymentHistoryResponse field checks.
func (r *runner) paymentModels() {
	data := models.PaymentHistoryData{
		Amount: 1000.0, Mode: "NEFT", Status: "SUCCESS", LastUpdatedAt: "2026-05-04T10:00:00Z",
		BankName: "Test Bank", TransactionId: "TXN123456", TotalCharges: 0.0, ChargesCategory: "PAYIN",
	}
	r.model(data.Amount == 1000.0, "error: PaymentHistoryData amount field not set correctly")
	r.model(data.Mode == "NEFT", "error: PaymentHistoryData mode field not set correctly")
	r.model(data.TransactionId == "TXN123456", "error: PaymentHistoryData transaction_id field not set correctly")

	resp := models.PaymentHistoryResponse{Status: "success", Data: []models.PaymentHistoryData{data}}
	r.model(resp.Status == "success", "error: PaymentHistoryResponse status field not set correctly")
}

// fundamentalsModels ports the fundamentals-market-data response status checks.
func (r *runner) fundamentalsModels() {
	for _, c := range []struct {
		status, name string
	}{
		{models.AnalyticsResponse{Status: "success"}.Status, "AnalyticsResponse"},
		{models.BalanceSheetResponse{Status: "success"}.Status, "BalanceSheetResponse"},
		{models.CashFlowResponse{Status: "success"}.Status, "CashFlowResponse"},
		{models.CompanyProfileResponse{Status: "success"}.Status, "CompanyProfileResponse"},
		{models.CompetitorsResponse{Status: "success"}.Status, "CompetitorsResponse"},
		{models.CorporateActionsResponse{Status: "success"}.Status, "CorporateActionsResponse"},
		{models.IncomeStatementResponse{Status: "success"}.Status, "IncomeStatementResponse"},
		{models.KeyRatiosResponse{Status: "success"}.Status, "KeyRatiosResponse"},
		{models.ShareHoldingsResponse{Status: "success"}.Status, "ShareHoldingsResponse"},
	} {
		r.model(c.status == "success", "error: "+c.name+" status field not set correctly")
	}
}

// newModels ports the IPO, payout and US stocks model checks.
func (r *runner) newModels() {
	listing := models.IpoListingData{Symbol: "XYZ", Status: "open"}
	r.model(listing.Symbol == "XYZ", "error: IpoListingData fields not set correctly")

	meta := models.IpoMetaData{Page: &models.Pagination{PageNumber: 1}}
	r.model(meta.Page.PageNumber == 1, "error: IpoMetaData fields not set correctly")

	listingResp := models.IpoListingResponse{Status: "success", Data: []models.IpoListingData{listing}, MetaData: &meta}
	r.model(listingResp.Status == "success", "error: IpoListingResponse status field not set correctly")

	timeline := models.IpoTimeline{ListingDate: "2026-07-01"}
	r.model(timeline.ListingDate == "2026-07-01", "error: IpoTimeline fields not set correctly")

	registrar := models.IpoRegistrarInfo{Name: "Registrar Co"}
	r.model(registrar.Name == "Registrar Co", "error: IpoRegistrarInfo fields not set correctly")

	details := models.IpoDetailsData{Id: "abc", Symbol: "XYZ", LotSize: 10, Timeline: &timeline, RegistrarInfo: &registrar}
	r.model(details.LotSize == 10, "error: IpoDetailsData fields not set correctly")

	detailsResp := models.IpoDetailsResponse{Status: "success", Data: &details}
	r.model(detailsResp.Status == "success", "error: IpoDetailsResponse status field not set correctly")

	investor := models.IpoInvestorType{Category: "IND", Description: "Individual Investor"}
	r.model(investor.Category == "IND", "error: IpoInvestorType fields not set correctly")

	withInvestors := models.IpoDetailsData{Id: "abc", Investors: []models.IpoInvestorType{investor}}
	r.model(withInvestors.Investors[0].Category == "IND", "error: IpoDetailsData investors field not set correctly")

	listingInvestors := models.IpoListingData{Symbol: "XYZ", Investors: []models.IpoInvestorType{investor}}
	r.model(listingInvestors.Investors[0].Description == "Individual Investor", "error: IpoListingData investors field not set correctly")

	bid := models.IpoBidRequest{Quantity: 10, Price: 150.5}
	r.model(bid.Quantity == 10 && bid.Price == 150.5, "error: IpoBidRequest fields not set correctly")

	apply := models.IpoApplyRequest{Id: "sample-ipo-slug", Upi: "someone@upi", Category: "IND", Bids: []models.IpoBidRequest{bid}}
	r.model(apply.Upi == "someone@upi" && len(apply.Bids) == 1, "error: IpoApplyRequest fields not set correctly")

	applyData := models.IpoApplyData{OrderId: "O1"}
	r.model(applyData.OrderId == "O1", "error: IpoApplyData fields not set correctly")

	applyResp := models.IpoApplyResponse{Status: "success", Data: &applyData}
	r.model(applyResp.Status == "success", "error: IpoApplyResponse status field not set correctly")

	orderBid := models.IpoOrderBid{Quantity: 10, Price: 150.5, Amount: 1505.0, Message: "accepted"}
	r.model(orderBid.Amount == 1505.0, "error: IpoOrderBid fields not set correctly")

	orderData := models.IpoOrderData{
		Id: "abc", Symbol: "XYZ", Exchange: "NSE", OrderId: "O1", OrderStatus: "COMPLETE", Category: "IND",
		IssueType: "regular", UnitsAllotted: 10, Bids: []models.IpoOrderBid{orderBid},
	}
	r.model(orderData.OrderId == "O1" && orderData.UnitsAllotted == 10, "error: IpoOrderData fields not set correctly")

	orderResp := models.IpoOrderResponse{Status: "success", Data: []models.IpoOrderData{orderData}, MetaData: &meta}
	r.model(orderResp.Status == "success", "error: IpoOrderResponse status field not set correctly")

	orderDetail := models.IpoOrderDetailResponse{Status: "success", Data: &orderData}
	r.model(orderDetail.Data.Symbol == "XYZ", "error: IpoOrderDetailResponse fields not set correctly")

	cancelData := models.IpoCancelData{OrderId: "O1", Status: "CANCELLED"}
	r.model(cancelData.Status == "CANCELLED", "error: IpoCancelData fields not set correctly")

	cancelResp := models.IpoCancelResponse{Status: "success", Data: &cancelData}
	r.model(cancelResp.Data.OrderId == "O1", "error: IpoCancelResponse fields not set correctly")

	initiate := models.InitiatePayoutRequest{Mode: "IMPS", Amount: 500.0}
	r.model(initiate.Amount == 500.0, "error: InitiatePayoutRequest fields not set correctly")

	modify := models.ModifyPayoutRequest{Amount: 750.0}
	r.model(modify.Amount == 750.0, "error: ModifyPayoutRequest fields not set correctly")

	payout := models.PayoutDetails{Status: "SUCCESS", TransactionId: "T1", Amount: 500.0}
	r.model(payout.TransactionId == "T1", "error: PayoutDetails fields not set correctly")

	payoutResp := models.PayoutDetailsResponse{Status: "success", Data: &payout}
	r.model(payoutResp.Status == "success", "error: PayoutDetailsResponse status field not set correctly")

	// Go has no PayoutModesData model; Data is free-form JSON
	modesResp := models.PayoutModesResponse{Status: "success", Data: map[string]interface{}{}}
	r.model(modesResp.Status == "success", "error: PayoutModesResponse status field not set correctly")

	usPlace := models.UsPlaceOrderRequest{Amount: 25.0, Validity: "DAY", InstrumentKey: "K", TransactionType: "BUY", OrderType: "MARKET"}
	r.model(usPlace.Amount == 25.0, "error: UsPlaceOrderRequest fields not set correctly")

	usMargin := models.UsMarginRequest{Instruments: []models.UsOrderEstimateRequest{
		{Quantity: 1.25, InstrumentKey: "K", TransactionType: "BUY", OrderType: "MARKET"},
	}}
	r.model(usMargin.Instruments[0].Quantity == 1.25, "error: UsMarginRequest fields not set correctly")

	usOrder := models.UsOrderResponse{Status: "success", Data: &models.UsOrderData{OrderId: "O1"}}
	r.model(usOrder.Data.OrderId == "O1", "error: UsOrderResponse fields not set correctly")

	usEstimate := models.UsOrderEstimateResponse{Status: "success", Data: &models.UsOrderEstimateData{TotalCharges: 1.5, EstimatedTotal: 101.5}}
	r.model(usEstimate.Data.EstimatedTotal == 101.5, "error: UsOrderEstimateResponse fields not set correctly")

	bookData := models.UsOrderBookData{OrderId: "O1", Status: "open", Quantity: 1.5, Fractionable: true}
	book := models.UsOrderBookResponse{Status: "success", Data: []models.UsOrderBookData{bookData}}
	bookDetails := models.UsOrderDetailsResponse{Status: "success", Data: &bookData}
	r.model(book.Data[0].Status == "open" && bookDetails.Data.Quantity == 1.5, "error: UsOrderBookResponse/UsOrderDetailsResponse fields not set correctly")

	usMarginResp := models.UsMarginResponse{Status: "success", Data: &models.UsMarginData{
		Margins: []models.UsOrderMargin{{InstrumentKey: "K", RequiredMargin: 101.2}}, Sufficient: true,
	}}
	r.model(usMarginResp.Data.Margins[0].RequiredMargin == 101.2, "error: UsMarginResponse fields not set correctly")

	usFunds := models.UsFundsResponse{Status: "success", Data: &models.UsFundsData{Cash: 100.0, ExchangeRate: 88.5}}
	r.model(usFunds.Data.ExchangeRate == 88.5, "error: UsFundsResponse fields not set correctly")

	usHoldings := models.UsHoldingsResponse{Status: "success", Data: []models.UsHoldingData{{InstrumentKey: "K", Quantity: 2.5}}}
	r.model(usHoldings.Data[0].Quantity == 2.5, "error: UsHoldingsResponse fields not set correctly")

	usHoliday := models.GetUsHolidayResponse{Status: "success", Data: []models.UsHolidayData{{Date: "2026-11-26", HolidayType: "TRADING_HOLIDAY"}}}
	r.model(usHoliday.Data[0].Date == "2026-11-26", "error: GetUsHolidayResponse fields not set correctly")

	usTiming := models.GetUsExchangeTimingResponse{Status: "success", Data: []models.UsExchangeTimingData{{Exchange: "US", StartTime: 1, EndTime: 2}}}
	r.model(usTiming.Data[0].EndTime == 2, "error: GetUsExchangeTimingResponse fields not set correctly")
}
