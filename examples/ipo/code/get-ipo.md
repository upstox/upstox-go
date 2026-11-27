## Get IPO listing

Lists IPOs.

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
	ipoApi := api.NewIPOApi(client)

	resp, _, err := ipoApi.GetIpoListing(context.Background(), nil)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling IPOApi.GetIpoListing:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling IPOApi.GetIpoListing:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```

## Get IPO listing with filters and pagination

Lists open regular (non-SME) IPOs, 20 per page.

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
	ipoApi := api.NewIPOApi(client)

	// Status: open | closed | listed | upcoming; IssueType: regular | sme
	resp, _, err := ipoApi.GetIpoListing(context.Background(), &api.IPOApiGetIpoListingOpts{Status: optional.NewString("open"), IssueType: optional.NewString("regular"), PageNumber: optional.NewString("1"), Records: optional.NewString("20")})
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling IPOApi.GetIpoListing:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling IPOApi.GetIpoListing:", err)
		}
		return
	}
	out, _ := json.MarshalIndent(resp.Data, "", "  ")
	fmt.Println(string(out))
}
```

## Get IPO details by id

Fetches price band, lot size and dates for one IPO.

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
	ipoApi := api.NewIPOApi(client)

	// The id is the IPO slug id returned by GetIpoListing
	ipoId := "{ipo_slug_id}"
	resp, _, err := ipoApi.GetIpoDetails(context.Background(), ipoId)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling IPOApi.GetIpoDetails:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling IPOApi.GetIpoDetails:", err)
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
