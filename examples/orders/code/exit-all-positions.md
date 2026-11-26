## Exit all the open positions

Squares off every open position.

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
	orderApi := api.NewOrderApi(client)

	resp, _, err := orderApi.ExitPositions(context.Background(), nil)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling OrderApi.ExitPositions:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling OrderApi.ExitPositions:", err)
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

## Exit all the open positions for a given segment

Squares off positions in the NSE_FO segment only.

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
	orderApi := api.NewOrderApi(client)

	resp, _, err := orderApi.ExitPositions(context.Background(), &api.OrderApiExitPositionsOpts{Segment: optional.NewString("NSE_FO")})
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling OrderApi.ExitPositions:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling OrderApi.ExitPositions:", err)
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

## Exit all the open positions for a given tag

Squares off positions opened with tag `xyz`.

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
	orderApi := api.NewOrderApi(client)

	resp, _, err := orderApi.ExitPositions(context.Background(), &api.OrderApiExitPositionsOpts{Tag: optional.NewString("xyz")})
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling OrderApi.ExitPositions:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling OrderApi.ExitPositions:", err)
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
