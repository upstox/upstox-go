## Get data with a 1-minute interval

Today's candles using unit `minutes` and interval `1`.

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
	historyV3Api := api.NewHistoryV3Api(client)

	resp, _, err := historyV3Api.GetIntraDayCandleData(context.Background(), "NSE_EQ|INE848E01016", "minutes", 1)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling HistoryV3Api.GetIntraDayCandleData:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling HistoryV3Api.GetIntraDayCandleData:", err)
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

Today's candles using unit `minutes` and interval `3`.

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
	historyV3Api := api.NewHistoryV3Api(client)

	resp, _, err := historyV3Api.GetIntraDayCandleData(context.Background(), "NSE_EQ|INE848E01016", "minutes", 3)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling HistoryV3Api.GetIntraDayCandleData:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling HistoryV3Api.GetIntraDayCandleData:", err)
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

## Get data with a 5-minute interval

Today's candles using unit `minutes` and interval `5`.

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
	historyV3Api := api.NewHistoryV3Api(client)

	resp, _, err := historyV3Api.GetIntraDayCandleData(context.Background(), "NSE_EQ|INE848E01016", "minutes", 5)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling HistoryV3Api.GetIntraDayCandleData:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling HistoryV3Api.GetIntraDayCandleData:", err)
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

Today's candles using unit `minutes` and interval `15`.

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
	historyV3Api := api.NewHistoryV3Api(client)

	resp, _, err := historyV3Api.GetIntraDayCandleData(context.Background(), "NSE_EQ|INE848E01016", "minutes", 15)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling HistoryV3Api.GetIntraDayCandleData:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling HistoryV3Api.GetIntraDayCandleData:", err)
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

Today's candles using unit `minutes` and interval `30`.

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
	historyV3Api := api.NewHistoryV3Api(client)

	resp, _, err := historyV3Api.GetIntraDayCandleData(context.Background(), "NSE_EQ|INE848E01016", "minutes", 30)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling HistoryV3Api.GetIntraDayCandleData:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling HistoryV3Api.GetIntraDayCandleData:", err)
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

Today's candles using unit `hours` and interval `1`.

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
	historyV3Api := api.NewHistoryV3Api(client)

	resp, _, err := historyV3Api.GetIntraDayCandleData(context.Background(), "NSE_EQ|INE848E01016", "hours", 1)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling HistoryV3Api.GetIntraDayCandleData:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling HistoryV3Api.GetIntraDayCandleData:", err)
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

## Get data with a 2-hour interval

Today's candles using unit `hours` and interval `2`.

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
	historyV3Api := api.NewHistoryV3Api(client)

	resp, _, err := historyV3Api.GetIntraDayCandleData(context.Background(), "NSE_EQ|INE848E01016", "hours", 2)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling HistoryV3Api.GetIntraDayCandleData:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling HistoryV3Api.GetIntraDayCandleData:", err)
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

## Get current day data

Today's candles using unit `days` and interval `1`.

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
	historyV3Api := api.NewHistoryV3Api(client)

	resp, _, err := historyV3Api.GetIntraDayCandleData(context.Background(), "NSE_EQ|INE848E01016", "days", 1)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling HistoryV3Api.GetIntraDayCandleData:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling HistoryV3Api.GetIntraDayCandleData:", err)
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
