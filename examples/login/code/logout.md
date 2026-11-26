## Logout Request

Invalidates the current access token.

```go
package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/upstox/upstox-go/upstox"
	"github.com/upstox/upstox-go/upstox/api"
)

func main() {
	cfg := upstox.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox.NewAPIClient(cfg)
	loginApi := api.NewLoginApi(client)

	resp, _, err := loginApi.Logout(context.Background())
	if err != nil {
		var apiErr upstox.GenericSwaggerError
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
