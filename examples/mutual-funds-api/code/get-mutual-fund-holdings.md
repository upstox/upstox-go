## Get mutual fund holdings

Lists your mutual fund holdings.

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
	mutualFundApi := api.NewMutualFundApi(client)

	resp, _, err := mutualFundApi.GetMutualFundHoldings(context.Background())
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling MutualFundApi.GetMutualFundHoldings:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling MutualFundApi.GetMutualFundHoldings:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```
