package kucoinv1

import "github.com/msw-x/moon/uhttp"

type Client struct {
	c          *uhttp.Client
	s          *Sign
	onNetError func(err error, attempt int) bool
}

func NewClient() *Client {
	o := new(Client)
	o.c = uhttp.NewClient()
	o.WithBaseUrl(MainBaseUrl)
	o.WithPath(ApiVersion)
	return o
}

func (o *Client) Clone() *Client {
	r := new(Client)
	r.c = o.c.Clone()
	r.s = o.s
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

func (o *Client) WithOnNetError(onNetError func(err error, attempt int) bool) *Client {
	o.onNetError = onNetError
	return o
}

func (o *Client) contracts() *Client {
	return o.Clone().WithAppendPath("contracts")
}
