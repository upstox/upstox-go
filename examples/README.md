# Upstox Developer API – Example Code

This folder contains **ready-to-use Go samples** for the [Upstox API](https://upstox.com/developer/api-documentation/open-api). Each example shows how to call the API using the official Upstox Go SDK (`github.com/upstox/upstox-go`).

## Why use these samples?

- **Quick start** — Copy-paste examples for common flows (login, orders, market data, portfolio).
- **Correct usage** — Request/response patterns, error handling, and API version usage as recommended by Upstox.
- **Reference** — See how to structure `models.PlaceOrderRequest`, historical data params, and other API payloads.

Use these samples to build trading apps, dashboards, or integrations without guessing request shapes or SDK usage.

## Prerequisites

- **Go** 1.23+
- **SDK**: `go get github.com/upstox/upstox-go`
- **Upstox developer account** and API credentials (client ID, client secret, redirect URI).
- **Access token** for authenticated APIs (obtain via [Login API](login/) samples).

For full setup, sandbox mode, and auth flow, see the main [Upstox Go SDK README](../README.md) in the repo root.

## Folder structure

Samples are grouped by API area. Each `.md` file contains one or more complete `package main` Go programs you can run after replacing placeholders like `{your_access_token}` and `{your_client_id}`.

Every program follows the same pattern: build a client with `upstox.NewAPIClient(cfg)`, bind a service with `api.New<Service>Api(client)`, pass request bodies from the `models` package, and unwrap API errors with `errors.As` into an `upstox.GenericSwaggerError`.

| Folder | Description |
|--------|-------------|
| [**login/**](login/) | Authentication: get token from auth code, access-token request, logout. |
| [**user/**](user/) | User profile, fund and margin details, IP management, kill switch. |
| [**orders/**](orders/) | Order lifecycle: place (single/multi, v2 & v3), modify, cancel, order book, order details, order history, trades, historical trades, exit all positions. |
| [**portfolio/**](portfolio/) | Positions, holdings, MTF positions, convert positions. |
| [**market-quote/**](market-quote/) | LTP, full market quotes (v2 & v3), OHLC (v2 & v3), option Greeks. |
| [**historical-data/**](historical-data/) | Historical and intraday candle data (v2 & v3). |
| [**option-chain/**](option-chain/) | Option contracts, put-call option chain. |
| [**expired-instruments/**](expired-instruments/) | Expiries, expired future/option contracts, expired historical candle data. |
| [**market-information/**](market-information/) | Exchange status, market timings, market holidays, OI, change in OI, PCR, max pain, FII, and DII. |
| [**smartlist/**](smartlist/) | Analytics-enriched smartlists for options, futures, and MTF stocks. |
| [**ipo/**](ipo/) | IPO listing (by status/issue type), IPO details by id, and IPO orders — apply, list, fetch by order id and cancel. |
| [**us-stocks/**](us-stocks/) | US stocks: place/modify/cancel orders, order book and details, charge estimate and margin, funds, holdings, and US market status, holidays and timings. |
| [**gtt-orders/**](gtt-orders/) | Place, modify, cancel, and get details for GTT (Good Till Triggered) orders. |
| [**margins/**](margins/) | Margin details. |
| [**charges/**](charges/) | Brokerage details. |
| [**mutual-funds-api/**](mutual-funds-api/) | Mutual fund holdings, orders, order details, and SIPs. |
| [**payments-api/**](payments-api/) | Payin and payout history. |
| [**payout/**](payout/) | Payout history, payout modes, and initiate/modify/cancel payout. |
| [**news/**](news/) | News articles by instrument keys, positions, or holdings. |
| [**fundamentals/**](fundamentals/) | Company profile, balance sheet, cash flow, income statement, key ratios, share holdings, corporate actions, and competitors. |
| [**trade-profit-and-loss/**](trade-profit-and-loss/) | P&amp;L report, report metadata, trade charges. |
| [**strategies/**](strategies/) | Ready-to-run options strategy examples for Nifty 50 (bullish, bearish, neutral, others). |
| [**websocket/**](websocket/) | Live market data (`MarketDataStreamerV3`) and order/portfolio updates (`PortfolioDataStreamer`) over WebSocket. |

### Options Strategies

Each strategy script searches for the required Nifty 50 option legs using the Instruments API (`api.NewInstrumentsApi`) and places market orders via the v3 Order API (`api.NewOrderV3Api`).

#### [Bullish](strategies/bullish/)

| File | Strategy | Legs |
|------|----------|------|
| [buy_call/main.go](strategies/bullish/code/buy_call/main.go) | [**Buy Call**](strategies/bullish/README.md#buy-call--codebuy_callmaingo) | BUY ATM CE |
| [sell_put/main.go](strategies/bullish/code/sell_put/main.go) | [**Sell Put**](strategies/bullish/README.md#sell-put--codesell_putmaingo) | SELL ATM PE |
| [bull_call_spread/main.go](strategies/bullish/code/bull_call_spread/main.go) | [**Bull Call Spread**](strategies/bullish/README.md#bull-call-spread--codebull_call_spreadmaingo) | BUY ATM CE + SELL ATM+1 CE |
| [bull_put_spread/main.go](strategies/bullish/code/bull_put_spread/main.go) | [**Bull Put Spread**](strategies/bullish/README.md#bull-put-spread--codebull_put_spreadmaingo) | SELL ATM PE + BUY ATM-1 PE |
| [bull_butterfly/main.go](strategies/bullish/code/bull_butterfly/main.go) | [**Bull Butterfly**](strategies/bullish/README.md#bull-butterfly--codebull_butterflymaingo) | BUY ATM CE + SELL 2× ATM+1 CE + BUY ATM+2 CE |
| [bull_condor/main.go](strategies/bullish/code/bull_condor/main.go) | [**Bull Condor**](strategies/bullish/README.md#bull-condor--codebull_condormaingo) | BUY ATM CE + SELL ATM+1 CE + SELL ATM+2 CE + BUY ATM+3 CE |
| [long_calendar_call/main.go](strategies/bullish/code/long_calendar_call/main.go) | [**Long Calendar with Calls**](strategies/bullish/README.md#long-calendar-with-calls--codelong_calendar_callmaingo) | SELL current-week ATM CE + BUY next-week ATM CE |
| [long_synthetic_future/main.go](strategies/bullish/code/long_synthetic_future/main.go) | [**Long Synthetic Future**](strategies/bullish/README.md#long-synthetic-future--codelong_synthetic_futuremaingo) | BUY ATM CE + SELL ATM PE |
| [call_ratio_back_spread/main.go](strategies/bullish/code/call_ratio_back_spread/main.go) | [**Call Ratio Back Spread**](strategies/bullish/README.md#call-ratio-back-spread--codecall_ratio_back_spreadmaingo) | SELL 1× ATM CE + BUY 2× ATM+1 CE |
| [range_forward/main.go](strategies/bullish/code/range_forward/main.go) | [**Range Forward**](strategies/bullish/README.md#range-forward--coderange_forwardmaingo) | SELL ATM-1 PE + BUY ATM+1 CE |

#### [Bearish](strategies/bearish/)

| File | Strategy | Legs |
|------|----------|------|
| [buy_put/main.go](strategies/bearish/code/buy_put/main.go) | [**Buy Put**](strategies/bearish/README.md#buy-put--codebuy_putmaingo) | BUY ATM PE |
| [sell_call/main.go](strategies/bearish/code/sell_call/main.go) | [**Sell Call**](strategies/bearish/README.md#sell-call--codesell_callmaingo) | SELL ATM CE |
| [bear_call_spread/main.go](strategies/bearish/code/bear_call_spread/main.go) | [**Bear Call Spread**](strategies/bearish/README.md#bear-call-spread--codebear_call_spreadmaingo) | SELL ATM CE + BUY ATM+1 CE |
| [bear_put_spread/main.go](strategies/bearish/code/bear_put_spread/main.go) | [**Bear Put Spread**](strategies/bearish/README.md#bear-put-spread--codebear_put_spreadmaingo) | BUY ATM PE + SELL ATM-1 PE |
| [bear_butterfly/main.go](strategies/bearish/code/bear_butterfly/main.go) | [**Bear Butterfly**](strategies/bearish/README.md#bear-butterfly--codebear_butterflymaingo) | BUY ATM PE + SELL 2× ATM-1 PE + BUY ATM-2 PE |
| [bear_condor/main.go](strategies/bearish/code/bear_condor/main.go) | [**Bear Condor**](strategies/bearish/README.md#bear-condor--codebear_condormaingo) | BUY ATM PE + SELL ATM-1 PE + SELL ATM-2 PE + BUY ATM-3 PE |
| [long_calendar_put/main.go](strategies/bearish/code/long_calendar_put/main.go) | [**Long Calendar with Puts**](strategies/bearish/README.md#long-calendar-with-puts--codelong_calendar_putmaingo) | SELL current-week ATM PE + BUY next-week ATM PE |
| [short_synthetic_future/main.go](strategies/bearish/code/short_synthetic_future/main.go) | [**Short Synthetic Future**](strategies/bearish/README.md#short-synthetic-future--codeshort_synthetic_futuremaingo) | SELL ATM CE + BUY ATM PE |
| [put_ratio_back_spread/main.go](strategies/bearish/code/put_ratio_back_spread/main.go) | [**Put Ratio Back Spread**](strategies/bearish/README.md#put-ratio-back-spread--codeput_ratio_back_spreadmaingo) | SELL 1× ATM PE + BUY 2× ATM-1 PE |
| [risk_reversal/main.go](strategies/bearish/code/risk_reversal/main.go) | [**Risk Reversal**](strategies/bearish/README.md#risk-reversal--coderisk_reversalmaingo) | SELL ATM+1 CE + BUY ATM-1 PE |

#### [Neutral](strategies/neutral/)

| File | Strategy | Legs |
|------|----------|------|
| [short_straddle/main.go](strategies/neutral/code/short_straddle/main.go) | [**Short Straddle**](strategies/neutral/README.md#short-straddle--codeshort_straddlemaingo) | SELL ATM CE + SELL ATM PE |
| [short_strangle/main.go](strategies/neutral/code/short_strangle/main.go) | [**Short Strangle**](strategies/neutral/README.md#short-strangle--codeshort_stranglemaingo) | SELL ATM+1 CE + SELL ATM-1 PE |
| [iron_butterfly/main.go](strategies/neutral/code/iron_butterfly/main.go) | [**Iron Butterfly**](strategies/neutral/README.md#iron-butterfly--codeiron_butterflymaingo) | SELL ATM CE + SELL ATM PE + BUY ATM+2 CE + BUY ATM-2 PE |
| [batman/main.go](strategies/neutral/code/batman/main.go) | [**Batman**](strategies/neutral/README.md#batman--codebatmanmaingo) | BUY ATM CE + SELL 2× ATM+1 CE + BUY ATM+2 CE + BUY ATM PE + SELL 2× ATM-1 PE + BUY ATM-2 PE |
| [short_iron_condor/main.go](strategies/neutral/code/short_iron_condor/main.go) | [**Short Iron Condor**](strategies/neutral/README.md#short-iron-condor--codeshort_iron_condormaingo) | SELL ATM+1 CE + BUY ATM+2 CE + SELL ATM-1 PE + BUY ATM-2 PE |

#### [Others](strategies/others/)

| File | Strategy | Legs |
|------|----------|------|
| [long_straddle/main.go](strategies/others/code/long_straddle/main.go) | [**Long Straddle**](strategies/others/README.md#long-straddle--codelong_straddlemaingo) | BUY ATM CE + BUY ATM PE |
| [long_strangle/main.go](strategies/others/code/long_strangle/main.go) | [**Long Strangle**](strategies/others/README.md#long-strangle--codelong_stranglemaingo) | BUY ATM+1 CE + BUY ATM-1 PE |
| [call_ratio_spread/main.go](strategies/others/code/call_ratio_spread/main.go) | [**Call Ratio Spread**](strategies/others/README.md#call-ratio-spread--codecall_ratio_spreadmaingo) | BUY 1× ATM CE + SELL 2× ATM+1 CE |
| [put_ratio_spread/main.go](strategies/others/code/put_ratio_spread/main.go) | [**Put Ratio Spread**](strategies/others/README.md#put-ratio-spread--codeput_ratio_spreadmaingo) | BUY 1× ATM PE + SELL 2× ATM-1 PE |
| [long_iron_butterfly/main.go](strategies/others/code/long_iron_butterfly/main.go) | [**Long Iron Butterfly**](strategies/others/README.md#long-iron-butterfly--codelong_iron_butterflymaingo) | BUY ATM CE + BUY ATM PE + SELL ATM+2 CE + SELL ATM-2 PE |
| [long_iron_condor/main.go](strategies/others/code/long_iron_condor/main.go) | [**Long Iron Condor**](strategies/others/README.md#long-iron-condor--codelong_iron_condormaingo) | BUY ATM+1 CE + SELL ATM+2 CE + BUY ATM-1 PE + SELL ATM-2 PE |
| [strip/main.go](strategies/others/code/strip/main.go) | [**Strip**](strategies/others/README.md#strip--codestripmaingo) | BUY 1× ATM CE + BUY 2× ATM PE |
| [strap/main.go](strategies/others/code/strap/main.go) | [**Strap**](strategies/others/README.md#strap--codestrapmaingo) | BUY 2× ATM CE + BUY 1× ATM PE |

## Documentation

- [Upstox API Documentation](https://upstox.com/developer/api-documentation)
- [Upstox Go SDK API reference](../docs/API_REFERENCE.md)
