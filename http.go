package core

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
)

// 默认值与go/pkg/http相同
const (
	defaultDialerTimeout       = 20 * time.Second
	defaultDialerKeepAlive     = 20 * time.Second
	defaultIdleConnTimeout     = 20 * time.Second
	defaultTLSHandshakeTimeout = 10 * time.Second
	defaultMaxIdleConnsPerHost = 64
	defaultMaxConnsPerHost     = 2048
	defaultWriteBufferSize     = 512 * 1024
	defaultReadBufferSize      = 512 * 1024
)

// HttpConfig http配置
type HttpConfig struct {
	Debug io.Writer
	/***********************************************
	 * Simple Setting
	 ***********************************************/
	// DialerTimeout 连接超时(默认3分钟)
	DialerTimeout time.Duration
	// DialerKeepAlive 保持活跃(默认30秒)
	DialerKeepAlive time.Duration
	// InsecureSkipVerify TLS是否跳过校验(默认false)
	InsecureSkipVerify bool
	/***********************************************
	 * Setting from http.Transport
	 ***********************************************/
	// Proxy specifies a function to return a proxy for a given
	// Request. If the function returns a non-nil error, the
	// request is aborted with the provided error.
	//
	// The proxy type is determined by the URL scheme. "http",
	// "https", "socks5", and "socks5h" are supported. If the scheme is empty,
	// "http" is assumed.
	// "socks5" is treated the same as "socks5h".
	//
	// If the proxy URL contains a userinfo subcomponent,
	// the proxy request will pass the username and password
	// in a Proxy-Authorization header.
	//
	// If Proxy is nil or returns a nil *URL, no proxy is used.
	Proxy func(*http.Request) (*url.URL, error)

	// OnProxyConnectResponse is called when the Transport gets an HTTP response from
	// a proxy for a CONNECT request. It's called before the check for a 200 OK response.
	// If it returns an error, the request fails with that error.
	OnProxyConnectResponse func(ctx context.Context, proxyURL *url.URL, connectReq *http.Request, connectRes *http.Response) error

	// DialContext specifies the dial function for creating unencrypted TCP connections.
	// If DialContext is nil (and the deprecated Dial below is also nil),
	// then the transport dials using package net.
	//
	// DialContext runs concurrently with calls to RoundTrip.
	// A RoundTrip call that initiates a dial may end up using
	// a connection dialed previously when the earlier connection
	// becomes idle before the later DialContext completes.
	DialContext func(ctx context.Context, network, addr string) (net.Conn, error)

	// DialTLSContext specifies an optional dial function for creating
	// TLS connections for non-proxied HTTPS requests.
	//
	// If DialTLSContext is nil (and the deprecated DialTLS below is also nil),
	// DialContext and TLSClientConfig are used.
	//
	// If DialTLSContext is set, the Dial and DialContext hooks are not used for HTTPS
	// requests and the TLSClientConfig and TLSHandshakeTimeout
	// are ignored. The returned net.Conn is assumed to already be
	// past the TLS handshake.
	DialTLSContext func(ctx context.Context, network, addr string) (net.Conn, error)

	// TLSClientConfig specifies the TLS configuration to use with
	// tls.Client.
	// If nil, the default configuration is used.
	// If non-nil, HTTP/2 support may not be enabled by default.
	TLSClientConfig *tls.Config

	// TLSHandshakeTimeout specifies the maximum amount of time to
	// wait for a TLS handshake. Zero means no timeout.
	TLSHandshakeTimeout time.Duration

	// DisableKeepAlives, if true, disables HTTP keep-alives and
	// will only use the connection to the server for a single
	// HTTP request.
	//
	// This is unrelated to the similarly named TCP keep-alives.
	DisableKeepAlives bool

	// DisableCompression, if true, prevents the Transport from
	// requesting compression with an "Accept-Encoding: gzip"
	// request header when the Request contains no existing
	// Accept-Encoding value. If the Transport requests gzip on
	// its own and gets a gzipped response, it's transparently
	// decoded in the Response.Body. However, if the user
	// explicitly requested gzip it is not automatically
	// uncompressed.
	DisableCompression bool

	// MaxIdleConns controls the maximum number of idle (keep-alive)
	// connections across all hosts. Zero means no limit.
	MaxIdleConns int

	// MaxIdleConnsPerHost, if non-zero, controls the maximum idle
	// (keep-alive) connections to keep per-host. If zero,
	// DefaultMaxIdleConnsPerHost is used.
	MaxIdleConnsPerHost int

	// MaxConnsPerHost optionally limits the total number of
	// connections per host, including connections in the dialing,
	// active, and idle states. On limit violation, dials will block.
	//
	// Zero means no limit.
	MaxConnsPerHost int

	// IdleConnTimeout is the maximum amount of time an idle
	// (keep-alive) connection will remain idle before closing
	// itself.
	// Zero means no limit.
	IdleConnTimeout time.Duration

	// ResponseHeaderTimeout, if non-zero, specifies the amount of
	// time to wait for a server's response headers after fully
	// writing the request (including its body, if any). This
	// time does not include the time to read the response body.
	ResponseHeaderTimeout time.Duration

	// ExpectContinueTimeout, if non-zero, specifies the amount of
	// time to wait for a server's first response headers after fully
	// writing the request headers if the request has an
	// "Expect: 100-continue" header. Zero means no timeout and
	// causes the body to be sent immediately, without
	// waiting for the server to approve.
	// This time does not include the time to send the request header.
	ExpectContinueTimeout time.Duration

	// TLSNextProto specifies how the Transport switches to an
	// alternate protocol (such as HTTP/2) after a TLS ALPN
	// protocol negotiation. If Transport dials a TLS connection
	// with a non-empty protocol name and TLSNextProto contains a
	// map entry for that key (such as "h2"), then the func is
	// called with the request's authority (such as "example.com"
	// or "example.com:1234") and the TLS connection. The function
	// must return a RoundTripper that then handles the request.
	// If TLSNextProto is not nil, HTTP/2 support is not enabled
	// automatically.
	TLSNextProto map[string]func(authority string, c *tls.Conn) http.RoundTripper

	// ProxyConnectHeader optionally specifies headers to send to
	// proxies during CONNECT requests.
	// To set the header dynamically, see GetProxyConnectHeader.
	ProxyConnectHeader http.Header

	// GetProxyConnectHeader optionally specifies a func to return
	// headers to send to proxyURL during a CONNECT request to the
	// ip:port target.
	// If it returns an error, the Transport's RoundTrip fails with
	// that error. It can return (nil, nil) to not add headers.
	// If GetProxyConnectHeader is non-nil, ProxyConnectHeader is
	// ignored.
	GetProxyConnectHeader func(ctx context.Context, proxyURL *url.URL, target string) (http.Header, error)

	// MaxResponseHeaderBytes specifies a limit on how many
	// response bytes are allowed in the server's response
	// header.
	//
	// Zero means to use a default limit.
	MaxResponseHeaderBytes int64

	// WriteBufferSize specifies the size of the write buffer used
	// when writing to the transport.
	// If zero, a default (currently 4KB) is used.
	WriteBufferSize int

	// ReadBufferSize specifies the size of the read buffer used
	// when reading from the transport.
	// If zero, a default (currently 4KB) is used.
	ReadBufferSize int

	// ForceAttemptHTTP2 controls whether HTTP/2 is enabled when a non-zero
	// Dial, DialTLS, or DialContext func or TLSClientConfig is provided.
	// By default, use of any those fields conservatively disables HTTP/2.
	// To use a custom dialer or TLS config and still attempt HTTP/2
	// upgrades, set this to true.
	ForceAttemptHTTP2 bool
}

