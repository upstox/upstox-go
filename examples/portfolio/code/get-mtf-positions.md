## Get user MTF positions

Lists your Margin Trade Funding positions.

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

	resp, _, err := portfolioApi.GetMtfPositions(context.Background())
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling PortfolioApi.GetMtfPositions:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling PortfolioApi.GetMtfPositions:", err)
		}
		return
	}
	fmt.Println("MTF positions:")
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```
