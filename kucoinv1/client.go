package kucoinv1

import (
	"net/http"
	"time"

	"github.com/msw-x/moon/uhttp"
)

type Client struct {
	c                *uhttp.Client
	s                *Sign
	onTransportError OnTransportError
}

func NewClient() *Client {
	o := new(Client)
	o.c = uhttp.NewClient()
	o.WithBaseUrl(MainBaseUrl)
	o.WithPath(ApiVersion)
	return o
}

func (o *Client) WithTimeout(timeout time.Duration) *Client {
	o.c.WithTimeout(timeout)
	return o
}

func (o *Client) WithTrace(trace func(uhttp.Response)) *Client {
	o.c.WithTrace(trace)
	return o
}

func (o *Client) WithTransport(tranport *http.Transport) *Client {
	o.c.WithTransport(tranport)
	return o
}

func (o *Client) WithProxy(proxy string) *Client {
	o.c.WithProxy(proxy)
	return o
}

func (o *Client) Copy() *Client {
	r := new(Client)
	r.c = o.c.Copy()
	r.s = o.s
	r.onTransportError = o.onTransportError
	return r
}

func (o *Client) WithBaseUrl(url string) *Client {
	o.c.WithBase(url)
	return o
}

func (o *Client) WithPath(path string) *Client {
	o.c.WithPath(path)
	return o
}

func (o *Client) WithAppendPath(path string) *Client {
	o.c.WithAppendPath(path)
	return o
}

func (o *Client) WithAuth(key, secret, password string) *Client {
	o.s = NewSign(key, secret, password)
	return o
}

func (o *Client) WithOnTransportError(f OnTransportError) *Client {
	o.onTransportError = f
	return o
}

func (o *Client) contracts() *Client {
	return o.Copy().WithAppendPath("contracts")
}

func (o *Client) orders() *Client {
	return o.Copy().WithAppendPath("orders")
}

func (o *Client) stopOrders() *Client {
	return o.Copy().WithAppendPath("stopOrders")
}

func (o *Client) positions() *Client {
	return o.Copy().WithAppendPath("positions")
}

func (o *Client) accountOverview() *Client {
	return o.Copy().WithAppendPath("account-overview")
}

func (o *Client) fundingHistory() *Client {
	return o.Copy().WithAppendPath("funding-history")
}

func (o *Client) bulletPrivate() *Client {
	return o.Copy().WithAppendPath("bullet-private")
}

func (o *Client) ticker() *Client {
	return o.Copy().WithAppendPath("ticker")
}

func (o *Client) kline() *Client {
	return o.Copy().WithAppendPath("kline")
}

type OnTransportError func(err error, method string, statusCode int, attempt int) bool
