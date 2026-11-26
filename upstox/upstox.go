package upstox

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// SDKVersion is the version of this SDK, sent with every request.
const SDKVersion = "1.0.0"

const (
	defaultHost      = "api.upstox.com"
	defaultOrderHost = "api-hft.upstox.com"
	sandboxHost      = "api-sandbox.upstox.com"
)

// sandboxPaths are the only endpoints the sandbox environment serves.
var sandboxPaths = map[string]bool{
	"/v2/order/place":       true,
	"/v2/order/modify":      true,
	"/v2/order/cancel":      true,
	"/v2/order/multi/place": true,
	"/v3/order/place":       true,
	"/v3/order/modify":      true,
	"/v3/order/cancel":      true,
}

// NewSandboxConfiguration returns a configuration that sends every call to the Upstox sandbox.
func NewSandboxConfiguration() *Configuration {
	cfg := NewConfiguration()
	cfg.Sandbox = true
	return cfg
}

// routeRequest sends order calls to the HFT host and enforces sandbox restrictions.
func (c *Configuration) routeRequest(u *url.URL) error {
	if c.Sandbox {
		if !sandboxPaths[u.Path] {
			return fmt.Errorf("upstox: %s is not available in sandbox mode", u.Path)
		}
		u.Host = sandboxHost
		return nil
	}
	// A custom BasePath (e.g. a proxy or test server) disables HFT routing
	if u.Host == defaultHost && c.OrderHost != "" && isOrderPath(u.Path) {
		u.Host = c.OrderHost
	}
	return nil
}

func isOrderPath(path string) bool {
	// US stock orders are served by the main host, not HFT
	if strings.HasPrefix(path, "/v2/us/") {
		return false
	}
	return strings.Contains(path, "/order/place") ||
		strings.Contains(path, "/order/modify") ||
		strings.Contains(path, "/order/cancel")
}

// The exports below are used by the generated upstox/api package.

// Config returns the client's configuration.
func (c *APIClient) Config() *Configuration { return c.cfg }

// PrepareRequest builds an HTTP request with auth, headers and routing applied.
func (c *APIClient) PrepareRequest(ctx context.Context, path string, method string, postBody interface{},
	headerParams map[string]string, queryParams url.Values, formParams url.Values,
	fileName string, fileBytes []byte) (*http.Request, error) {
	return c.prepareRequest(ctx, path, method, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
}

// CallAPI sends a prepared request.
func (c *APIClient) CallAPI(request *http.Request) (*http.Response, error) { return c.callAPI(request) }

// Decode unmarshals a response body by content type.
func (c *APIClient) Decode(v interface{}, b []byte, contentType string) error {
	return c.decode(v, b, contentType)
}

// NewGenericSwaggerError builds the error returned for non-2xx responses.
func NewGenericSwaggerError(body []byte, msg string, model interface{}) GenericSwaggerError {
	return GenericSwaggerError{body: body, error: msg, model: model}
}

// ParameterToString formats a parameter for a path, query or header value.
func ParameterToString(obj interface{}, collectionFormat string) string {
	return parameterToString(obj, collectionFormat)
}

// SelectHeaderContentType picks the request Content-Type.
func SelectHeaderContentType(contentTypes []string) string {
	return selectHeaderContentType(contentTypes)
}

// SelectHeaderAccept builds the Accept header.
func SelectHeaderAccept(accepts []string) string { return selectHeaderAccept(accepts) }

// ReportError formats a parameter validation error.
func ReportError(format string, a ...interface{}) error { return reportError(format, a...) }

// Strlen counts runes, as required-parameter checks expect.
func Strlen(s string) int { return strlen(s) }
