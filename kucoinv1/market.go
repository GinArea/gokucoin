package kucoinv1

import (
	"fmt"

	"github.com/msw-x/moon/ujson"
)

// Contract represents a futures contract specification with market data
// https://www.kucoin.com/docs-new/rest/futures-trading/market-data/get-all-symbols
type Contract struct {
	// Basic info
	Symbol                            string
	RootSymbol                        string
	Type                              string
	FirstOpenDate                     ujson.Int64
	ExpireDate                        ujson.Int64
	SettleDate                        ujson.Int64
	BaseCurrency                      string
	QuoteCurrency                     string
	SettleCurrency                    string
	MaxOrderQty                       ujson.Int64
	MaxPrice                          ujson.Float64
	LotSize                           ujson.Int64
	TickSize                          ujson.Float64
	IndexPriceTickSize                ujson.Float64
	Multiplier                        ujson.Float64
	InitialMargin                     ujson.Float64
	MaintainMargin                    ujson.Float64
	MaxRiskLimit                      ujson.Int64
	MinRiskLimit                      ujson.Int64
	RiskStep                          ujson.Int64
	MakerFeeRate                      ujson.Float64
	TakerFeeRate                      ujson.Float64
	TakerFixFee                       ujson.Float64
	MakerFixFee                       ujson.Float64
	SettlementFee                     ujson.Float64
	IsDeleverage                      bool
	IsQuanto                          bool
	IsInverse                         bool
	MarkMethod                        string
	FairMethod                        string
	FundingBaseSymbol                 string
	FundingQuoteSymbol                string
	FundingRateSymbol                 string
	IndexSymbol                       string
	SettlementSymbol                  string
	Status                            string
	FundingFeeRate                    ujson.Float64
	PredictedFundingFeeRate           ujson.Float64
	FundingRateGranularity            ujson.Int64
	OpenInterest                      string
	TurnoverOf24h                     ujson.Float64
	VolumeOf24h                       ujson.Float64
	MarkPrice                         ujson.Float64
	IndexPrice                        ujson.Float64
	LastTradePrice                    ujson.Float64
	NextFundingRateTime               ujson.Int64
	MaxLeverage                       ujson.Int64
	SourceExchanges                   []string
	PremiumsSymbol1M                  string
	PremiumsSymbol8H                  string
	FundingBaseSymbol1M               string
	FundingQuoteSymbol1M              string
	LowPrice                          ujson.Float64
	HighPrice                         ujson.Float64
	PriceChgPct                       ujson.Float64
	PriceChg                          ujson.Float64
	DisplaySymbol                     string
	DisplayBaseCurrency               string
	MarketMaxOrderQty                 ujson.Int64
	DailyInterestRate                 ujson.Float64
	FundingRateCap                    ujson.Float64
	FundingRateFloor                  ujson.Float64
	Period                            ujson.Int64
	EffectiveFundingRateCycleStartTime ujson.Int64
	CurrentFundingRateGranularity     ujson.Int64
	NextFundingRateDateTime           ujson.Int64
	K                                 ujson.Float64
	M                                 ujson.Float64
	F                                 ujson.Float64
	MmrLimit                          ujson.Float64
	MmrLevConstant                    ujson.Float64
	SupportCross                      bool
	BuyLimit                          ujson.Float64
	SellLimit                         ujson.Float64
	CrossRiskLimit                    ujson.Float64
	MarketStage                       string
	OrderPriceRange                   ujson.Float64
	AdjustK                           ujson.Float64
	AdjustM                           ujson.Float64
	AdjustMmrLevConstant              ujson.Float64
	AdjustActiveTime                  ujson.Int64
	PreMarketToPerpDate               ujson.Int64
}

// Ticker represents real-time ticker data
// https://www.kucoin.com/docs-new/rest/futures-trading/market-data/get-ticker
type Ticker struct {
	Sequence     ujson.Int64
	Symbol       string
	Side         string
	Size         ujson.Int64
	TradeId      string
	Price        ujson.Float64
	BestBidPrice ujson.Float64
	BestBidSize  ujson.Int64
	BestAskPrice ujson.Float64
	BestAskSize  ujson.Int64
	Ts           ujson.Int64
}

