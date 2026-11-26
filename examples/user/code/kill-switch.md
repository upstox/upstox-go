## Get kill switch status

Kill switch state of every segment.

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

	resp, _, err := userApi.GetKillSwitch(context.Background())
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling UserApi.GetKillSwitch:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling UserApi.GetKillSwitch:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```

## Update kill switch

Disables trading in the NSE_EQ and NSE_FO segments.

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

	body := []models.KillSwitchSegmentUpdateRequest{
		{Segment: "NSE_EQ", Action: "DISABLE"},
		{Segment: "NSE_FO", Action: "DISABLE"},
	}
	resp, _, err := userApi.UpdateKillSwitch(context.Background(), body)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling UserApi.UpdateKillSwitch:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling UserApi.UpdateKillSwitch:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```
