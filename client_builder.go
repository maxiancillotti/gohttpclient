package gohttpclient

import (
	"net"
	"net/http"
	"time"
)

// Methods returning interface can concatenate method calls
type ClientBuilder interface {

	// SetHeaders: set common headers to use during all client life
	// If Content-Type or Accept aren't set, default is application/json.
	SetHeaders(headers http.Header) ClientBuilder

	// SetConnectionTimeout sets the request connection timeout.
	// Default is 10 seconds.
	SetConnectionTimeout(connectionTimeout time.Duration) ClientBuilder

	// SetResponseTimeout sets the response timeout after we have sent the Request.
	// Default is 30 seconds.
	// ResponseHeaderTimeout, if non-zero, specifies the amount of
	// time to wait for a server's response headers after fully
	// writing the request (including its body, if any). This
	// time does not include the time to read the response body.
	SetResponseTimeout(requestTimeout time.Duration) ClientBuilder

	// ExpectContinueTimeout limits the time the client will wait between sending the request headers
	// when including an "Expect: 100-continue" and receiving the go-ahead to send the body
	// ExpectContinueTimeout, if non-zero, specifies the amount of
	// time to wait for a server's first response headers after fully
	// writing the request headers if the request has an
	// "Expect: 100-continue" header. Zero means no timeout and
	// causes the body to be sent immediately, without
	// waiting for the server to approve.
	// This time does not include the time to send the request header
	// Default is 1 second.
	SetExpectContinueTimeout(timeout time.Duration) ClientBuilder

	// TLSHandshakeTimeout specifies the maximum amount of time waiting to
	// wait for a TLS handshake. Zero means no timeout. Default 10 seconds.
	SetTLSHandshakeTimeout(timeout time.Duration) ClientBuilder

	// IdleConnTimeout controls how long an idle connection is kept in the connection pool.
	// It does not control a blocking phase of a client request.
	// IdleConnTimeout is the maximum amount of time an idle
	// (keep-alive) connection will remain idle before closing
	// itself.
	// Zero means no limit.
	// Default 90 seconds.
	SetIdleConnTimeout(timeout time.Duration) ClientBuilder

	// MaxIdleConns controls the maximum number of idle (keep-alive)
	// connections across all hosts. Zero means no limit.
	// Default is 100.
	SetMaxIdleConnections(maxIdleConnections int) ClientBuilder

	// MaxIdleConnsPerHost, if non-zero, controls the maximum idle
	// (keep-alive) connections to keep per-host. If zero,
	// standard library default is used.
	// Requests per minute is a good metric to set this value.
	// Default is 100.
	SetMaxIdleConnectionsPerHost(maxIdleConnsPerHost int) ClientBuilder

	// MaxConnsPerHost optionally limits the total number of
	// connections per host, including connections in the dialing,
	// active, and idle states. On limit violation, dials will block.
	// Zero means no limit.
	// Default is 512.
	SetMaxConnectionsPerHost(maxConnsPerHost int) ClientBuilder

	// KeepAlive specifies the interval between keep-alive
	// probes for an active network connection.
	// If zero, keep-alive probes are sent with a default value
	// (currently 15 seconds), if supported by the protocol and operating
	// system. Network protocols or operating systems that do
	// not support keep-alives ignore this field.
	// If negative, keep-alive probes are disabled.
	SetDialerKeepAlive(keepTime time.Duration) ClientBuilder

	// FallbackDelay specifies the length of time to wait before
	// spawning a RFC 6555 Fast Fallback connection. That is, this
	// is the amount of time to wait for IPv6 to succeed before
	// assuming that IPv6 is misconfigured and falling back to
	// IPv4.
	//
	// If zero, a default delay of 300ms is used.
	// A negative value disables Fast Fallback support.
	SetDialerFallbackDelay(delay time.Duration) ClientBuilder

	// LocalAddr is the local address to use when dialing an
	// address. The address must be of a compatible type for the
	// network being dialed.
	// If nil, a local address is automatically chosen.
	SetDialerLocalAddr(localAddr net.Addr) ClientBuilder

	// ForceAttemptHTTP2 controls whether HTTP/2 is enabled when a non-zero
	// Dial, DialTLS, or DialContext func or TLSClientConfig is provided.
	// By default, use of any those fields conservatively disables HTTP/2.
	// To use a custom dialer or TLS config and still attempt HTTP/2
	// upgrades, set this to true.
	ForceAttemptHTTP2(enable bool) ClientBuilder

	// A CookieJar manages storage and use of cookies in HTTP requests.
	// Implementations of CookieJar must be safe for concurrent use by multiple
	// goroutines.
	// The net/http/cookiejar package provides a CookieJar implementation.
	SetCookieJar(cookieJar http.CookieJar) ClientBuilder

	// Build sets the previously configured parameters into our HTTP client
	// and returns it to perform the desired HTTP calls.
	Build() Client
}

