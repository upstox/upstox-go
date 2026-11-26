## Get OI Data

Open interest of Nifty 50 options for one expiry on one date.

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
	marketHolidaysAndTimingsApi := api.NewMarketHolidaysAndTimingsApi(client)

	resp, _, err := marketHolidaysAndTimingsApi.GetOiData(context.Background(), "NSE_INDEX|Nifty 50", "2026-05-29", "2026-05-12")
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling MarketHolidaysAndTimingsApi.GetOiData:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling MarketHolidaysAndTimingsApi.GetOiData:", err)
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

## Get Change in OI Data

Change in open interest over 5-minute intervals.

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
	marketHolidaysAndTimingsApi := api.NewMarketHolidaysAndTimingsApi(client)

	resp, _, err := marketHolidaysAndTimingsApi.GetChangeOiData(context.Background(), "NSE_INDEX|Nifty 50", "2026-05-29", "2026-05-12", 5)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling MarketHolidaysAndTimingsApi.GetChangeOiData:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling MarketHolidaysAndTimingsApi.GetChangeOiData:", err)
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

## Get PCR Data

Put-call ratio in 30-minute buckets.

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
	marketHolidaysAndTimingsApi := api.NewMarketHolidaysAndTimingsApi(client)

	resp, _, err := marketHolidaysAndTimingsApi.GetPcrData(context.Background(), "NSE_INDEX|Nifty 50", "2026-05-29", "2026-05-12", 30)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling MarketHolidaysAndTimingsApi.GetPcrData:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling MarketHolidaysAndTimingsApi.GetPcrData:", err)
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

## Get Max Pain Data

Max pain strike in 30-minute buckets.

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
	marketHolidaysAndTimingsApi := api.NewMarketHolidaysAndTimingsApi(client)

	resp, _, err := marketHolidaysAndTimingsApi.GetMaxPainData(context.Background(), "NSE_INDEX|Nifty 50", "2026-05-29", "2026-05-12", 30)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling MarketHolidaysAndTimingsApi.GetMaxPainData:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling MarketHolidaysAndTimingsApi.GetMaxPainData:", err)
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

## Get FII Data

Daily FII activity in the NSE cash segment.

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
	marketHolidaysAndTimingsApi := api.NewMarketHolidaysAndTimingsApi(client)

	resp, _, err := marketHolidaysAndTimingsApi.GetFiiData(context.Background(), "NSE_EQ|CASH", "1D", nil)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling MarketHolidaysAndTimingsApi.GetFiiData:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling MarketHolidaysAndTimingsApi.GetFiiData:", err)
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

## Get FII Data with Date Filter

Monthly FII activity in index futures since 2026-01-01.

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
	marketHolidaysAndTimingsApi := api.NewMarketHolidaysAndTimingsApi(client)

	resp, _, err := marketHolidaysAndTimingsApi.GetFiiData(context.Background(), "NSE_FO|INDEX_FUTURES", "1M", &api.MarketHolidaysAndTimingsApiGetFiiDataOpts{From: optional.NewString("2026-01-01")})
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling MarketHolidaysAndTimingsApi.GetFiiData:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling MarketHolidaysAndTimingsApi.GetFiiData:", err)
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

## Get DII Data

Daily DII activity in the NSE cash segment.

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
	marketHolidaysAndTimingsApi := api.NewMarketHolidaysAndTimingsApi(client)

	resp, _, err := marketHolidaysAndTimingsApi.GetDiiData(context.Background(), "NSE_EQ|CASH", "1D", nil)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling MarketHolidaysAndTimingsApi.GetDiiData:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling MarketHolidaysAndTimingsApi.GetDiiData:", err)
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
