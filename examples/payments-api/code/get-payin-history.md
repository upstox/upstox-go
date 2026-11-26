## Get Payin History

Lists funds added to your trading account.

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
	userApi := api.NewUserApi(client)

	resp, _, err := userApi.GetPayinHistory(context.Background())
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling UserApi.GetPayinHistory:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling UserApi.GetPayinHistory:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```
