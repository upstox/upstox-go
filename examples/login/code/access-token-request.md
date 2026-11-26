## Access token request

Starts the access-token request flow for an individual (indie) user.

```go
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/upstox/upstox-go/upstox"
	"github.com/upstox/upstox-go/upstox/api"
	"github.com/upstox/upstox-go/upstox/models"
)

func main() {
	cfg := upstox.NewConfiguration()
	client := upstox.NewAPIClient(cfg)
	loginApi := api.NewLoginApi(client)

	body := models.IndieUserTokenRequest{ClientSecret: "****"}
	resp, _, err := loginApi.InitTokenRequestForIndieUser(context.Background(), body, "*****")
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling LoginApi.InitTokenRequestForIndieUser:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling LoginApi.InitTokenRequestForIndieUser:", err)
		}
		return
	}
	if resp.Data == nil {
		fmt.Println("no data returned")
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```
