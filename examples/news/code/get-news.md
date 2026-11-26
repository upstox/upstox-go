## Get news by instrument keys

News articles for one or more instrument keys.

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
	newsApi := api.NewNewsApi(client)

	resp, _, err := newsApi.GetNews(context.Background(), "instrument_keys", &api.NewsApiGetNewsOpts{InstrumentKeys: optional.NewString("NSE_EQ|INE669E01016")})
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling NewsApi.GetNews:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling NewsApi.GetNews:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```

## Get news by instrument keys with pagination

First page of 10 articles for an instrument key.

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
	newsApi := api.NewNewsApi(client)

	resp, _, err := newsApi.GetNews(context.Background(), "instrument_keys", &api.NewsApiGetNewsOpts{InstrumentKeys: optional.NewString("NSE_EQ|INE669E01016"), PageNumber: optional.NewInt32(1), PageSize: optional.NewInt32(10)})
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling NewsApi.GetNews:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling NewsApi.GetNews:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```

## Get news by positions

News for the instruments in your open positions.

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
	newsApi := api.NewNewsApi(client)

	resp, _, err := newsApi.GetNews(context.Background(), "positions", nil)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling NewsApi.GetNews:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling NewsApi.GetNews:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```

## Get news by holdings

News for the instruments in your holdings.

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
	newsApi := api.NewNewsApi(client)

	resp, _, err := newsApi.GetNews(context.Background(), "holdings", nil)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling NewsApi.GetNews:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling NewsApi.GetNews:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```
