## Get brokerage details for equity delivery orders

Brokerage and charges for an equity delivery BUY of 10 shares.

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
	chargeApi := api.NewChargeApi(client)

	resp, _, err := chargeApi.GetBrokerage(context.Background(), "NSE_EQ|INE669E01016", 10, "D", "BUY", 13.4)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling ChargeApi.GetBrokerage:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling ChargeApi.GetBrokerage:", err)
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

## Get brokerage details for equity intraday orders

Brokerage and charges for an equity intraday BUY of 10 shares.

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
	chargeApi := api.NewChargeApi(client)

	resp, _, err := chargeApi.GetBrokerage(context.Background(), "NSE_EQ|INE669E01016", 10, "I", "BUY", 13.4)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling ChargeApi.GetBrokerage:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling ChargeApi.GetBrokerage:", err)
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

## Get brokerage details for equity futures and options delivery orders

Brokerage and charges for an F&O delivery BUY of 10 units.

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
	chargeApi := api.NewChargeApi(client)

	resp, _, err := chargeApi.GetBrokerage(context.Background(), "NSE_FO|35271", 10, "D", "BUY", 1333.4)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling ChargeApi.GetBrokerage:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling ChargeApi.GetBrokerage:", err)
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

## Get brokerage details for equity futures and options intraday orders

Brokerage and charges for an F&O intraday BUY of 10 units.

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
	chargeApi := api.NewChargeApi(client)

	resp, _, err := chargeApi.GetBrokerage(context.Background(), "NSE_FO|35271", 10, "I", "BUY", 1333.4)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling ChargeApi.GetBrokerage:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling ChargeApi.GetBrokerage:", err)
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
