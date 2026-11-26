## Get user holdings

Lists your long-term holdings.

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
	portfolioApi := api.NewPortfolioApi(client)

	resp, _, err := portfolioApi.GetHoldings(context.Background())
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling PortfolioApi.GetHoldings:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling PortfolioApi.GetHoldings:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```
