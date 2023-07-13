package v1

import (
	"github.com/msw-x/moon/uhttp"
	"github.com/msw-x/moon/ujson"
)

// Get Account Overview
// https://docs.kucoin.com/futures/#get-account-overview

//parameters
//currency	String	[Optional] Currecny ,including XBT,USDT,Default XBT

type GetAccountOverview struct {
	Currency string `url:",omitempty"`
}

type AccountOverview struct {
	AccountEquity    ujson.Float64
	UnrealisedPNL    ujson.Float64
	MmarginBalance   ujson.Float64
	PositionMargin   ujson.Float64
	OrderMargin      ujson.Float64
	FrozenFunds      ujson.Float64
	AvailableBalance ujson.Float64
	Currency         string
}

func (o *Client) GetAccountOverview(v GetAccountOverview) Response[AccountOverview] {
	return v.Do(o)
}

func (o GetAccountOverview) Do(c *Client) Response[AccountOverview] {
	return getAccountOverview[AccountOverview](o, c)
}

func getAccountOverview[T any](o GetAccountOverview, c *Client) Response[T] {
	return Get[T](c, "account-overview", o, func(h uhttp.Responce) (r Response[T], er error) {
		if h.BodyExists() {
			raw := new(item[T])
			h.Json(raw)
			r.Time = getCurrentTime()
			r.Error = raw.Error()
			if r.Ok() {
				r.Data = []T{raw.Data}
			}
		}
		return
	})
}