// HttpClient 客户端接口
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// HttpAspect 切面接口
type HttpAspect interface {
	Proxy(client HttpClient) HttpClient
}

// NewHttpClient 创建客户端
func NewHttpClient(c *HttpConfig, ps ...HttpAspect) (result HttpClient) {

	transport := new(http.Transport)
	transport.Proxy = c.Proxy
	if transport.Proxy == nil {
		transport.Proxy = http.ProxyFromEnvironment
	}
	transport.OnProxyConnectResponse = c.OnProxyConnectResponse
	transport.DialContext = c.DialContext
	if transport.DialContext == nil && (c.DialerTimeout != 0 || c.DialerKeepAlive != 0) {
		transport.DialContext = (&net.Dialer{
			Timeout:   NvlD(c.DialerTimeout, defaultDialerTimeout),
			KeepAlive: NvlD(c.DialerKeepAlive, defaultDialerKeepAlive),
		}).DialContext
	}
	transport.DialTLSContext = c.DialTLSContext
	transport.TLSClientConfig = c.TLSClientConfig
	transport.TLSHandshakeTimeout = NvlD(c.TLSHandshakeTimeout, defaultTLSHandshakeTimeout)
	transport.DisableKeepAlives = c.DisableKeepAlives
	transport.DisableCompression = c.DisableCompression
	transport.MaxIdleConns = c.MaxIdleConns
	transport.MaxIdleConnsPerHost = NvlI(c.MaxIdleConnsPerHost, defaultMaxIdleConnsPerHost)
	transport.MaxConnsPerHost = NvlI(c.MaxConnsPerHost, defaultMaxConnsPerHost)
	transport.IdleConnTimeout = NvlD(c.IdleConnTimeout, defaultIdleConnTimeout)
	transport.ResponseHeaderTimeout = c.ResponseHeaderTimeout
	transport.ExpectContinueTimeout = c.ExpectContinueTimeout
	transport.TLSNextProto = c.TLSNextProto
	transport.ProxyConnectHeader = c.ProxyConnectHeader
	transport.GetProxyConnectHeader = c.GetProxyConnectHeader
	transport.MaxResponseHeaderBytes = c.MaxResponseHeaderBytes
	transport.WriteBufferSize = NvlI(c.WriteBufferSize, defaultWriteBufferSize)
	transport.ReadBufferSize = NvlI(c.ReadBufferSize, defaultReadBufferSize)
	transport.ForceAttemptHTTP2 = c.ForceAttemptHTTP2

	result = &http.Client{
		Transport: transport,
	}
	for _, p := range ps {
		result = p.Proxy(result)
	}
	if c.Debug != nil {
		result = &DebugHttpClient{
			HttpClient: result,
			Debug:      c.Debug,
		}
	}
	return result
}

