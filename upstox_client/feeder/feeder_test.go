package feeder

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/upstox/upstox-go/upstox_client"
	"github.com/upstox/upstox-go/upstox_client/feeder/marketdatav3"
	"google.golang.org/protobuf/proto"
)

const testTimeout = 5 * time.Second

var upgrader = websocket.Upgrader{}

type serverConn struct {
	conn *websocket.Conn
	msgs chan []byte
}

// testServer upgrades every request on /ws and hands each connection to the test.
type testServer struct {
	*httptest.Server
	conns     chan *serverConn
	reject401 atomic.Bool
	mu        sync.Mutex
	lastReq   *http.Request
}

func newTestServer(t *testing.T) *testServer {
	t.Helper()
	ts := &testServer{conns: make(chan *serverConn, 8)}
	mux := http.NewServeMux()
	mux.HandleFunc("/feed", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer test-token" {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}
		http.Redirect(w, r, "/ws?code=abc&"+r.URL.RawQuery, http.StatusFound)
	})
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ts.mu.Lock()
		ts.lastReq = r
		ts.mu.Unlock()
		if ts.reject401.Load() {
			http.Error(w, `{"status":"error"}`, http.StatusUnauthorized)
			return
		}
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		sc := &serverConn{conn: c, msgs: make(chan []byte, 16)}
		go func() {
			defer close(sc.msgs)
			for {
				_, data, err := c.ReadMessage()
				if err != nil {
					return
				}
				sc.msgs <- data
			}
		}()
		ts.conns <- sc
	})
	ts.Server = httptest.NewServer(mux)
	t.Cleanup(ts.Close)
	return ts
}

func (ts *testServer) feedURL() string {
	return "ws" + strings.TrimPrefix(ts.URL, "http") + "/feed"
}

func (ts *testServer) nextConn(t *testing.T) *serverConn {
	t.Helper()
	select {
	case c := <-ts.conns:
		return c
	case <-time.After(testTimeout):
		t.Fatal("timed out waiting for client connection")
		return nil
	}
}

func (sc *serverConn) nextRequest(t *testing.T) request {
	t.Helper()
	select {
	case data := <-sc.msgs:
		var req request
		if err := json.Unmarshal(data, &req); err != nil {
			t.Fatalf("decode request %q: %v", data, err)
		}
		return req
	case <-time.After(testTimeout):
		t.Fatal("timed out waiting for client request")
		return request{}
	}
}

func testConfig() *upstox_client.Configuration {
	cfg := upstox_client.NewConfiguration()
	cfg.AccessToken = "test-token"
	return cfg
}

func newTestMarketStreamer(t *testing.T, ts *testServer, keys []string, mode Mode) *MarketDataStreamerV3 {
	t.Helper()
	m, err := NewMarketDataStreamerV3(testConfig(), keys, mode)
	if err != nil {
		t.Fatal(err)
	}
	m.url = ts.feedURL()
	t.Cleanup(m.Disconnect)
	return m
}

func TestMarketDataConnectAndDecode(t *testing.T) {
	ts := newTestServer(t)
	m := newTestMarketStreamer(t, ts, []string{"NSE_INDEX|Nifty 50"}, ModeFull)

	opened := make(chan struct{}, 1)
	got := make(chan *marketdatav3.FeedResponse, 1)
	m.OnOpen(func() { opened <- struct{}{} })
	m.OnMessage(func(r *marketdatav3.FeedResponse) { got <- r })

	if err := m.Connect(); err != nil {
		t.Fatal(err)
	}
	sc := ts.nextConn(t)
	<-opened

	req := sc.nextRequest(t)
	if req.Method != "sub" || req.Data.Mode != ModeFull || len(req.Data.InstrumentKeys) != 1 {
		t.Fatalf("unexpected initial subscribe: %+v", req)
	}

	msg := &marketdatav3.FeedResponse{
		Type:      marketdatav3.Type_live_feed,
		CurrentTs: 1700000000000,
		Feeds: map[string]*marketdatav3.Feed{
			"NSE_INDEX|Nifty 50": {FeedUnion: &marketdatav3.Feed_Ltpc{Ltpc: &marketdatav3.LTPC{Ltp: 22000.5}}},
		},
	}
	data, err := proto.Marshal(msg)
	if err != nil {
		t.Fatal(err)
	}
	if err := sc.conn.WriteMessage(websocket.BinaryMessage, data); err != nil {
		t.Fatal(err)
	}
	select {
	case r := <-got:
		if r.GetFeeds()["NSE_INDEX|Nifty 50"].GetLtpc().GetLtp() != 22000.5 || r.GetType() != marketdatav3.Type_live_feed {
			t.Fatalf("unexpected decoded message: %v", r)
		}
	case <-time.After(testTimeout):
		t.Fatal("no message delivered")
	}

	ts.mu.Lock()
	r := ts.lastReq
	ts.mu.Unlock()
	if r.URL.Query().Get("code") != "abc" {
		t.Fatalf("redirect not followed: %s", r.URL)
	}
}

