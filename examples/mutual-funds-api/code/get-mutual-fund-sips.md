## Get all mutual fund SIPs

Lists your SIPs.

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
	mutualFundApi := api.NewMutualFundApi(client)

	resp, _, err := mutualFundApi.GetMutualFundSips(context.Background(), nil)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling MutualFundApi.GetMutualFundSips:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling MutualFundApi.GetMutualFundSips:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```

## Get mutual fund SIPs with pagination

Lists the first page of 10 SIPs.

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
	mutualFundApi := api.NewMutualFundApi(client)

	resp, _, err := mutualFundApi.GetMutualFundSips(context.Background(), &api.MutualFundApiGetMutualFundSipsOpts{PageNumber: optional.NewInt32(1), Records: optional.NewInt32(10)})
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling MutualFundApi.GetMutualFundSips:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling MutualFundApi.GetMutualFundSips:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```
