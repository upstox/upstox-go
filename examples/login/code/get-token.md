## Get access token using auth code

Exchanges the authorization code from the login redirect for an access token.

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
	client := upstox.NewAPIClient(cfg)
	loginApi := api.NewLoginApi(client)

	code := "{your_auth_code}"
	clientId := "{your_client_id}"
	clientSecret := "{your_client_secret}"
	redirectUri := "{your_redirect_url}"
	grantType := "authorization_code"
	resp, _, err := loginApi.Token(context.Background(), code, clientId, clientSecret, redirectUri, grantType, false)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling LoginApi.Token:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling LoginApi.Token:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(out))
}
```
