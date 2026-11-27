## Get Expiries for given instrument

Lists every past expiry date available for Nifty 50.

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

	resp, _, err := expiredInstrumentApi.GetExpiries(context.Background(), "NSE_INDEX|Nifty 50")
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling ExpiredInstrumentApi.GetExpiries:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling ExpiredInstrumentApi.GetExpiries:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```
