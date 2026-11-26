## Apply for an IPO

Submits a single-bid IPO application paid through a UPI mandate.

```go
package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/upstox/upstox-go/upstox"
	"github.com/upstox/upstox-go/upstox/api"
	"github.com/upstox/upstox-go/upstox/models"
)

func main() {
	cfg := upstox.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox.NewAPIClient(cfg)
	ipoApi := api.NewIPOApi(client)

	// Id is the IPO slug id returned by GetIpoListing; Category is IND or HNI; 1 to 3 bids
	body := models.IpoApplyRequest{
		Id:       "{ipo_slug_id}",
		Upi:      "{your_upi_id}",
		Category: "IND",
		Bids: []models.IpoBidRequest{
			{Quantity: 10, Price: 150.0},
		},
	}
	resp, _, err := ipoApi.ApplyForIpo(context.Background(), body)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling IPOApi.ApplyForIpo:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling IPOApi.ApplyForIpo:", err)
		}
		return
	}
	if resp.Data == nil {
		fmt.Println("no data returned")
		return
	}
	fmt.Println("IPO order id:", resp.Data.OrderId)
}
```

## Apply for an IPO with multiple bids

Submits an HNI application with three bids at different prices.

```go
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/upstox/upstox-go/upstox"
	"github.com/upstox/upstox-go/upstox/api"
	"github.com/upstox/upstox-go/upstox/models"
)

func main() {
	cfg := upstox.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox.NewAPIClient(cfg)
	ipoApi := api.NewIPOApi(client)

	// Build valid bids from the cut-off price and lot size in GetIpoDetails
	body := models.IpoApplyRequest{
		Id:       "{ipo_slug_id}",
		Upi:      "{your_upi_id}",
		Category: "HNI",
		Bids: []models.IpoBidRequest{
			{Quantity: 10, Price: 150.0},
			{Quantity: 20, Price: 155.0},
			{Quantity: 30, Price: 160.0},
		},
	}
	resp, _, err := ipoApi.ApplyForIpo(context.Background(), body)
	if err != nil {
		var apiErr upstox.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling IPOApi.ApplyForIpo:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling IPOApi.ApplyForIpo:", err)
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

> Note: after a successful application, approve the UPI mandate request in your
> UPI app to block the funds. Until the mandate is approved, the order remains
> pending.
