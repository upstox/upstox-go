package feeder

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/gorilla/websocket"
	"github.com/upstox/upstox-go/upstox_client"
	"github.com/upstox/upstox-go/upstox_client/feeder/marketdatav3"
	"google.golang.org/protobuf/proto"
)

const marketDataFeedV3URL = "wss://api.upstox.com/v3/feed/market-data-feed"

// Mode is the market data subscription mode.
type Mode string

// Market data feed modes.
const (
	ModeLTPC         Mode = "ltpc"
	ModeFull         Mode = "full"
	ModeOptionGreeks Mode = "option_greeks"
	ModeFullD30      Mode = "full_d30"
)

var allModes = []Mode{ModeLTPC, ModeFull, ModeOptionGreeks, ModeFullD30}

func (m Mode) valid() bool {
	for _, v := range allModes {
		if m == v {
			return true
		}
	}
	return false
}

const (
	methodSubscribe   = "sub"
	methodChangeMode  = "change_mode"
	methodUnsubscribe = "unsub"
)

// MarketDataStreamerV3 streams decoded protobuf market data from the v3 market data feed.
type MarketDataStreamerV3 struct {
	*Streamer

	// subs is guarded by Streamer.mu.
	subs      map[Mode]map[string]struct{}
	onMessage func(*marketdatav3.FeedResponse)
}

// NewMarketDataStreamerV3 creates a streamer that subscribes instrumentKeys in mode on every (re)connect.
func NewMarketDataStreamerV3(cfg *upstox_client.Configuration, instrumentKeys []string, mode Mode) (*MarketDataStreamerV3, error) {
	if !mode.valid() {
		return nil, fmt.Errorf("feeder: invalid mode provided %q", mode)
	}
	base, err := newStreamer(cfg, marketDataFeedV3URL)
	if err != nil {
		return nil, err
	}
	m := &MarketDataStreamerV3{Streamer: base, subs: map[Mode]map[string]struct{}{}}
	for _, md := range allModes {
		m.subs[md] = map[string]struct{}{}
	}
	for _, k := range instrumentKeys {
		m.subs[mode][k] = struct{}{}
	}
	base.onConnected = m.resubscribe
	base.onData = m.handleData
	return m, nil
}

// OnMessage registers the callback that receives each decoded feed message.
func (m *MarketDataStreamerV3) OnMessage(fn func(*marketdatav3.FeedResponse)) {
	m.mu.Lock()
	m.onMessage = fn
	m.mu.Unlock()
}

// Subscribe adds instrumentKeys in mode; returns ErrNotConnected before Connect.
func (m *MarketDataStreamerV3) Subscribe(instrumentKeys []string, mode Mode) error {
	if !mode.valid() {
		return fmt.Errorf("feeder: invalid mode provided %q", mode)
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	if err := m.send(methodSubscribe, instrumentKeys, mode); err != nil {
		return err
	}
	for _, k := range instrumentKeys {
		m.subs[mode][k] = struct{}{}
	}
	return nil
}

// Unsubscribe removes instrumentKeys from every mode.
func (m *MarketDataStreamerV3) Unsubscribe(instrumentKeys []string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if err := m.send(methodUnsubscribe, instrumentKeys, ""); err != nil {
		return err
	}
	m.removeKeys(instrumentKeys)
	return nil
}

// ChangeMode moves already subscribed instrumentKeys to mode.
func (m *MarketDataStreamerV3) ChangeMode(instrumentKeys []string, mode Mode) error {
	if !mode.valid() {
		return fmt.Errorf("feeder: invalid mode provided %q", mode)
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	if err := m.send(methodChangeMode, instrumentKeys, mode); err != nil {
		return err
	}
	m.removeKeys(instrumentKeys)
	for _, k := range instrumentKeys {
		m.subs[mode][k] = struct{}{}
	}
	return nil
}

// ClearSubscriptions forgets all tracked subscriptions so none are replayed on reconnect; it sends nothing.
func (m *MarketDataStreamerV3) ClearSubscriptions() {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, keys := range m.subs {
		clear(keys)
	}
}

func (m *MarketDataStreamerV3) removeKeys(instrumentKeys []string) {
	for _, keys := range m.subs {
		for _, k := range instrumentKeys {
			delete(keys, k)
		}
	}
}

// resubscribe replays every tracked subscription; called with mu held.
func (m *MarketDataStreamerV3) resubscribe(_ *websocket.Conn) error {
	for _, mode := range allModes {
		keys := make([]string, 0, len(m.subs[mode]))
		for k := range m.subs[mode] {
			keys = append(keys, k)
		}
		if len(keys) == 0 {
			continue
		}
		sort.Strings(keys)
		if err := m.send(methodSubscribe, keys, mode); err != nil {
			return err
		}
	}
	return nil
}

// send writes a binary JSON request; called with mu held.
func (m *MarketDataStreamerV3) send(method string, instrumentKeys []string, mode Mode) error {
	payload, err := buildRequest(method, instrumentKeys, mode)
	if err != nil {
		return err
	}
	return m.write(websocket.BinaryMessage, payload)
}

func (m *MarketDataStreamerV3) handleData(_ int, data []byte) {
	var resp marketdatav3.FeedResponse
	err := proto.Unmarshal(data, &resp)
	m.mu.Lock()
	onMessage, onError := m.onMessage, m.onError
	m.mu.Unlock()
	if err != nil {
		if onError != nil {
			onError(fmt.Errorf("feeder: decode feed message: %w", err))
		}
		return
	}
	if onMessage != nil {
		onMessage(&resp)
	}
}

type requestData struct {
	Mode           Mode     `json:"mode,omitempty"`
	InstrumentKeys []string `json:"instrumentKeys"`
}

type request struct {
	GUID   string      `json:"guid"`
	Method string      `json:"method"`
	Data   requestData `json:"data"`
}

func buildRequest(method string, instrumentKeys []string, mode Mode) ([]byte, error) {
	if instrumentKeys == nil {
		instrumentKeys = []string{}
	}
	return json.Marshal(request{
		GUID:   newGUID(),
		Method: method,
		Data:   requestData{Mode: mode, InstrumentKeys: instrumentKeys},
	})
}

// newGUID returns a random RFC 4122 version 4 UUID.
func newGUID() string {
	var b [16]byte
	_, _ = rand.Read(b[:])
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}
