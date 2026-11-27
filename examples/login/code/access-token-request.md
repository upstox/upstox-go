## Access token request

Starts the access-token request flow for an individual (indie) user.

```go
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/upstox/upstox-go/upstox_client"
	"github.com/upstox/upstox-go/upstox_client/api"
	"github.com/upstox/upstox-go/upstox_client/models"
)

func main() {
	cfg := upstox_client.NewConfiguration()
	client := upstox_client.NewAPIClient(cfg)
	loginApi := api.NewLoginApi(client)

	body := models.IndieUserTokenRequest{ClientSecret: "****"}
	resp, _, err := loginApi.InitTokenRequestForIndieUser(context.Background(), body, "*****")
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
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
