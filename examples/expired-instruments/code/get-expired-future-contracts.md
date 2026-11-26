## Get Expired Future Contracts for given instrument with expiry date

Lists the expired future contracts of Nifty 50 for one expiry date.

```go
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/upstox/upstox-go/upstox"
	"github.com/upstox/upstox-go/upstox/api"
)

func main() {
	cfg := upstox.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox.NewAPIClient(cfg)
	expiredInstrumentApi := api.NewExpiredInstrumentApi(client)

	resp, _, err := expiredInstrumentApi.GetExpiredFutureContracts(context.Background(), "NSE_INDEX|Nifty 50", "2024-11-27")
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling ExpiredInstrumentApi.GetExpiredFutureContracts:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling ExpiredInstrumentApi.GetExpiredFutureContracts:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```
