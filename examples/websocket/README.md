# WebSocket Feeds

The `feeder` package connects to the Upstox WebSocket feeds for live market data and portfolio updates:

1. **MarketDataStreamerV3**: real-time market updates for the instruments you subscribe to, decoded from protobuf.
2. **PortfolioDataStreamer**: real-time order, position, holding and GTT order updates as JSON.

Both streamers authenticate with the access token on your `upstox_client.Configuration`, reconnect automatically after an abnormal disconnect, and replay your market data subscriptions after every reconnect.

```sh
go get github.com/upstox/upstox-go
```

Runnable samples:

- [market_data/v3](market_data/v3) streams market data for Nifty 50 and Nifty Bank.
- [order_updates](order_updates) streams order updates.

## MarketDataStreamerV3

```go
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/upstox/upstox-go/upstox_client"
	"github.com/upstox/upstox-go/upstox_client/feeder"
	"github.com/upstox/upstox-go/upstox_client/feeder/marketdatav3"
)

func main() {
	cfg := upstox_client.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"

	streamer, err := feeder.NewMarketDataStreamerV3(cfg,
		[]string{"NSE_INDEX|Nifty 50", "NSE_INDEX|Nifty Bank"}, feeder.ModeFull)
	if err != nil {
		log.Fatal(err)
	}
	streamer.OnMessage(func(feed *marketdatav3.FeedResponse) {
		fmt.Println(feed)
	})
	if err := streamer.Connect(); err != nil {
		log.Fatal(err)
	}
	defer streamer.Disconnect()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
}
```

You authenticate with an access token, create the streamer with instrument keys and a mode, and connect. Each update arrives as a `*marketdatav3.FeedResponse`; use `protojson.Format(feed)` from `google.golang.org/protobuf/encoding/protojson` to print it as JSON.

### Modes

- `feeder.ModeLTPC` (`ltpc`): the latest trade only: last traded price, time and quantity, and the previous close.
- `feeder.ModeFull` (`full`): latest trade prices, D5 depth, 1-minute, 30-minute and daily candles, plus additional details.
- `feeder.ModeOptionGreeks` (`option_greeks`): option greeks only.
- `feeder.ModeFullD30` (`full_d30`): full mode data plus 30 market level quotes.

### Methods

1. `NewMarketDataStreamerV3(cfg, instrumentKeys, mode)`: creates the streamer; the keys are subscribed in `mode` on connect. Pass `nil` keys to subscribe later.
2. `Connect()`: opens the WebSocket connection. Returns a `*feeder.HandshakeError` (status code and body) if the server rejects it.
3. `Subscribe(instrumentKeys, mode)`: subscribes the keys in the given mode. Returns `feeder.ErrNotConnected` before `Connect`.
4. `Unsubscribe(instrumentKeys)`: stops updates for the keys.
5. `ChangeMode(instrumentKeys, mode)`: switches already subscribed keys to another mode.
6. `ClearSubscriptions()`: forgets tracked subscriptions so they are not replayed on reconnect.
7. `Disconnect()`: closes the connection and stops auto-reconnect.
8. `AutoReconnect(enable, interval, retryCount)`: configures auto-reconnect. Defaults: enabled, 1 second, 5 retries.

### Callbacks

- `OnOpen(func())`: the connection is established (also after each reconnect).
- `OnClose(func(code int, reason string))`: the connection closed.
- `OnMessage(func(*marketdatav3.FeedResponse))`: a market update arrived.
- `OnError(func(error))`: a connection, decode or reconnect error occurred.
- `OnReconnecting(func(attempt int))`: a reconnect attempt is starting.
- `OnAutoReconnectStopped(func(reason string))`: auto-reconnect stopped (retries exhausted, disabled by the client, or the token was rejected with 401).

Callbacks run on the streamer's read goroutine, so keep them short and hand heavy work to another goroutine. Calling `Subscribe`, `Unsubscribe` or `ChangeMode` from a callback is safe.

### Subscribe, change mode and unsubscribe

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/upstox/upstox-go/upstox_client"
	"github.com/upstox/upstox-go/upstox_client/feeder"
	"github.com/upstox/upstox-go/upstox_client/feeder/marketdatav3"
)

func main() {
	cfg := upstox_client.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"

	streamer, err := feeder.NewMarketDataStreamerV3(cfg, nil, feeder.ModeFull)
	if err != nil {
		log.Fatal(err)
	}
	keys := []string{"NSE_EQ|INE020B01018", "NSE_EQ|INE467B01029"}

	streamer.OnOpen(func() {
		fmt.Println("Connected. Subscribing to instrument keys.")
		if err := streamer.Subscribe(keys, feeder.ModeFull); err != nil {
			log.Println(err)
		}
	})
	streamer.OnMessage(func(feed *marketdatav3.FeedResponse) {
		fmt.Println(feed)
	})
	if err := streamer.Connect(); err != nil {
		log.Fatal(err)
	}
	defer streamer.Disconnect()

	time.Sleep(5 * time.Second)
	fmt.Println("Changing subscription mode to ltpc...")
	if err := streamer.ChangeMode(keys, feeder.ModeLTPC); err != nil {
		log.Println(err)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Unsubscribing from instrument keys.")
	if err := streamer.Unsubscribe(keys); err != nil {
		log.Println(err)
	}
}
```

### Auto-reconnect

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/upstox/upstox-go/upstox_client"
	"github.com/upstox/upstox-go/upstox_client/feeder"
)

func main() {
	cfg := upstox_client.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"

	streamer, err := feeder.NewMarketDataStreamerV3(cfg, []string{"NSE_INDEX|Nifty 50"}, feeder.ModeLTPC)
	if err != nil {
		log.Fatal(err)
	}
	streamer.OnReconnecting(func(attempt int) { fmt.Println("reconnect attempt", attempt) })
	streamer.OnAutoReconnectStopped(func(reason string) { fmt.Println(reason) })

	// Retry every 10 seconds, at most 3 times; use AutoReconnect(false, 0, 0) to disable.
	streamer.AutoReconnect(true, 10*time.Second, 3)

	if err := streamer.Connect(); err != nil {
		log.Fatal(err)
	}
	defer streamer.Disconnect()
	time.Sleep(time.Minute)
}
```

## PortfolioDataStreamer

```go
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/upstox/upstox-go/upstox_client"
	"github.com/upstox/upstox-go/upstox_client/feeder"
)

func main() {
	cfg := upstox_client.NewConfiguration()
	cfg.AccessToken = "{your_access_token}"

	streamer, err := feeder.NewPortfolioDataStreamer(cfg, feeder.PortfolioUpdates{
		Order:    true,
		Position: true,
		Holding:  true,
		GTTOrder: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	streamer.OnOpen(func() { fmt.Println("connection opened") })
	streamer.OnMessage(func(message []byte) { fmt.Println(string(message)) })
	if err := streamer.Connect(); err != nil {
		log.Fatal(err)
	}
	defer streamer.Disconnect()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
}
```

`PortfolioUpdates` selects the update types: `Order`, `Position`, `Holding` and `GTTOrder`. Set at least one; with none set, no `update_types` parameter is sent. Messages are delivered as raw JSON bytes. `Connect`, `Disconnect`, `AutoReconnect` and all callbacks except `OnMessage` behave as for `MarketDataStreamerV3`.
