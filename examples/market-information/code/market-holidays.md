## Get market holidays for current year

Lists all exchange holidays of the current year.

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
	client := upstox_client.NewAPIClient(cfg)
	marketHolidaysAndTimingsApi := api.NewMarketHolidaysAndTimingsApi(client)

	resp, _, err := marketHolidaysAndTimingsApi.GetHolidays(context.Background())
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling MarketHolidaysAndTimingsApi.GetHolidays:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling MarketHolidaysAndTimingsApi.GetHolidays:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```

## Get market holiday status of a date

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
	client := upstox_client.NewAPIClient(cfg)
	marketHolidaysAndTimingsApi := api.NewMarketHolidaysAndTimingsApi(client)

	resp, _, err := marketHolidaysAndTimingsApi.GetHoliday(context.Background(), "2024-01-22")
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling MarketHolidaysAndTimingsApi.GetHoliday:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling MarketHolidaysAndTimingsApi.GetHoliday:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```
