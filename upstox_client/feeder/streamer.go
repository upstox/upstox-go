// Package feeder streams live market data and portfolio updates from the Upstox WebSocket feeds.
package feeder

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/upstox/upstox-go/upstox_client"
)

const (
	defaultReconnectInterval = time.Second
	defaultRetryCount        = 5
	maxRedirects             = 3
	writeTimeout             = 10 * time.Second
	handshakeTimeout         = 30 * time.Second
)

// ErrNotConnected is returned when an operation needs an open WebSocket connection.
var ErrNotConnected = errors.New("feeder: websocket is not open")

// HandshakeError reports a WebSocket upgrade rejected by the server.
type HandshakeError struct {
	StatusCode int
	Body       string
}

func (e *HandshakeError) Error() string {
	return fmt.Sprintf("feeder: websocket handshake failed: %d %s: %s", e.StatusCode, http.StatusText(e.StatusCode), e.Body)
}

// Streamer is the connection and auto-reconnect core shared by all feed streamers.
// All methods are safe for concurrent use; callbacks run on the streamer's read goroutine.
type Streamer struct {
	cfg    *upstox_client.Configuration
	url    string
	dialer *websocket.Dialer

	// onConnected runs with mu held right after each successful dial, before OnOpen.
	onConnected func(conn *websocket.Conn) error
	onData      func(messageType int, data []byte)

	mu             sync.Mutex
	conn           *websocket.Conn
	done           chan struct{}
	userClosed     bool
	autoReconnect  bool
	interval       time.Duration
	retryCount     int
	onOpen         func()
	onError        func(error)
	onClose        func(code int, reason string)
	onReconnecting func(attempt int)
	onStopped      func(reason string)
}

func newStreamer(cfg *upstox_client.Configuration, feedURL string) (*Streamer, error) {
	if cfg == nil {
		return nil, errors.New("feeder: configuration is nil")
	}
	return &Streamer{
		cfg: cfg,
		url: feedURL,
		dialer: &websocket.Dialer{
			Proxy:            http.ProxyFromEnvironment,
			HandshakeTimeout: handshakeTimeout,
		},
		autoReconnect: true,
		interval:      defaultReconnectInterval,
		retryCount:    defaultRetryCount,
	}, nil
}

// OnOpen registers a callback invoked after every successful (re)connection.
func (s *Streamer) OnOpen(fn func()) { s.mu.Lock(); s.onOpen = fn; s.mu.Unlock() }

// OnError registers a callback for connection, decode and reconnect errors.
func (s *Streamer) OnError(fn func(error)) { s.mu.Lock(); s.onError = fn; s.mu.Unlock() }

// OnClose registers a callback invoked when the connection closes.
func (s *Streamer) OnClose(fn func(code int, reason string)) {
	s.mu.Lock()
	s.onClose = fn
	s.mu.Unlock()
}

// OnReconnecting registers a callback invoked before each reconnect attempt (1-based).
func (s *Streamer) OnReconnecting(fn func(attempt int)) {
	s.mu.Lock()
	s.onReconnecting = fn
	s.mu.Unlock()
}

// OnAutoReconnectStopped registers a callback invoked when auto-reconnect gives up or is disabled.
func (s *Streamer) OnAutoReconnectStopped(fn func(reason string)) {
	s.mu.Lock()
	s.onStopped = fn
	s.mu.Unlock()
}

// AutoReconnect configures reconnection after an abnormal close; defaults are enabled, 1s, 5 retries.
func (s *Streamer) AutoReconnect(enable bool, interval time.Duration, retryCount int) {
	s.mu.Lock()
	s.autoReconnect = enable
	s.interval = interval
	s.retryCount = retryCount
	stopped := s.onStopped
	s.mu.Unlock()
	if !enable && stopped != nil {
		stopped("Disabled by client.")
	}
}

// Connect dials the feed and starts reading; it is a no-op when already connected.
func (s *Streamer) Connect() error {
	s.mu.Lock()
	if s.conn != nil {
		s.mu.Unlock()
		return nil
	}
	s.userClosed = false
	s.done = make(chan struct{})
	done := s.done
	s.mu.Unlock()

	conn, err := s.dial()
	if err != nil {
		return err
	}
	return s.attach(conn, done)
}

// Disconnect closes the connection with status 1000 and disables reconnection until the next Connect.
func (s *Streamer) Disconnect() {
	s.mu.Lock()
	if !s.userClosed && s.done != nil {
		close(s.done)
	}
	s.userClosed = true
	conn := s.conn
	s.conn = nil
	if conn != nil {
		_ = conn.SetWriteDeadline(time.Now().Add(writeTimeout))
		_ = conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	}
	s.mu.Unlock()
	if conn != nil {
		_ = conn.Close()
	}
}

// write sends one message on the open connection; the caller must hold mu.
func (s *Streamer) write(messageType int, data []byte) error {
	if s.conn == nil {
		return ErrNotConnected
	}
	_ = s.conn.SetWriteDeadline(time.Now().Add(writeTimeout))
	return s.conn.WriteMessage(messageType, data)
}

