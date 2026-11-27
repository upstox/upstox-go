## Get US market status

Current open/close status of the US market.

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
	usMarketHolidaysAndTimingsApi := api.NewUSMarketHolidaysAndTimingsApi(client)

	resp, _, err := usMarketHolidaysAndTimingsApi.GetUsMarketStatus(context.Background())
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling USMarketHolidaysAndTimingsApi.GetUsMarketStatus:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling USMarketHolidaysAndTimingsApi.GetUsMarketStatus:", err)
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

## Get US market holidays for the year

Lists all US market holidays of the current year.

```go
package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/upstox/upstox-go/upstox_client"
	"github.com/upstox/upstox-go/upstox_client/api"
)

func main() {
	cfg := upstox_client.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox_client.NewAPIClient(cfg)
	usMarketHolidaysAndTimingsApi := api.NewUSMarketHolidaysAndTimingsApi(client)

	resp, _, err := usMarketHolidaysAndTimingsApi.GetUsHolidays(context.Background())
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling USMarketHolidaysAndTimingsApi.GetUsHolidays:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling USMarketHolidaysAndTimingsApi.GetUsHolidays:", err)
		}
		return
	}
	for _, holiday := range resp.Data {
		fmt.Println(holiday.Date, holiday.HolidayType)
	}
}
```

## Get US market holiday on a date

Holiday details for one date.

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
	usMarketHolidaysAndTimingsApi := api.NewUSMarketHolidaysAndTimingsApi(client)

	// Empty list when the date is not a holiday
	resp, _, err := usMarketHolidaysAndTimingsApi.GetUsHoliday(context.Background(), "2026-11-26")
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling USMarketHolidaysAndTimingsApi.GetUsHoliday:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling USMarketHolidaysAndTimingsApi.GetUsHoliday:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```

## Get US market timings on a date

US trading sessions on one date.

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
	usMarketHolidaysAndTimingsApi := api.NewUSMarketHolidaysAndTimingsApi(client)

	// Regular and overnight sessions; empty list when the market is closed
	resp, _, err := usMarketHolidaysAndTimingsApi.GetUsExchangeTimings(context.Background(), "2026-11-27")
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling USMarketHolidaysAndTimingsApi.GetUsExchangeTimings:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling USMarketHolidaysAndTimingsApi.GetUsExchangeTimings:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```
