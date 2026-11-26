package feeder

import (
	"strings"

	"github.com/upstox/upstox-go/upstox_client"
)

const portfolioStreamFeedURL = "wss://api.upstox.com/v2/feed/portfolio-stream-feed"

// PortfolioUpdates selects which update types the portfolio stream delivers.
type PortfolioUpdates struct {
	Order    bool
	Position bool
	Holding  bool
	GTTOrder bool
}

// PortfolioDataStreamer streams order, position, holding and GTT updates as raw JSON.
type PortfolioDataStreamer struct {
	*Streamer
	onMessage func([]byte)
}

// NewPortfolioDataStreamer creates a portfolio stream for the selected update types.
func NewPortfolioDataStreamer(cfg *upstox_client.Configuration, updates PortfolioUpdates) (*PortfolioDataStreamer, error) {
	base, err := newStreamer(cfg, portfolioURL(updates))
	if err != nil {
		return nil, err
	}
	p := &PortfolioDataStreamer{Streamer: base}
	base.onData = p.handleData
	return p, nil
}

// OnMessage registers the callback that receives each JSON update message.
func (p *PortfolioDataStreamer) OnMessage(fn func([]byte)) {
	p.mu.Lock()
	p.onMessage = fn
	p.mu.Unlock()
}

func (p *PortfolioDataStreamer) handleData(_ int, data []byte) {
	p.mu.Lock()
	fn := p.onMessage
	p.mu.Unlock()
	if fn != nil {
		fn(data)
	}
}

// portfolioURL builds the feed URL; the order of update_types matches the Python SDK.
func portfolioURL(u PortfolioUpdates) string {
	var types []string
	if u.Order {
		types = append(types, "order")
	}
	if u.Holding {
		types = append(types, "holding")
	}
	if u.Position {
		types = append(types, "position")
	}
	if u.GTTOrder {
		types = append(types, "gtt_order")
	}
	if len(types) == 0 {
		return portfolioStreamFeedURL
	}
	return portfolioStreamFeedURL + "?update_types=" + strings.Join(types, "%2C")
}
