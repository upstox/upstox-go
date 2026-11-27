## Get Historical Candle Data for Expired Instruments

Fetches 1-minute candles for an expired instrument key (`<segment>|<token>|<expiry>`).

```go
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/upstox/upstox-go/upstox_client"
	"github.com/upstox/upstox-go/upstox_client/api"
)

func main() {
	cfg := upstox_client.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox_client.NewAPIClient(cfg)
	expiredInstrumentApi := api.NewExpiredInstrumentApi(client)

	resp, _, err := expiredInstrumentApi.GetExpiredHistoricalCandleData(context.Background(), "NSE_FO|54452|24-04-2025", "1minute", "2025-04-24", "2025-04-24")
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling ExpiredInstrumentApi.GetExpiredHistoricalCandleData:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling ExpiredInstrumentApi.GetExpiredHistoricalCandleData:", err)
		}
		return
	}
	if resp.Data == nil {
		fmt.Println("no data returned")
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```
