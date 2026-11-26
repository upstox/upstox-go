# Market Stream feed websocket client

This Go program demonstrates how to connect to the Upstox WebSocket API for streaming live market data. It subscribes to a list of instrument keys and prints the decoded protobuf data as JSON.

## Getting Started

These instructions will help you run the sample v3 websocket client.

### Prerequisites

You need Go 1.23 or later. If you haven't installed Go yet, you can download it from the official website:

[Download Go](https://go.dev/dl/)

The program uses the Upstox Go SDK, which brings in `github.com/gorilla/websocket` and `google.golang.org/protobuf`:

```sh
go get github.com/upstox/upstox-go
```

### Protocol Buffers (Protobuf) Classes

The SDK ships the generated Go types for the v3 feed in `github.com/upstox/upstox-go/upstox/feeder/marketdatav3`, so you do not need to run `protoc` yourself.

To regenerate them from the [proto file](https://assets.upstox.com/feed/market-data-feed/v3/MarketDataFeed.proto), install the Protocol Buffers compiler from the [Google Protocol Buffers GitHub repository](https://github.com/protocolbuffers/protobuf/releases) and the Go plugin:

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.6
protoc --go_out=. --go_opt=paths=source_relative MarketDataFeedV3.proto
```

### Configuration

The program requires an Upstox API access token for authorization. Look for the line below in `main.go` and replace `ACCESS_TOKEN` with your actual access token.

```
cfg.AccessToken = "ACCESS_TOKEN"
```

### Running the Program

After setting up your access token, run the following command from this directory:

```
go run .
```

Press Ctrl+C to stop streaming.

## Understanding the Code

The program creates an SDK configuration with your access token and a `feeder.MarketDataStreamerV3` for "NSE_INDEX|Nifty Bank" and "NSE_INDEX|Nifty 50" in `full` mode. On connect, the streamer authorizes with the access token, opens the WebSocket and sends the subscription request.

Each message from the server is decoded into a `marketdatav3.FeedResponse`, which the program prints as JSON with `protojson.Format`. If the connection drops, the streamer reconnects and subscribes again.

## Support

If you encounter any problems or have any questions about this project, feel free to post it on our [Developer Community](https://community.upstox.com/c/developer-api/15).

## Disclaimer

This is a sample program meant for educational purposes. It may require modifications to work with your specific requirements.

Please replace `ACCESS_TOKEN` with your actual access token. Modify any other details as needed to fit your project.
