## Get payout history

Lists withdrawals from your trading account.

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

	resp, _, err := userApi.GetPayoutHistory(context.Background())
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling UserApi.GetPayoutHistory:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling UserApi.GetPayoutHistory:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```

## Get payout modes

Lists the payout modes available to you.

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

	resp, _, err := userApi.GetPayoutModes(context.Background())
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling UserApi.GetPayoutModes:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling UserApi.GetPayoutModes:", err)
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

## Initiate payout

Requests a withdrawal of 1000 via IMPS (moves real funds).

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
	cfg.AccessToken = "{your_access_token}"
	client := upstox.NewAPIClient(cfg)
	userApi := api.NewUserApi(client)

	// Mode is one of the values returned by GetPayoutModes (e.g. IMPS)
	body := models.InitiatePayoutRequest{Mode: "IMPS", Amount: 1000.0}
	resp, _, err := userApi.InitiatePayout(context.Background(), body)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling UserApi.InitiatePayout:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling UserApi.InitiatePayout:", err)
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

## Modify payout

Changes the amount of a pending withdrawal.

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
	cfg.AccessToken = "{your_access_token}"
	client := upstox.NewAPIClient(cfg)
	userApi := api.NewUserApi(client)

	transactionId := "{transaction_id}"
	body := models.ModifyPayoutRequest{Amount: 2000.0}
	resp, _, err := userApi.ModifyPayout(context.Background(), body, transactionId)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling UserApi.ModifyPayout:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling UserApi.ModifyPayout:", err)
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

## Cancel payout

Cancels a pending withdrawal.

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

	transactionId := "{transaction_id}"
	resp, _, err := userApi.CancelPayout(context.Background(), transactionId)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling UserApi.CancelPayout:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling UserApi.CancelPayout:", err)
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
