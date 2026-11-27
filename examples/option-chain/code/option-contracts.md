## Get option contracts of an instrument key

All option contracts of Nifty 50 across expiries.

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
	optionsApi := api.NewOptionsApi(client)

	resp, _, err := optionsApi.GetOptionContracts(context.Background(), "NSE_INDEX|Nifty 50", nil)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling OptionsApi.GetOptionContracts:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling OptionsApi.GetOptionContracts:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```

## Get option contracts of an instrument key with expiry date

Nifty 50 option contracts for a single expiry.

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
	optionsApi := api.NewOptionsApi(client)

	resp, _, err := optionsApi.GetOptionContracts(context.Background(), "NSE_INDEX|Nifty 50", &api.OptionsApiGetOptionContractsOpts{ExpiryDate: optional.NewString("2024-10-31")})
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling OptionsApi.GetOptionContracts:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling OptionsApi.GetOptionContracts:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```
