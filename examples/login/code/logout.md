## Logout Request

Invalidates the current access token.

```go
package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/upstox/upstox-go/upstox_client"
	"github.com/upstox/upstox-go/upstox_client/api"
)

func main() {
	cfg := upstox_client.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox_client.NewAPIClient(cfg)
	loginApi := api.NewLoginApi(client)

	resp, _, err := loginApi.Logout(context.Background())
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling LoginApi.Logout:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling LoginApi.Logout:", err)
		}
		return
	}
	fmt.Println("logged out:", resp.Data)
}
```
