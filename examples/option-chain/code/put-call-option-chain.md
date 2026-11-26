## Get put/call option chain

Put and call legs per strike for one Nifty 50 expiry.

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
	optionsApi := api.NewOptionsApi(client)

	resp, _, err := optionsApi.GetPutCallOptionChain(context.Background(), "NSE_INDEX|Nifty 50", "2024-10-31")
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling OptionsApi.GetPutCallOptionChain:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling OptionsApi.GetPutCallOptionChain:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```
