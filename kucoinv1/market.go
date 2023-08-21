package kucoinv1

import (
	"github.com/msw-x/moon/ujson"
)

// Get Open Contract List
// https://docs.kucoin.com/futures/#get-open-contract-list
type GetOpenContracts struct {
}

type OpenContract struct {
	Symbol                  string
	RootSymbol              string
	Type                    string
	FirstOpenDate           string
	ExpireDate              string
	SettleDate              string
	BaseCurrency            string
	QuoteCurrency           string
	SettleCurrency          string
	MaxOrderQty             ujson.Float64
	MaxPrice                ujson.Float64
	LotSize                 ujson.Float64
	TickSize                ujson.Float64
	IndexPriceTickSize      ujson.Float64
	Multiplier              ujson.Float64
	InitialMargin           ujson.Float64
	MaintainMargin          ujson.Float64
	MaxRiskLimit            ujson.Float64
	MinRiskLimit            ujson.Float64
	RiskStep                ujson.Float64
	MakerFeeRate            ujson.Float64
	TakerFeeRate            ujson.Float64
	TakerFixFee             ujson.Float64
	MakerFixFee             ujson.Float64
	SettlementFee           ujson.Float64
	IsDeleverage            bool
	IsQuanto                bool
	IsInverse               bool
	MarkMethod              string
	FairMethod              string
	FundingBaseSymbol       string
	FundingQuoteSymbol      string
	FundingRateSymbol       string
	IndexSymbol             string
	SettlementSymbol        string
	Status                  ContractStatus
	FundingFeeRate          ujson.Float64
	PredictedFundingFeeRate ujson.Float64
	OpenInterest            string
	TurnoverOf24h           ujson.Float64
	VolumeOf24h             ujson.Float64
	MarkPrice               ujson.Float64
	IndexPrice              ujson.Float64
	LastTradePrice          ujson.Float64
	NextFundingRateTime     ujson.Float64
	MaxLeverage             ujson.Float64
	SourceExchanges         []string
	PremiumsSymbol1M        string
	PremiumsSymbol8H        string
	FundingBaseSymbol1M     string
	FundingQuoteSymbol1M    string
	LowPrice                ujson.Float64
	HighPrice               ujson.Float64
	PriceChgPct             ujson.Float64
	PriceChg                ujson.Float64
}

func (o GetOpenContracts) Do(c *Client) Response[[]OpenContract] {
	return GetPub(c.contracts(), "active", o, forward[[]OpenContract])
}

func (o GetOpenContracts) DoSingle(c *Client, market string) Response[OpenContract] {
	return GetPub(c.contracts(), market, o, forward[OpenContract])
}

func (o *Client) GetOpenContracts() Response[[]OpenContract] {
	return GetOpenContracts{}.Do(o)
}

func (o *Client) GetOpenContract(market string) Response[OpenContract] {
	return GetOpenContracts{}.DoSingle(o, market)
}

// Get Real-Time Ticker
// https://docs.kucoin.com/futures/#get-real-time-ticker
//
//	ticker -> String [Symbol of the contract]
type GetTicker struct {
	Symbol string
}

type Ticker struct {
	Sequence     int64
	Symbol       string
	Side         string
	Size         ujson.Float64
	Price        ujson.Float64
	BestBidSize  ujson.Float64
	BestBidPrice ujson.Float64
	BestAskSize  ujson.Float64
	BestAskPrice ujson.Float64
	TradeId      string
	Ts           int64
}

func (o GetTicker) Do(c *Client) Response[Ticker] {
	return GetPub(c, "ticker", o, forward[Ticker])
}

func (o *Client) GetTicker(symbol string) Response[Ticker] {
	return GetTicker{
		Symbol: symbol,
	}.Do(o)
}