// GetContracts retrieves all active futures contracts
// GET /api/v1/contracts/active
type GetContracts struct{}

func (o GetContracts) Do(c *Client) Response[[]Contract] {
	return GetPub(c.contracts(), "active", o, forward[[]Contract])
}

func (o *Client) GetContracts() Response[[]Contract] {
	return GetContracts{}.Do(o)
}

// GetContract retrieves a single contract by symbol
// GET /api/v1/contracts/{symbol}
type GetContract struct {
	Symbol string
}

func (o GetContract) Do(c *Client) Response[Contract] {
	return GetPub(c.contracts(), o.Symbol, struct{}{}, forward[Contract])
}

func (o *Client) GetContract(symbol string) Response[Contract] {
	return GetContract{Symbol: symbol}.Do(o)
}

// GetTicker retrieves real-time ticker for a symbol
// GET /api/v1/ticker?symbol=XBTUSDTM
type GetTicker struct {
	Symbol string `url:"symbol"`
}

func (o GetTicker) Do(c *Client) Response[Ticker] {
	return GetPub(c.ticker(), "", o, forward[Ticker])
}

func (o *Client) GetTicker(symbol string) Response[Ticker] {
	return GetTicker{Symbol: symbol}.Do(o)
}

// GetTickers retrieves all tickers
// GET /api/v1/allTickers
// https://www.kucoin.com/docs-new/rest/futures-trading/market-data/get-all-tickers
type GetTickers struct{}

func (o GetTickers) Do(c *Client) Response[[]Ticker] {
	return GetPub(c, "allTickers", o, forward[[]Ticker])
}

func (o *Client) GetTickers() Response[[]Ticker] {
	return GetTickers{}.Do(o)
}

// Kline represents a single candlestick data point
type Kline struct {
	Ts       int64
	Open     float64
	High     float64
	Low      float64
	Close    float64
	Volume   float64
	Turnover float64
}

// GetKlines retrieves candlestick/kline data
// GET /api/v1/kline/query
// https://www.kucoin.com/docs-new/rest/futures-trading/market-data/get-klines
type GetKlines struct {
	Symbol      string `url:"symbol"`
	Granularity int    `url:"granularity"`
	From        int64  `url:"from,omitempty"`
	To          int64  `url:"to,omitempty"`
}

func (o GetKlines) Do(c *Client) Response[[]Kline] {
	return GetPub(c.kline(), "query", o, transformKlines)
}

func (o *Client) GetKlines(v GetKlines) Response[[]Kline] {
	return v.Do(o)
}

func transformKlines(data [][]any) ([]Kline, error) {
	result := make([]Kline, len(data))
	for i, row := range data {
		if len(row) < 7 {
			return nil, fmt.Errorf("kline row %d: expected 7 elements, got %d", i, len(row))
		}
		k := Kline{}
		// Timestamp (int64 ms)
		if v, ok := row[0].(float64); ok {
			k.Ts = int64(v)
		} else {
			return nil, fmt.Errorf("kline row %d: invalid timestamp type", i)
		}
		// Open
		if v, ok := row[1].(float64); ok {
			k.Open = v
		} else {
			return nil, fmt.Errorf("kline row %d: invalid open type", i)
		}
		// High
		if v, ok := row[2].(float64); ok {
			k.High = v
		} else {
			return nil, fmt.Errorf("kline row %d: invalid high type", i)
		}
		// Low
		if v, ok := row[3].(float64); ok {
			k.Low = v
		} else {
			return nil, fmt.Errorf("kline row %d: invalid low type", i)
		}
		// Close
		if v, ok := row[4].(float64); ok {
			k.Close = v
		} else {
			return nil, fmt.Errorf("kline row %d: invalid close type", i)
		}
		// Volume
		if v, ok := row[5].(float64); ok {
			k.Volume = v
		} else {
			return nil, fmt.Errorf("kline row %d: invalid volume type", i)
		}
		// Turnover
		if v, ok := row[6].(float64); ok {
			k.Turnover = v
		} else {
			return nil, fmt.Errorf("kline row %d: invalid turnover type", i)
		}
		result[i] = k
	}
	return result, nil
}
