## Get historical candle data with a 1-minute interval

V2 `1minute` candles between 2023-11-12 and 2023-11-13.

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
	historyApi := api.NewHistoryApi(client)

	instrumentKey := "NSE_EQ|INE669E01016"
	interval := "1minute"
	toDate := "2023-11-13"
	fromDate := "2023-11-12"
	resp, _, err := historyApi.GetHistoricalCandleDataRange(context.Background(), instrumentKey, interval, toDate, fromDate)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling HistoryApi.GetHistoricalCandleDataRange:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling HistoryApi.GetHistoricalCandleDataRange:", err)
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

## Get data with a 30-minute interval

V2 `30minute` candles between 2023-11-12 and 2023-11-13.

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
	historyApi := api.NewHistoryApi(client)

	instrumentKey := "NSE_EQ|INE669E01016"
	interval := "30minute"
	toDate := "2023-11-13"
	fromDate := "2023-11-12"
	resp, _, err := historyApi.GetHistoricalCandleDataRange(context.Background(), instrumentKey, interval, toDate, fromDate)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling HistoryApi.GetHistoricalCandleDataRange:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling HistoryApi.GetHistoricalCandleDataRange:", err)
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

## Get data with a daily interval

V2 `day` candles between 2023-11-12 and 2023-11-13.

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
	historyApi := api.NewHistoryApi(client)

	instrumentKey := "NSE_EQ|INE669E01016"
	interval := "day"
	toDate := "2023-11-13"
	fromDate := "2023-11-12"
	resp, _, err := historyApi.GetHistoricalCandleDataRange(context.Background(), instrumentKey, interval, toDate, fromDate)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling HistoryApi.GetHistoricalCandleDataRange:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling HistoryApi.GetHistoricalCandleDataRange:", err)
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

## Get data with a weekly interval

V2 `week` candles between 2023-10-13 and 2023-11-13.

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
	historyApi := api.NewHistoryApi(client)

	instrumentKey := "NSE_EQ|INE669E01016"
	interval := "week"
	toDate := "2023-11-13"
	fromDate := "2023-10-13"
	resp, _, err := historyApi.GetHistoricalCandleDataRange(context.Background(), instrumentKey, interval, toDate, fromDate)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling HistoryApi.GetHistoricalCandleDataRange:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling HistoryApi.GetHistoricalCandleDataRange:", err)
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

## Get data with a monthly interval

V2 `month` candles between 2022-10-13 and 2023-11-13.

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
	historyApi := api.NewHistoryApi(client)

	instrumentKey := "NSE_EQ|INE669E01016"
	interval := "month"
	toDate := "2023-11-13"
	fromDate := "2022-10-13"
	resp, _, err := historyApi.GetHistoricalCandleDataRange(context.Background(), instrumentKey, interval, toDate, fromDate)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling HistoryApi.GetHistoricalCandleDataRange:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling HistoryApi.GetHistoricalCandleDataRange:", err)
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

## Get historical candle data with a 1-minute interval

V2 `1minute` candles up to 2023-11-13 (no from date).

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
	historyApi := api.NewHistoryApi(client)

	instrumentKey := "NSE_EQ|INE669E01016"
	interval := "1minute"
	toDate := "2023-11-13"
	resp, _, err := historyApi.GetHistoricalCandleData(context.Background(), instrumentKey, interval, toDate)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling HistoryApi.GetHistoricalCandleData:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling HistoryApi.GetHistoricalCandleData:", err)
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

## Get data with a 30-minute interval

V2 `30minute` candles up to 2023-11-13 (no from date).

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
	historyApi := api.NewHistoryApi(client)

	instrumentKey := "NSE_EQ|INE669E01016"
	interval := "30minute"
	toDate := "2023-11-13"
	resp, _, err := historyApi.GetHistoricalCandleData(context.Background(), instrumentKey, interval, toDate)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling HistoryApi.GetHistoricalCandleData:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling HistoryApi.GetHistoricalCandleData:", err)
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

## Get data with a daily interval

V2 `day` candles up to 2023-11-13 (no from date).

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
	historyApi := api.NewHistoryApi(client)

	instrumentKey := "NSE_EQ|INE669E01016"
	interval := "day"
	toDate := "2023-11-13"
	resp, _, err := historyApi.GetHistoricalCandleData(context.Background(), instrumentKey, interval, toDate)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling HistoryApi.GetHistoricalCandleData:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling HistoryApi.GetHistoricalCandleData:", err)
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

## Get data with a weekly interval

V2 `week` candles up to 2023-11-13 (no from date).

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
	historyApi := api.NewHistoryApi(client)

	instrumentKey := "NSE_EQ|INE669E01016"
	interval := "week"
	toDate := "2023-11-13"
	resp, _, err := historyApi.GetHistoricalCandleData(context.Background(), instrumentKey, interval, toDate)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling HistoryApi.GetHistoricalCandleData:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling HistoryApi.GetHistoricalCandleData:", err)
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

## Get data with a monthly interval

V2 `month` candles up to 2023-11-13 (no from date).

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
	historyApi := api.NewHistoryApi(client)

	instrumentKey := "NSE_EQ|INE669E01016"
	interval := "month"
	toDate := "2023-11-13"
	resp, _, err := historyApi.GetHistoricalCandleData(context.Background(), instrumentKey, interval, toDate)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling HistoryApi.GetHistoricalCandleData:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling HistoryApi.GetHistoricalCandleData:", err)
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
