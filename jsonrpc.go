package base

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type JsonMarshal func(v any) ([]byte, error)
type JsonUnmarshal func(r io.Reader, v any) error

// HttpHeader http头部
type HttpHeader interface {
	Init(furl string, body []byte, header http.Header)
}

// JsonRpcClient rpc客户端
type JsonRpcClient struct {
	endpoint string
	client   HttpClient
	header   HttpHeader
	encoder  JsonMarshal
	decoder  JsonUnmarshal
}

func (c *JsonRpcClient) GET(uri string, req any, rsp any, status ...int) error {
	return c.Do(http.MethodGet, uri, req, rsp, status...)
}

func (c *JsonRpcClient) POST(uri string, req any, rsp any, status ...int) error {
	return c.Do(http.MethodPost, uri, req, rsp, status...)
}

func (c *JsonRpcClient) PUT(uri string, req any, rsp any, status ...int) error {
	return c.Do(http.MethodPut, uri, req, rsp, status...)
}

func (c *JsonRpcClient) DELETE(uri string, req any, rsp any, status ...int) error {
	return c.Do(http.MethodDelete, uri, req, rsp, status...)
}

func (c *JsonRpcClient) HEAD(uri string, req any, rsp any, status ...int) error {
	return c.Do(http.MethodHead, uri, req, rsp, status...)
}

func (c *JsonRpcClient) PATCH(uri string, req any, rsp any, status ...int) error {
	return c.Do(http.MethodPatch, uri, req, rsp, status...)
}

func (c *JsonRpcClient) OPTIONS(uri string, req any, rsp any, status ...int) error {
	return c.Do(http.MethodOptions, uri, req, rsp, status...)
}

func (c *JsonRpcClient) CONNECT(uri string, req any, rsp any, status ...int) error {
	return c.Do(http.MethodConnect, uri, req, rsp, status...)
}

// Do 远程调用. 可以指定期望的status
func (c *JsonRpcClient) Do(method string, uri string, req any, rsp any, status ...int) error {
	furl := c.endpoint + uri
	body, err := c.encoder(req)
	if err != nil {
		return err
	}
	hreq, err := http.NewRequest(method, furl, content(body))
	if err != nil {
		return err
	}
	if c.header != nil {
		c.header.Init(furl, body, hreq.Header)
	}
	hreq.Header.Set("Content-Type", "application/json")
	hreq.ContentLength = int64(len(body))

	hrsp, err := c.client.Do(hreq)
	if err != nil {
		return err
	}
	defer DiscardResponse(hrsp)

	if len(status) > 0 {
		if !contains(hrsp.StatusCode, status) {
			return fmt.Errorf("unexpected status code: %v, expected %v", hrsp.StatusCode, status)
		}
	}
	if rsp != nil {
		return c.decoder(hrsp.Body, rsp)
	}
	return nil
}

// NewJsonRpcClient 创建rpc客户端
func NewJsonRpcClient(endpoint string, config *HttpConfig, header HttpHeader, encoder JsonMarshal, decoder JsonUnmarshal) *JsonRpcClient {
	if config == nil {
		config = new(HttpConfig)
	}
	if encoder == nil {
		encoder = EncodeProtoJsonData
	}
	if decoder == nil {
		decoder = DecodeProtoJsonReader
	}

	return &JsonRpcClient{
		endpoint: endpoint,
		client:   NewHttpClient(config),
		header:   header,
		encoder:  encoder,
		decoder:  decoder,
	}
}

func contains(p int, vs []int) bool {
	for _, v := range vs {
		if v == p {
			return true
		}
	}
	return false
}

func content(data []byte) io.Reader {
	if len(data) == 0 {
		return http.NoBody
	}
	return bytes.NewReader(data)
}
