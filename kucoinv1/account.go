package kucoinv1

import (
	"github.com/msw-x/moon/ujson"
)

// Get Account Overview
// https://docs.kucoin.com/futures/#get-account-overview
//
//	currency String [Optional] Currecny, including XBT,USDT,Default XBT
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

func (o GetAccountOverview) Do(c *Client) Response[AccountOverview] {
	return Get(c, "account-overview", o, forward[AccountOverview])
}

func (o *Client) GetAccountOverview(currency string) Response[AccountOverview] {
	return GetAccountOverview{
		Currency: currency,
	}.Do(o)
}