type clientBuilder struct {
	connectionTimeout     time.Duration
	responseTimeOut       time.Duration
	expectContinueTimeout time.Duration
	tlsHandshakeTimeout   time.Duration

	idleConnTimeout time.Duration

	maxIdleConns        int
	maxIdleConnsPerHost int
	maxConnsPerHost     int

	keepAliveTime time.Duration
	fallbackDelay time.Duration
	localAddr     net.Addr

	forceAttemptHTTP2Enabled bool

	headers http.Header

	cookieJar http.CookieJar
}

// NewBuiler returns a ClientBuilder that you can configure to build
// finally your HTTP client.
func NewBuilder() ClientBuilder {
	return &clientBuilder{

		connectionTimeout:     defaultConnectionTimeout,
		responseTimeOut:       defaultResponseTimeOut,
		expectContinueTimeout: defaultExpectContinueTimeout,
		tlsHandshakeTimeout:   defaultTLSHandshakeTimeout,

		idleConnTimeout: defaultIdleConnTimeout,

		maxIdleConns: defaultMaxIdleConnections,

		keepAliveTime: defaultKeepAliveTime,
		fallbackDelay: defaultFallbackDelay,
		localAddr:     nil,

		forceAttemptHTTP2Enabled: defaultForceAttemptHTTP2Enabled,

		cookieJar: nil,
	}
}

func (b *clientBuilder) Build() Client {
	return &client{
		builder: b,
	}
}

func (b *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	b.headers = headers
	return b
}

func (b *clientBuilder) SetConnectionTimeout(connectionTimeout time.Duration) ClientBuilder {
	b.connectionTimeout = connectionTimeout
	return b
}

func (b *clientBuilder) SetResponseTimeout(responseTimeOut time.Duration) ClientBuilder {
	b.responseTimeOut = responseTimeOut
	return b
}

func (b *clientBuilder) SetExpectContinueTimeout(timeout time.Duration) ClientBuilder {
	b.expectContinueTimeout = timeout
	return b
}

func (b *clientBuilder) SetTLSHandshakeTimeout(timeout time.Duration) ClientBuilder {
	b.tlsHandshakeTimeout = timeout
	return b
}

func (b *clientBuilder) SetIdleConnTimeout(timeout time.Duration) ClientBuilder {
	b.idleConnTimeout = timeout
	return b
}

func (b *clientBuilder) SetMaxIdleConnections(maxIdleConns int) ClientBuilder {
	b.maxIdleConns = maxIdleConns
	return b
}

func (b *clientBuilder) SetMaxIdleConnectionsPerHost(maxIdleConnsPerHost int) ClientBuilder {
	b.maxIdleConnsPerHost = maxIdleConnsPerHost
	return b
}

func (b *clientBuilder) SetMaxConnectionsPerHost(maxConnsPerHost int) ClientBuilder {
	b.maxConnsPerHost = maxConnsPerHost
	return b
}

func (b *clientBuilder) SetDialerKeepAlive(keepTime time.Duration) ClientBuilder {
	b.keepAliveTime = keepTime
	return b
}

func (b *clientBuilder) SetDialerFallbackDelay(delay time.Duration) ClientBuilder {
	b.fallbackDelay = delay
	return b
}

func (b *clientBuilder) SetDialerLocalAddr(localAddr net.Addr) ClientBuilder {
	b.localAddr = localAddr
	return b
}

func (b *clientBuilder) ForceAttemptHTTP2(enable bool) ClientBuilder {
	b.forceAttemptHTTP2Enabled = enable
	return b
}

func (b *clientBuilder) SetCookieJar(cookieJar http.CookieJar) ClientBuilder {
	b.cookieJar = cookieJar
	return b
}