type DebugHttpClient struct {
	HttpClient
	Debug io.Writer
}

func (dc DebugHttpClient) Do(req *http.Request) (*http.Response, error) {
	data, _ := io.ReadAll(req.Body)
	req.Body = &BuffBody{
		Data: data,
	}
	res, err := dc.HttpClient.Do(req)
	if err != nil {
		return res, err
	}
	buf := new(bytes.Buffer)
	fmt.Fprintln(buf, "------Http Request------")
	fmt.Fprintln(buf, req.Method, req.URL.String())
	for k, vs := range req.Header {
		fmt.Fprintln(buf, k, ":", vs)
	}
	buf.Write(data)
	fmt.Fprintln(buf)
	fmt.Fprintln(buf, "------Http Response------")
	fmt.Fprintln(buf, res.Proto, res.Status)
	for k, vs := range res.Header {
		fmt.Fprintln(buf, k, ":", vs)
	}
	data, _ = io.ReadAll(res.Body)
	res.Body = &BuffBody{
		Data: data,
	}
	buf.Write(data)
	buf.WriteString("\n")

	dc.Debug.Write(buf.Bytes())

	return res, nil
}

// BuffBody 一次性读写Body, 配合ReadBody()一块使用
type BuffBody struct {
	mark int
	Data []byte
}

func (b *BuffBody) Read(p []byte) (int, error) {
	if b.mark >= len(b.Data) {
		return 0, io.EOF
	}
	n := copy(p, b.Data[b.mark:])
	b.mark += n
	return n, nil
}

func (b *BuffBody) Close() error {
	b.mark = len(b.Data)
	return nil
}

func (b *BuffBody) Reset() {
	b.mark = 0
}

var _ io.ReadCloser = (*BuffBody)(nil)

var DiscardBuffer = make([]byte, 32*1024)

func DiscardResponse(res *http.Response) error {
	for {
		_, err := res.Body.Read(DiscardBuffer)
		if err != nil {
			break
		}
	}
	return res.Body.Close()
}
