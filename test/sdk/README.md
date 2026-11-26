# Live SDK sanity suite

These tests call the real Upstox API through the Go SDK. Every REST endpoint in `upstox/api` is called at least once. The suite is the Go counterpart of the Python SDK's `test/sdk_tests/`.

Every file carries the `sdk` build tag, so `go test ./...` never runs the suite. When `UPSTOX_ACCESS_TOKEN` is unset, every test skips.

## Running

```bash
# Read-only checks plus safe probes (all market data, account and report endpoints)
UPSTOX_ACCESS_TOKEN=... go test -tags sdk ./test/sdk/ -v

# One service
UPSTOX_ACCESS_TOKEN=... go test -tags sdk ./test/sdk/ -v -run 'TestGetOptionContracts|TestGetPutCallOptionChain'

# Also place, modify and cancel real orders
UPSTOX_ACCESS_TOKEN=... UPSTOX_RUN_ORDER_TESTS=1 go test -tags sdk ./test/sdk/ -v

# Sandbox order tests only
UPSTOX_SANDBOX_ACCESS_TOKEN=... go test -tags sdk ./test/sdk/ -v -run TestSandbox
```

Check that the suite compiles and skips cleanly without credentials:

```bash
go vet -tags sdk ./test/sdk/
env -u UPSTOX_ACCESS_TOKEN go test -tags sdk ./test/sdk/ -v
```

## Environment variables

| Variable | Needed by | Effect |
| --- | --- | --- |
| `UPSTOX_ACCESS_TOKEN` | every live test | Bearer token for `api.upstox.com`. Tests skip when it is unset. |
| `UPSTOX_SANDBOX_ACCESS_TOKEN` | `TestSandbox*` | Token for the sandbox environment. |
| `UPSTOX_RUN_ORDER_TESTS=1` | order, GTT, multi-order, US order and convert-position tests | Places real orders. See [Order safety](#order-safety). |
| `UPSTOX_RUN_EXIT_POSITIONS=1` | `TestExitPositionsByTag` | Squares off open positions from orders tagged `sdk-go-test`. |
| `UPSTOX_RUN_PAYOUT_TESTS=1` | `TestInitiateModifyCancelPayout` | Requests a real ₹1 withdrawal, changes it to ₹2, then cancels it. |
| `UPSTOX_PAYOUT_MODE` | payout test | Withdrawal mode. Defaults to `IMPS`. |
| `UPSTOX_TEST_IPO_ID`, `UPSTOX_TEST_UPI` | `TestApplyAndCancelIpo` | Applies for that IPO through the UPI handle at the lowest bid, then cancels. This blocks funds through a real UPI mandate. |
| `UPSTOX_RUN_ACCOUNT_SETTINGS_TESTS=1` | `TestUpdateUserIp`, `TestUpdateKillSwitch` | Writes the current static IP and kill-switch state back unchanged. |
| `UPSTOX_RUN_LOGOUT_TEST=1` | `TestLogout` | Invalidates `UPSTOX_ACCESS_TOKEN`. Run it on its own with `-run TestLogout`. |
| `UPSTOX_CLIENT_ID`, `UPSTOX_REDIRECT_URI` | `TestAuthorize` | App credentials for the authorization-dialog redirect check. |
| `UPSTOX_CLIENT_ID`, `UPSTOX_CLIENT_SECRET`, `UPSTOX_REDIRECT_URI`, `UPSTOX_AUTH_CODE` | `TestToken` | Exchanges an auth code for a token. Each code works only once. |
| `UPSTOX_ALGO_NAME` | order writes | Value sent in the `X-Algo-Name` header. Defaults to `sdk-go-test`. |
| `UPSTOX_TEST_US_INSTRUMENT` | US stock tests | US instrument key. Defaults to `US_EQ\|COUR`. |

## What runs without a gate

When only `UPSTOX_ACCESS_TOKEN` is set, the suite runs:

- **Reads:** every GET endpoint. This covers quotes, history, options, instruments, analytics, fundamentals, news, the portfolio, funds, reports, IPO and mutual fund listings, US market data, and the websocket authorize and redirect endpoints.
- **Charge calculators:** brokerage, margin, US order estimate and US margin. These place nothing.
- **Not-found probes:** modify and cancel calls against order, GTT, payout, IPO and US order IDs that do not exist. These exercise the write path with no side effect. The test passes when the API rejects the ID, either with the documented error code (for example `UDAPI100010`) or with any structured 4xx error.
- **Invalid-credential login probes:** `Token`, `InitTokenRequestForIndieUser` and `GenerateTokenViaTotpAndPin`, sent with dummy credentials. The TOTP probe sends no `X-User-Id`, so it cannot count against a real user's PIN attempts.

## Order safety

With `UPSTOX_RUN_ORDER_TESTS=1`, each test follows these rules:

- **Small, far-from-market orders:** each order is a 1-share LIMIT BUY of Vodafone Idea (`NSE_EQ|INE669E01016`), priced 5% below the live LTP. Outside market hours the order is sent as an AMO.
- **Tagged orders:** every order is tagged `sdk-go-test`. The bulk cancel and exit tests filter on that tag, so they never touch other orders.
- **Cleanup:** every order, GTT or US order a test places is cancelled in `t.Cleanup`, even when the test fails.
- **GTT triggers:** GTT entries trigger only after a 20% fall, so they do not fire during the test.
- **US orders:** the US order is a 1-share LIMIT BUY at half the estimated price.
- **Position conversion:** `TestConvertPositions` converts a 1-share intraday Yes Bank position, if one exists.

## Pass criteria

- **`check`:** fails on any error, or on a response whose `status` is not `success`.
- **`checkAllow`:** also accepts the listed business error codes, read from the error body's `errors[].errorCode` (or `error_code`).
- **`checkRejected`:** requires a 4xx response. It is used for deliberately invalid calls.
- **`checkRedirect`:** requires a 2xx or 3xx response. It is used for redirect endpoints, which the client is set not to follow.

Some endpoints depend on the account. For example, US stocks, MTF and expired-instrument data need those features enabled. On an account without them, the matching tests fail with the API's error body printed.
