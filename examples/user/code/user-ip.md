## Get user IPs

Static IPs registered for order placement.

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
	userApi := api.NewUserApi(client)

	resp, _, err := userApi.GetUserIps(context.Background())
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling UserApi.GetUserIps:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling UserApi.GetUserIps:", err)
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

## Update user IP

Sets the primary static IP.

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
	cfg.AccessToken = "{your_access_token}"
	client := upstox_client.NewAPIClient(cfg)
	userApi := api.NewUserApi(client)

	body := models.UpdateUserIpRequest{PrimaryIp: "1.2.3.4"}
	resp, _, err := userApi.UpdateUserIp(context.Background(), body)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling UserApi.UpdateUserIp:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling UserApi.UpdateUserIp:", err)
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
