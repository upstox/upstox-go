# Portfolio Stream Feed WebSocket Client

This Go program demonstrates how to connect to the Upstox WebSocket API for streaming live order updates. It receives the order updates and prints them to the console.

## Getting Started

These instructions will help you run the sample websocket client.

### Prerequisites

You need Go 1.23 or later. If you haven't installed Go yet, you can download it from the official website:

[Download Go](https://go.dev/dl/)

Add the Upstox Go SDK to your module:

```sh
go get github.com/upstox/upstox-go
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

The program creates an SDK configuration with your access token and a `feeder.PortfolioDataStreamer` with order updates enabled. On connect, the streamer authorizes with the access token and opens the WebSocket connection.

Each order update arrives as a JSON message, which the program prints to the console. Set `Position`, `Holding` or `GTTOrder` in `feeder.PortfolioUpdates` to receive those updates as well.

## Support

If you encounter any problems or have any questions about this project, feel free to post it on our [Developer Community](https://community.upstox.com/c/developer-api/15).

## Disclaimer

This is a sample program meant for educational purposes. It may require modifications to work with your specific requirements.

Please replace `ACCESS_TOKEN` with your actual access token. Modify any other details as needed to fit your project.