func (s *Streamer) headers() http.Header {
	h := http.Header{}
	for k, v := range s.cfg.DefaultHeader {
		h.Set(k, v)
	}
	if s.cfg.UserAgent != "" {
		h.Set("User-Agent", s.cfg.UserAgent)
	}
	h.Set("Authorization", "Bearer "+s.cfg.AccessToken)
	return h
}

// dial opens the feed, following up to maxRedirects authorize redirects like the Python client.
func (s *Streamer) dial() (*websocket.Conn, error) {
	if s.cfg.AccessToken == "" {
		return nil, errors.New("feeder: configuration AccessToken is empty")
	}
	target := s.url
	for i := 0; ; i++ {
		conn, resp, err := s.dialer.Dial(target, s.headers())
		if err == nil {
			return conn, nil
		}
		if resp == nil {
			return nil, err
		}
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
		resp.Body.Close()
		loc := resp.Header.Get("Location")
		if isRedirect(resp.StatusCode) && loc != "" && i < maxRedirects {
			next, perr := resolveWS(target, loc)
			if perr != nil {
				return nil, perr
			}
			target = next
			continue
		}
		return nil, &HandshakeError{StatusCode: resp.StatusCode, Body: string(body)}
	}
}

func isRedirect(code int) bool {
	switch code {
	case http.StatusMovedPermanently, http.StatusFound, http.StatusSeeOther,
		http.StatusTemporaryRedirect, http.StatusPermanentRedirect:
		return true
	}
	return false
}

// resolveWS resolves a redirect Location against base and maps http(s) to ws(s).
func resolveWS(base, loc string) (string, error) {
	b, err := url.Parse(base)
	if err != nil {
		return "", err
	}
	u, err := b.Parse(loc)
	if err != nil {
		return "", err
	}
	switch u.Scheme {
	case "http":
		u.Scheme = "ws"
	case "https":
		u.Scheme = "wss"
	}
	return u.String(), nil
}

// attach installs a freshly dialed connection, replays state, fires OnOpen and starts the read loop.
func (s *Streamer) attach(conn *websocket.Conn, done chan struct{}) error {
	s.mu.Lock()
	if s.userClosed || s.done != done || s.conn != nil {
		// Lost a race with Disconnect or a concurrent Connect.
		alreadyOpen := s.conn != nil
		s.mu.Unlock()
		_ = conn.Close()
		if alreadyOpen {
			return nil
		}
		return ErrNotConnected
	}
	s.conn = conn
	var hookErr error
	if s.onConnected != nil {
		hookErr = s.onConnected(conn)
	}
	onOpen, onError := s.onOpen, s.onError
	s.mu.Unlock()

	if hookErr != nil && onError != nil {
		onError(hookErr)
	}
	if onOpen != nil {
		onOpen()
	}
	go s.readLoop(conn, done)
	return nil
}

func (s *Streamer) readLoop(conn *websocket.Conn, done chan struct{}) {
	for {
		mt, data, err := conn.ReadMessage()
		if err != nil {
			s.handleDrop(conn, done, err)
			return
		}
		if s.onData != nil {
			s.onData(mt, data)
		}
	}
}

func (s *Streamer) handleDrop(conn *websocket.Conn, done chan struct{}, err error) {
	code, reason := websocket.CloseAbnormalClosure, err.Error()
	var ce *websocket.CloseError
	if errors.As(err, &ce) {
		code, reason = ce.Code, ce.Text
	}

	s.mu.Lock()
	userClosed := s.userClosed || s.done != done
	if s.conn == conn {
		s.conn = nil
	}
	onClose, onError := s.onClose, s.onError
	reconnect := s.autoReconnect && !userClosed && code != websocket.CloseNormalClosure
	s.mu.Unlock()
	_ = conn.Close()

	if userClosed {
		code, reason = websocket.CloseNormalClosure, ""
	} else if ce == nil && onError != nil {
		onError(err)
	}
	if onClose != nil {
		onClose(code, reason)
	}
	if reconnect {
		s.reconnectLoop(done)
	}
}

func (s *Streamer) reconnectLoop(done chan struct{}) {
	s.mu.Lock()
	interval, retries := s.interval, s.retryCount
	s.mu.Unlock()

	for attempt := 1; attempt <= retries; attempt++ {
		select {
		case <-done:
			return
		case <-time.After(interval):
		}
		s.mu.Lock()
		onReconnecting, onError, onStopped, enabled := s.onReconnecting, s.onError, s.onStopped, s.autoReconnect
		s.mu.Unlock()
		if !enabled {
			return
		}
		if onReconnecting != nil {
			onReconnecting(attempt)
		}
		conn, err := s.dial()
		if err == nil {
			_ = s.attach(conn, done)
			return
		}
		if onError != nil {
			onError(err)
		}
		var he *HandshakeError
		if errors.As(err, &he) && he.StatusCode == http.StatusUnauthorized {
			// A rejected token will not recover by retrying.
			if onStopped != nil {
				onStopped("401 Unauthorized: access token rejected.")
			}
			return
		}
	}
	s.mu.Lock()
	onStopped := s.onStopped
	s.mu.Unlock()
	if onStopped != nil {
		onStopped(fmt.Sprintf("retryCount of %d exhausted.", retries))
	}
}
