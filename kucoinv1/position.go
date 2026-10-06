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

// Switch Position Mode
// https://www.kucoin.com/docs-new/rest/futures-trading/positions/switch-position-mode
type SwitchPositionMode struct {
	// the endpoint expects a string, ujson.Int64 always marshals as a quoted value
	PositionMode ujson.Int64
}

func (o SwitchPositionMode) Do(c *Client) Response[bool] {
	// futures position mode lives on api/v2 while the rest of the futures api is v1
	return Post(c.Copy().WithPath(ApiVersion2), "position/switchPositionMode", o, func(struct{}) (bool, error) {
		return true, nil
	})
}

func (o *Client) SwitchPositionMode(mode PositionMode) Response[bool] {
	return SwitchPositionMode{PositionMode: ujson.Int64(mode)}.Do(o)
}

// Get Position Mode
// https://www.kucoin.com/docs-new/rest/futures-trading/positions/get-position-mode
type GetPositionMode struct{}

func (o GetPositionMode) Do(c *Client) Response[PositionMode] {
	type result struct {
		// unlike switchPositionMode this endpoint answers with a number
		PositionMode ujson.Int64
	}
	return Get(c.Copy().WithPath(ApiVersion2), "position/getPositionMode", o, func(r result) (PositionMode, error) {
		return PositionMode(r.PositionMode.Value()), nil
	})
}

func (o *Client) GetPositionMode() Response[PositionMode] {
	return GetPositionMode{}.Do(o)
}

// CrossUserLeverage - response for GET /api/v2/getCrossUserLeverage
// https://www.kucoin.com/docs-new/rest/futures-trading/positions/get-cross-margin-leverage
type CrossUserLeverage struct {
	// Symbol - futures contract symbol
	Symbol string
	// Leverage - cross margin leverage of the account for the symbol (quoted number: "3")
	Leverage ujson.Int64
}

// Get Cross Margin Leverage
// https://www.kucoin.com/docs-new/rest/futures-trading/positions/get-cross-margin-leverage
type GetCrossUserLeverage struct {
	Symbol string `url:",omitempty"`
}

func (o GetCrossUserLeverage) Do(c *Client) Response[CrossUserLeverage] {
	// cross margin leverage lives on api/v2 while the rest of the futures api is v1
	return Get(c.Copy().WithPath(ApiVersion2), "getCrossUserLeverage", o, forward[CrossUserLeverage])
}

func (o *Client) GetCrossUserLeverage(symbol string) Response[CrossUserLeverage] {
	return GetCrossUserLeverage{
		Symbol: symbol,
	}.Do(o)
}

// Change Cross Margin Leverage
// https://www.kucoin.com/docs-new/rest/futures-trading/positions/modify-cross-margin-leverage
type ChangeCrossUserLeverage struct {
	// Symbol - futures contract symbol
	Symbol string
	// the endpoint expects a string, ujson.Int64 always marshals as a quoted value
	Leverage ujson.Int64
}

func (o ChangeCrossUserLeverage) Do(c *Client) Response[bool] {
	// changeCrossUserLeverage lives on api/v2 and answers with data: true
	return Post(c.Copy().WithPath(ApiVersion2), "changeCrossUserLeverage", o, forward[bool])
}

func (o *Client) ChangeCrossUserLeverage(symbol string, leverage int) Response[bool] {
	return ChangeCrossUserLeverage{
		Symbol:   symbol,
		Leverage: ujson.Int64(leverage),
	}.Do(o)
}
