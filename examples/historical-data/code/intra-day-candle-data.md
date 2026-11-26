## Get intraday candle data with a 1-minute interval

Today's V2 `1minute` candles.

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
	client := upstox.NewAPIClient(cfg)
	historyApi := api.NewHistoryApi(client)

	instrumentKey := "NSE_EQ|INE669E01016"
	interval := "1minute"
	resp, _, err := historyApi.GetIntraDayCandleData(context.Background(), instrumentKey, interval)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling HistoryApi.GetIntraDayCandleData:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling HistoryApi.GetIntraDayCandleData:", err)
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

## Get intraday candle data with a 30-minute interval

Today's V2 `30minute` candles.

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
	client := upstox.NewAPIClient(cfg)
	historyApi := api.NewHistoryApi(client)

	instrumentKey := "NSE_EQ|INE669E01016"
	interval := "30minute"
	resp, _, err := historyApi.GetIntraDayCandleData(context.Background(), instrumentKey, interval)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling HistoryApi.GetIntraDayCandleData:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling HistoryApi.GetIntraDayCandleData:", err)
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
