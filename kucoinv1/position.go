package kucoinv1

import (
	"github.com/msw-x/moon/ujson"
)

// Get Position List
// https://www.kucoin.com/docs-new/rest/futures-trading/positions/get-position-list
type GetPositions struct {
	Currency string `url:",omitempty"`
}

type Position struct {
	Id                string
	Symbol            string
	AutoDeposit       bool
	CrossMode         bool
	MaintMarginReq    ujson.Float64
	RiskLimit         ujson.Float64
	RealLeverage      ujson.Float64
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
	PosCrossMargin    ujson.Float64
	PosInit           ujson.Float64
	PosComm           ujson.Float64
	PosCommCommon     ujson.Float64
	PosLoss           ujson.Float64
	PosMargin         ujson.Float64
	PosFunding        ujson.Float64
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
	MarginMode        string
	PositionSide      PositionSide
	Leverage          ujson.Float64
	DealComm          ujson.Float64
	FundingFee        ujson.Float64
	Tax               ujson.Float64
	WithdrawPnl       ujson.Float64
}

func (o GetPositions) Do(c *Client) Response[[]Position] {
	return Get(c, "positions", o, forward[[]Position])
}

func (o *Client) GetPositions(currency string) Response[[]Position] {
	return GetPositions{
		Currency: currency,
	}.Do(o)
}
