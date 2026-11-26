## Get equity and commodity funds

Funds and margin for both segments.

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

	resp, _, err := userApi.GetUserFundMargin(context.Background(), nil)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling UserApi.GetUserFundMargin:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling UserApi.GetUserFundMargin:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```

## Get equity funds

Funds and margin for the securities (SEC) segment.

```go
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/antihax/optional"
	"github.com/upstox/upstox-go/upstox"
	"github.com/upstox/upstox-go/upstox/api"
)

func main() {
	cfg := upstox.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox.NewAPIClient(cfg)
	userApi := api.NewUserApi(client)

	resp, _, err := userApi.GetUserFundMargin(context.Background(), &api.UserApiGetUserFundMarginOpts{Segment: optional.NewString("SEC")})
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling UserApi.GetUserFundMargin:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling UserApi.GetUserFundMargin:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```

## Get commodity funds

Funds and margin for the commodity (COM) segment.

```go
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/antihax/optional"
	"github.com/upstox/upstox-go/upstox"
	"github.com/upstox/upstox-go/upstox/api"
)

func main() {
	cfg := upstox.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox.NewAPIClient(cfg)
	userApi := api.NewUserApi(client)

	resp, _, err := userApi.GetUserFundMargin(context.Background(), &api.UserApiGetUserFundMarginOpts{Segment: optional.NewString("COM")})
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling UserApi.GetUserFundMargin:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling UserApi.GetUserFundMargin:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```
