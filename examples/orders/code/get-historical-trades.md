## Get trade history for equity segment

Past equity trades between 2023-04-01 and 2025-03-31, first page of 1000.

```go
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/antihax/optional"
	"github.com/upstox/upstox-go/upstox_client"
	"github.com/upstox/upstox-go/upstox_client/api"
)

func main() {
	cfg := upstox_client.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox_client.NewAPIClient(cfg)
	postTradeApi := api.NewPostTradeApi(client)

	resp, _, err := postTradeApi.GetTradesByDateRange(context.Background(), "2023-04-01", "2025-03-31", 1, 1000, &api.PostTradeApiGetTradesByDateRangeOpts{Segment: optional.NewString("EQ")})
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling PostTradeApi.GetTradesByDateRange:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling PostTradeApi.GetTradesByDateRange:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```

## Get trade history for futures and options segment

Past F&O trades between 2023-04-01 and 2025-03-31, first page of 1000.

```go
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/antihax/optional"
	"github.com/upstox/upstox-go/upstox_client"
	"github.com/upstox/upstox-go/upstox_client/api"
)

func main() {
	cfg := upstox_client.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox_client.NewAPIClient(cfg)
	postTradeApi := api.NewPostTradeApi(client)

	resp, _, err := postTradeApi.GetTradesByDateRange(context.Background(), "2023-04-01", "2025-03-31", 1, 1000, &api.PostTradeApiGetTradesByDateRangeOpts{Segment: optional.NewString("FO")})
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling PostTradeApi.GetTradesByDateRange:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling PostTradeApi.GetTradesByDateRange:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```