func TestSubscribePayloadShape(t *testing.T) {
	ts := newTestServer(t)
	m := newTestMarketStreamer(t, ts, nil, ModeLTPC)

	if err := m.Subscribe([]string{"NSE_EQ|INE002A01018"}, ModeFull); !errors.Is(err, ErrNotConnected) {
		t.Fatalf("Subscribe before Connect = %v, want ErrNotConnected", err)
	}
	if err := m.Connect(); err != nil {
		t.Fatal(err)
	}
	sc := ts.nextConn(t)

	if err := m.Subscribe([]string{"NSE_EQ|INE002A01018"}, ModeFull); err != nil {
		t.Fatal(err)
	}
	data := <-sc.msgs
	var raw map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatal(err)
	}
	if raw["method"] != "sub" || raw["guid"] == "" {
		t.Fatalf("unexpected payload %s", data)
	}
	d := raw["data"].(map[string]any)
	if d["mode"] != "full" || d["instrumentKeys"].([]any)[0] != "NSE_EQ|INE002A01018" {
		t.Fatalf("unexpected data %s", data)
	}

	if err := m.ChangeMode([]string{"NSE_EQ|INE002A01018"}, ModeLTPC); err != nil {
		t.Fatal(err)
	}
	if req := sc.nextRequest(t); req.Method != "change_mode" || req.Data.Mode != ModeLTPC {
		t.Fatalf("unexpected change_mode %+v", req)
	}
	if err := m.Unsubscribe([]string{"NSE_EQ|INE002A01018"}); err != nil {
		t.Fatal(err)
	}
	data = <-sc.msgs
	if !strings.Contains(string(data), `"method":"unsub"`) || strings.Contains(string(data), `"mode"`) {
		t.Fatalf("unexpected unsub payload %s", data)
	}
	if err := m.Subscribe([]string{"X"}, Mode("bogus")); err == nil {
		t.Fatal("invalid mode accepted")
	}
}

func TestResubscribeAfterReconnect(t *testing.T) {
	ts := newTestServer(t)
	m := newTestMarketStreamer(t, ts, []string{"NSE_INDEX|Nifty 50"}, ModeLTPC)
	m.AutoReconnect(true, 10*time.Millisecond, 3)

	reconnecting := make(chan int, 4)
	m.OnReconnecting(func(n int) { reconnecting <- n })
	if err := m.Connect(); err != nil {
		t.Fatal(err)
	}
	sc := ts.nextConn(t)
	sc.nextRequest(t)
	if err := m.Subscribe([]string{"NSE_FO|45450"}, ModeFull); err != nil {
		t.Fatal(err)
	}
	sc.nextRequest(t)

	// Drop the TCP connection without a close frame to simulate a network failure.
	_ = sc.conn.NetConn().Close()

	sc2 := ts.nextConn(t)
	if n := <-reconnecting; n != 1 {
		t.Fatalf("first reconnect attempt = %d", n)
	}
	seen := map[Mode][]string{}
	for i := 0; i < 2; i++ {
		req := sc2.nextRequest(t)
		seen[req.Data.Mode] = req.Data.InstrumentKeys
	}
	if len(seen[ModeLTPC]) != 1 || seen[ModeLTPC][0] != "NSE_INDEX|Nifty 50" ||
		len(seen[ModeFull]) != 1 || seen[ModeFull][0] != "NSE_FO|45450" {
		t.Fatalf("resubscribed %v", seen)
	}
}

