## Get user positions

Lists your open short-term positions.

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
	portfolioApi := api.NewPortfolioApi(client)

	resp, _, err := portfolioApi.GetPositions(context.Background())
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling PortfolioApi.GetPositions:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling PortfolioApi.GetPositions:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```
