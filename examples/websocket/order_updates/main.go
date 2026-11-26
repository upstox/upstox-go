// Command order_updates streams live order updates from the Upstox portfolio stream feed.
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/upstox/upstox-go/upstox"
	"github.com/upstox/upstox-go/upstox/feeder"
)

func main() {
	cfg := upstox.NewConfiguration()
	cfg.AccessToken = "ACCESS_TOKEN"

	streamer, err := feeder.NewPortfolioDataStreamer(cfg, feeder.PortfolioUpdates{Order: true})
	if err != nil {
		log.Fatal(err)
	}

	streamer.OnOpen(func() { fmt.Println("Connection established") })
	streamer.OnMessage(func(message []byte) { fmt.Println(string(message)) })
	streamer.OnError(func(err error) { log.Println("error:", err) })
	streamer.OnClose(func(code int, reason string) { log.Printf("closed: %d %s", code, reason) })

	if err := streamer.Connect(); err != nil {
		log.Fatal(err)
	}
	defer streamer.Disconnect()

	// Stream until Ctrl+C.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
}