func TestUnauthorizedStopsReconnect(t *testing.T) {
	ts := newTestServer(t)
	m := newTestMarketStreamer(t, ts, nil, ModeLTPC)
	m.AutoReconnect(true, 10*time.Millisecond, 5)

	var attempts atomic.Int32
	stopped := make(chan string, 1)
	m.OnReconnecting(func(int) { attempts.Add(1) })
	m.OnAutoReconnectStopped(func(reason string) { stopped <- reason })
	if err := m.Connect(); err != nil {
		t.Fatal(err)
	}
	sc := ts.nextConn(t)
	ts.reject401.Store(true)
	_ = sc.conn.NetConn().Close()

	select {
	case reason := <-stopped:
		if !strings.Contains(reason, "401") {
			t.Fatalf("stop reason %q", reason)
		}
	case <-time.After(testTimeout):
		t.Fatal("auto-reconnect did not stop")
	}
	time.Sleep(50 * time.Millisecond)
	if n := attempts.Load(); n != 1 {
		t.Fatalf("reconnect attempts = %d, want 1", n)
	}

	m.Disconnect()
	err := m.Connect()
	var he *HandshakeError
	if !errors.As(err, &he) || he.StatusCode != http.StatusUnauthorized {
		t.Fatalf("Connect after 401 = %v, want HandshakeError 401", err)
	}
}

func TestPortfolioQueryAndMessages(t *testing.T) {
	cases := []struct {
		in   PortfolioUpdates
		want string
	}{
		{PortfolioUpdates{}, portfolioStreamFeedURL},
		{PortfolioUpdates{Order: true}, portfolioStreamFeedURL + "?update_types=order"},
		{PortfolioUpdates{Order: true, Position: true, Holding: true, GTTOrder: true},
			portfolioStreamFeedURL + "?update_types=order%2Cholding%2Cposition%2Cgtt_order"},
	}
	for _, c := range cases {
		p, err := NewPortfolioDataStreamer(testConfig(), c.in)
		if err != nil {
			t.Fatal(err)
		}
		if p.url != c.want {
			t.Errorf("url for %+v = %s, want %s", c.in, p.url, c.want)
		}
	}

	ts := newTestServer(t)
	p, err := NewPortfolioDataStreamer(testConfig(), PortfolioUpdates{Order: true, Position: true})
	if err != nil {
		t.Fatal(err)
	}
	p.url = ts.feedURL() + "?" + strings.SplitN(p.url, "?", 2)[1]
	t.Cleanup(p.Disconnect)

	got := make(chan []byte, 1)
	closed := make(chan int, 1)
	p.OnMessage(func(b []byte) { got <- b })
	p.OnClose(func(code int, _ string) { closed <- code })
	if err := p.Connect(); err != nil {
		t.Fatal(err)
	}
	sc := ts.nextConn(t)
	ts.mu.Lock()
	q := ts.lastReq.URL.Query().Get("update_types")
	ts.mu.Unlock()
	if q != "order,position" {
		t.Fatalf("update_types = %q", q)
	}

	payload := `{"update_type":"order","status":"complete"}`
	if err := sc.conn.WriteMessage(websocket.TextMessage, []byte(payload)); err != nil {
		t.Fatal(err)
	}
	select {
	case b := <-got:
		if string(b) != payload {
			t.Fatalf("message = %s", b)
		}
	case <-time.After(testTimeout):
		t.Fatal("no message delivered")
	}

	p.Disconnect()
	select {
	case code := <-closed:
		if code != websocket.CloseNormalClosure {
			t.Fatalf("close code = %d", code)
		}
	case <-time.After(testTimeout):
		t.Fatal("OnClose not called")
	}
}

func TestNewMarketDataStreamerV3Validation(t *testing.T) {
	if _, err := NewMarketDataStreamerV3(testConfig(), nil, Mode("bad")); err == nil {
		t.Fatal("invalid mode accepted")
	}
	if _, err := NewMarketDataStreamerV3(nil, nil, ModeFull); err == nil {
		t.Fatal("nil configuration accepted")
	}
}
