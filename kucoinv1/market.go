package kucoinv1

import (
	"github.com/msw-x/moon/uhttp"
	"github.com/msw-x/moon/ujson"
)

// Get Open Contract List <- получение информации по всем маркетам (Submit request to get the info of all open contracts)
// https://docs.kucoin.com/futures/#get-open-contract-list
//
// no parameters
type GetOpenContracts struct{}

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

func getOpenContracts[T any](o GetOpenContracts, c *Client) Response[T] {
	return GetPub[T](c.contracts(), "active", o, func(h uhttp.Responce) (r Response[T], er error) {
		if h.BodyExists() {
			raw := new(response[T])
			h.Json(raw)
			r.Error = raw.Error()
			if r.Ok() {
				r.Data, r.Error = raw.Data, nil
			}
		}
		return
	})
}

func (o GetOpenContracts) Do(c *Client) Response[OpenContract] {
	return getOpenContracts[OpenContract](o, c)
}

func (o *Client) GetOpenContracts(v GetOpenContracts) Response[OpenContract] {
	return v.Do(o)
}

func (o *Client) GetOpenContract(market string) Response[OpenContract] {
	return getOpenContract[OpenContract](market, o)
}

func getOpenContract[T any](market string, c *Client) Response[T] {
	return GetPub[T](c.contracts(), market, GetOpenContracts{}, func(h uhttp.Responce) (r Response[T], er error) {
		if h.BodyExists() {
			raw := new(item[T])
			h.Json(raw)
			r.Error = raw.Error()
			if r.Ok() {
				r.Data, r.Error = []T{raw.Data}, nil
			}
		}
		return
	})
}

// Get Real-Time Ticker <- получение информации по ценам переданного тикера
// https://docs.kucoin.com/futures/#get-real-time-ticker
//
// parameters
// ticker -> String [Symbol of the contract]
type GetTicker struct {
	Symbol string `url:",omitempty"`
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

func (o *Client) GetTicker(v GetTicker) Response[Ticker] {
	return v.Do(o)
}

func (o GetTicker) Do(c *Client) Response[Ticker] {
	return getTickers[Ticker](o, c)
}

func getTickers[T any](o GetTicker, c *Client) Response[T] {
	return GetPub[T](c, "ticker", o, func(h uhttp.Responce) (r Response[T], er error) {
		if h.BodyExists() {
			raw := new(item[T])
			h.Json(raw)
			r.Error = raw.Error()
			if r.Ok() {
				r.Data = []T{raw.Data}
			}
		}
		return
	})
}

// Get Full Order Book - Level 2
// https://docs.kucoin.com/futures/#get-full-order-book-level-2

type Orderbook struct {
	Symbol    string            `json:"symbol"`
	Ask       [][]ujson.Float64 `json:"asks"`
	Bid       [][]ujson.Float64 `json:"bids"`
	Timestamp int               `json:"ts"`
	Sequence  int               `json:"sequence"`
}
