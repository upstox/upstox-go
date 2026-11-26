// Command v3 streams live market data from the Upstox v3 market data feed and prints it as JSON.
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/upstox/upstox-go/upstox"
	"github.com/upstox/upstox-go/upstox/feeder"
	"github.com/upstox/upstox-go/upstox/feeder/marketdatav3"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	cfg := upstox.NewConfiguration()
	cfg.AccessToken = "ACCESS_TOKEN"

	streamer, err := feeder.NewMarketDataStreamerV3(cfg,
		[]string{"NSE_INDEX|Nifty Bank", "NSE_INDEX|Nifty 50"}, feeder.ModeFull)
	if err != nil {
		log.Fatal(err)
	}

	streamer.OnOpen(func() { fmt.Println("Connection established") })
	streamer.OnMessage(func(feed *marketdatav3.FeedResponse) {
		fmt.Println(protojson.Format(feed))
	})
	streamer.OnError(func(err error) { log.Println("error:", err) })
	streamer.OnClose(func(code int, reason string) { log.Printf("closed: %d %s", code, reason) })
	streamer.OnReconnecting(func(attempt int) { log.Printf("reconnecting, attempt %d", attempt) })
	streamer.OnAutoReconnectStopped(func(reason string) { log.Println("auto reconnect stopped:", reason) })
	streamer.AutoReconnect(true, 5*time.Second, 10)

	if err := streamer.Connect(); err != nil {
		log.Fatal(err)
	}
	defer streamer.Disconnect()

	// Stream until Ctrl+C.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
}
