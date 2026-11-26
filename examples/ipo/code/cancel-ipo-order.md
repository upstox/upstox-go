## Cancel an IPO order

Withdraws an IPO application by its order id.

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
	ipoApi := api.NewIPOApi(client)

	// The order id comes from ApplyForIpo or GetIpoOrders
	orderId := "{ipo_order_id}"
	resp, _, err := ipoApi.CancelIpoOrder(context.Background(), orderId)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling IPOApi.CancelIpoOrder:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling IPOApi.CancelIpoOrder:", err)
		}
		return
	}
	if resp.Data == nil {
		fmt.Println("no data returned")
		return
	}
	fmt.Println("cancelled order id:", resp.Data.OrderId)
	fmt.Println("status:", resp.Data.Status)
}
```

## Apply, then cancel the same IPO order

Applies for an IPO and immediately withdraws that application.

```go
package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/upstox/upstox-go/upstox_client"
	"github.com/upstox/upstox-go/upstox_client/api"
	"github.com/upstox/upstox-go/upstox_client/models"
)

func main() {
	cfg := upstox_client.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"
	client := upstox_client.NewAPIClient(cfg)
	ipoApi := api.NewIPOApi(client)

	body := models.IpoApplyRequest{
		Id:       "{ipo_slug_id}",
		Upi:      "{your_upi_id}",
		Category: "IND",
		Bids: []models.IpoBidRequest{
			{Quantity: 10, Price: 150.0},
		},
	}
	applyResp, _, err := ipoApi.ApplyForIpo(context.Background(), body)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling IPOApi.ApplyForIpo:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling IPOApi.ApplyForIpo:", err)
		}
		return
	}
	if applyResp.Data == nil {
		fmt.Println("no data returned")
		return
	}
	orderId := applyResp.Data.OrderId
	fmt.Println("applied, order id:", orderId)

	cancelResp, _, err := ipoApi.CancelIpoOrder(context.Background(), orderId)
	if err != nil {
		var apiErr upstox_client.GenericSwaggerError
		if errors.As(err, &apiErr) {
			fmt.Println("Exception when calling IPOApi.CancelIpoOrder:", string(apiErr.Body()))
		} else {
			fmt.Println("Exception when calling IPOApi.CancelIpoOrder:", err)
		}
		return
	}
	if cancelResp.Data != nil {
		fmt.Println("cancelled:", cancelResp.Data.Status)
	}
}
```

> Note: cancellation is only accepted while the IPO bidding window is still
> open. Once the issue closes, the order can no longer be withdrawn — check
> `BiddingEndDate` from `GetIpoDetails` before cancelling.
