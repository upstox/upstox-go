## Get Expiries for given instrument

Lists every past expiry date available for Nifty 50.

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

	resp, _, err := expiredInstrumentApi.GetExpiriesResponse(context.Background(), "NSE_INDEX|Nifty 50")
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling ExpiredInstrumentApi.GetExpiriesResponse:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling ExpiredInstrumentApi.GetExpiriesResponse:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```
