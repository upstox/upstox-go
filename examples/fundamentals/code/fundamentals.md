## Get Company Profile

Company profile for an ISIN.

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
	fundamentalsApi := api.NewFundamentalsApi(client)

	resp, _, err := fundamentalsApi.GetCompanyProfile(context.Background(), "INE669E01016")
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling FundamentalsApi.GetCompanyProfile:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling FundamentalsApi.GetCompanyProfile:", err)
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

## Get Key Ratios

Key financial ratios for an ISIN.

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
	fundamentalsApi := api.NewFundamentalsApi(client)

	resp, _, err := fundamentalsApi.GetKeyRatios(context.Background(), "INE669E01016")
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling FundamentalsApi.GetKeyRatios:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling FundamentalsApi.GetKeyRatios:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```

## Get Balance Sheet

Balance sheet for an ISIN.

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
	fundamentalsApi := api.NewFundamentalsApi(client)

	resp, _, err := fundamentalsApi.GetBalanceSheet(context.Background(), "INE669E01016", nil)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling FundamentalsApi.GetBalanceSheet:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling FundamentalsApi.GetBalanceSheet:", err)
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

## Get Income Statement

Income statement for an ISIN.

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
	fundamentalsApi := api.NewFundamentalsApi(client)

	resp, _, err := fundamentalsApi.GetIncomeStatement(context.Background(), "INE669E01016", nil)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling FundamentalsApi.GetIncomeStatement:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling FundamentalsApi.GetIncomeStatement:", err)
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

## Get Cash Flow

Cash flow statement for an ISIN.

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
	fundamentalsApi := api.NewFundamentalsApi(client)

	resp, _, err := fundamentalsApi.GetCashFlow(context.Background(), "INE669E01016", nil)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling FundamentalsApi.GetCashFlow:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling FundamentalsApi.GetCashFlow:", err)
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

## Get Competitors

Peer companies for an instrument key.

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
	fundamentalsApi := api.NewFundamentalsApi(client)

	resp, _, err := fundamentalsApi.GetCompetitors(context.Background(), "NSE_EQ|INE669E01016")
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling FundamentalsApi.GetCompetitors:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling FundamentalsApi.GetCompetitors:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```

## Get Corporate Actions

Dividends, splits, bonuses and other corporate actions for an ISIN.

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
	fundamentalsApi := api.NewFundamentalsApi(client)

	resp, _, err := fundamentalsApi.GetCorporateActions(context.Background(), "INE669E01016")
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling FundamentalsApi.GetCorporateActions:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling FundamentalsApi.GetCorporateActions:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```

## Get Share Holdings

Shareholding pattern for an ISIN.

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
	fundamentalsApi := api.NewFundamentalsApi(client)

	resp, _, err := fundamentalsApi.GetShareHoldings(context.Background(), "INE669E01016")
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling FundamentalsApi.GetShareHoldings:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling FundamentalsApi.GetShareHoldings:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```
