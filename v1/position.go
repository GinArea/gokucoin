package v1

import (
	"github.com/msw-x/moon/uhttp"
	"github.com/msw-x/moon/ujson"
)

// Get Position List
// https://docs.kucoin.com/futures/#get-position-list
type GetPositions struct {
	Currency string `url:",omitempty"`
}

type Position struct {
	Id                string
	Symbol            string
	AutoDeposit       bool
	MaintMarginReq    ujson.Float64
	RiskLimit         ujson.Float64
	RealLeverage      ujson.Float64
	CrossMode         bool
	DelevPercentage   ujson.Float64
	OpeningTimestamp  ujson.Int64
	CurrentTimestamp  ujson.Int64
	CurrentQty        ujson.Float64
	CurrentCost       ujson.Float64
	CurrentComm       ujson.Float64
	UnrealisedCost    ujson.Float64
	RealisedGrossCost ujson.Float64
	RealisedCost      ujson.Float64
	IsOpen            bool
	MarkPrice         ujson.Float64
	MarkValue         ujson.Float64
	PosCost           ujson.Float64
	PosCross          ujson.Float64
	PosInit           ujson.Float64
	PosComm           ujson.Float64
	PosLoss           ujson.Float64
	PosMargin         ujson.Float64
	PosMaint          ujson.Float64
	MaintMargin       ujson.Float64
	RealisedGrossPnl  ujson.Float64
	RealisedPnl       ujson.Float64
	UnrealisedPnl     ujson.Float64
	UnrealisedPnlPcnt ujson.Float64
	UnrealisedRoePcnt ujson.Float64
	AvgEntryPrice     ujson.Float64
	LiquidationPrice  ujson.Float64
	BankruptPrice     ujson.Float64
	SettleCurrency    string
	IsInverse         bool
	MaintainMargin    ujson.Float64
	UserId            ujson.Float64
}

func (o *Client) GetPositions(v GetPositions) Response[Position] {
	return v.Do(o)
}

func (o GetPositions) Do(c *Client) Response[Position] {
	return getPositions[Position](o, c)
}

func getPositions[T any](o GetPositions, c *Client) Response[T] {
	return Get[T](c, "positions", o, func(h uhttp.Responce) (r Response[T], er error) {
		if h.BodyExists() {
			raw := new(response[T])
			h.Json(raw)
			r.Time = getCurrentTime()
			r.Error = raw.Error()
			if r.Ok() {
				r.Data = raw.Data
			}
		}
		return
	})
}
