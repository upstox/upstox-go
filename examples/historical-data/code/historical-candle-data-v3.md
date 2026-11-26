## Get data with a 1-minute interval

Historical candles from 2025-01-01 to 2025-01-02 using unit `minutes` and interval `1`.

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
	historyV3Api := api.NewHistoryV3Api(client)

	resp, _, err := historyV3Api.GetHistoricalCandleDataRangeV3(context.Background(), "NSE_EQ|INE848E01016", "minutes", 1, "2025-01-02", "2025-01-01")
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling HistoryV3Api.GetHistoricalCandleDataRangeV3:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling HistoryV3Api.GetHistoricalCandleDataRangeV3:", err)
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

## Get data with a 3-minute interval

Historical candles from 2025-01-01 to 2025-01-02 using unit `minutes` and interval `3`.

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
	historyV3Api := api.NewHistoryV3Api(client)

	resp, _, err := historyV3Api.GetHistoricalCandleDataRangeV3(context.Background(), "NSE_EQ|INE848E01016", "minutes", 3, "2025-01-02", "2025-01-01")
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling HistoryV3Api.GetHistoricalCandleDataRangeV3:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling HistoryV3Api.GetHistoricalCandleDataRangeV3:", err)
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

## Get data with a 15-minute interval

Historical candles from 2025-01-01 to 2025-01-04 using unit `minutes` and interval `15`.

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
	historyV3Api := api.NewHistoryV3Api(client)

	resp, _, err := historyV3Api.GetHistoricalCandleDataRangeV3(context.Background(), "NSE_EQ|INE848E01016", "minutes", 15, "2025-01-04", "2025-01-01")
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling HistoryV3Api.GetHistoricalCandleDataRangeV3:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling HistoryV3Api.GetHistoricalCandleDataRangeV3:", err)
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

## Get data with a 1-hour interval

Historical candles from 2025-01-01 to 2025-02-01 using unit `hours` and interval `1`.

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
	historyV3Api := api.NewHistoryV3Api(client)

	resp, _, err := historyV3Api.GetHistoricalCandleDataRangeV3(context.Background(), "NSE_EQ|INE848E01016", "hours", 1, "2025-02-01", "2025-01-01")
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling HistoryV3Api.GetHistoricalCandleDataRangeV3:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling HistoryV3Api.GetHistoricalCandleDataRangeV3:", err)
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

## Get data with a 4-hour interval

Historical candles from 2025-01-01 to 2025-02-01 using unit `hours` and interval `4`.

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
	historyV3Api := api.NewHistoryV3Api(client)

	resp, _, err := historyV3Api.GetHistoricalCandleDataRangeV3(context.Background(), "NSE_EQ|INE848E01016", "hours", 4, "2025-02-01", "2025-01-01")
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling HistoryV3Api.GetHistoricalCandleDataRangeV3:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling HistoryV3Api.GetHistoricalCandleDataRangeV3:", err)
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

Historical candles from 2025-01-01 to 2025-03-01 using unit `days` and interval `1`.

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
	historyV3Api := api.NewHistoryV3Api(client)

	resp, _, err := historyV3Api.GetHistoricalCandleDataRangeV3(context.Background(), "NSE_EQ|INE848E01016", "days", 1, "2025-03-01", "2025-01-01")
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling HistoryV3Api.GetHistoricalCandleDataRangeV3:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling HistoryV3Api.GetHistoricalCandleDataRangeV3:", err)
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

Historical candles from 2024-01-01 to 2025-01-01 using unit `weeks` and interval `1`.

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
	historyV3Api := api.NewHistoryV3Api(client)

	resp, _, err := historyV3Api.GetHistoricalCandleDataRangeV3(context.Background(), "NSE_EQ|INE848E01016", "weeks", 1, "2025-01-01", "2024-01-01")
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling HistoryV3Api.GetHistoricalCandleDataRangeV3:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling HistoryV3Api.GetHistoricalCandleDataRangeV3:", err)
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

Historical candles from 2010-01-01 to 2025-01-01 using unit `months` and interval `1`.

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
	historyV3Api := api.NewHistoryV3Api(client)

	resp, _, err := historyV3Api.GetHistoricalCandleDataRangeV3(context.Background(), "NSE_EQ|INE848E01016", "months", 1, "2025-01-01", "2010-01-01")
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling HistoryV3Api.GetHistoricalCandleDataRangeV3:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling HistoryV3Api.GetHistoricalCandleDataRangeV3:", err)
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
