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

// ============================================================================
// Spot Market Data
// ============================================================================

// SymbolSpot - spot trading pair specification
// https://www.kucoin.com/docs-new/rest/spot-trading/market-data/get-all-symbols
type SymbolSpot struct {
	// Symbol - Trading pair identifier (e.g., "BTC-USDT")
	Symbol string `json:"symbol"`
	// Name - Display name
	Name string `json:"name"`
	// BaseCurrency - Base currency code
	BaseCurrency string `json:"baseCurrency"`
	// QuoteCurrency - Quote currency code
	QuoteCurrency string `json:"quoteCurrency"`
	// FeeCurrency - Currency used for fees
	FeeCurrency string `json:"feeCurrency"`
	// Market - Market classification
	Market string `json:"market"`
	// BaseMinSize - Minimum base currency amount
	BaseMinSize ujson.Float64 `json:"baseMinSize"`
	// QuoteMinSize - Minimum quote currency amount
	QuoteMinSize ujson.Float64 `json:"quoteMinSize"`
	// BaseMaxSize - Maximum base currency amount
	BaseMaxSize ujson.Float64 `json:"baseMaxSize"`
	// QuoteMaxSize - Maximum quote currency amount
	QuoteMaxSize ujson.Float64 `json:"quoteMaxSize"`
	// BaseIncrement - Base currency precision increment
	BaseIncrement ujson.Float64 `json:"baseIncrement"`
	// QuoteIncrement - Quote currency precision increment
	QuoteIncrement ujson.Float64 `json:"quoteIncrement"`
	// PriceIncrement - Price precision increment
	PriceIncrement ujson.Float64 `json:"priceIncrement"`
	// PriceLimitRate - Price limit rate
	PriceLimitRate ujson.Float64 `json:"priceLimitRate"`
	// MinFunds - Minimum order value
	MinFunds ujson.Float64 `json:"minFunds"`
	// IsMarginEnabled - Margin trading availability
	IsMarginEnabled bool `json:"isMarginEnabled"`
	// EnableTrading - Trading enabled status
	EnableTrading bool `json:"enableTrading"`
	// FeeCategory - Fee tier classification
	FeeCategory int `json:"feeCategory"`
	// MakerFeeCoefficient - Maker fee multiplier
	MakerFeeCoefficient ujson.Float64 `json:"makerFeeCoefficient"`
	// TakerFeeCoefficient - Taker fee multiplier
	TakerFeeCoefficient ujson.Float64 `json:"takerFeeCoefficient"`
	// St - Special trading flag
	St bool `json:"st"`
	// CallauctionIsEnabled - Call auction availability
	CallauctionIsEnabled bool `json:"callauctionIsEnabled"`
	// CallauctionPriceFloor - Call auction minimum price (nullable)
	CallauctionPriceFloor *string `json:"callauctionPriceFloor"`
	// CallauctionPriceCeiling - Call auction maximum price (nullable)
	CallauctionPriceCeiling *string `json:"callauctionPriceCeiling"`
	// CallauctionFirstStageStartTime - First stage auction start time (nullable)
	CallauctionFirstStageStartTime *ujson.Int64 `json:"callauctionFirstStageStartTime"`
	// CallauctionSecondStageStartTime - Second stage auction start time (nullable)
	CallauctionSecondStageStartTime *ujson.Int64 `json:"callauctionSecondStageStartTime"`
	// CallauctionThirdStageStartTime - Third stage auction start time (nullable)
	CallauctionThirdStageStartTime *ujson.Int64 `json:"callauctionThirdStageStartTime"`
	// TradingStartTime - Trading commencement timestamp (nullable)
	TradingStartTime *ujson.Int64 `json:"tradingStartTime"`
}

// TickerSpot - spot Level 1 market data
// https://www.kucoin.com/docs-new/rest/spot-trading/market-data/get-ticker
type TickerSpot struct {
	// Time - Server timestamp in milliseconds
	Time ujson.Int64 `json:"time"`
	// Sequence - Message sequence number
	Sequence string `json:"sequence"`
	// Price - Last traded price
	Price ujson.Float64 `json:"price"`
	// Size - Last traded size
	Size ujson.Float64 `json:"size"`
	// BestBid - Best bid price
	BestBid ujson.Float64 `json:"bestBid"`
	// BestBidSize - Best bid size
	BestBidSize ujson.Float64 `json:"bestBidSize"`
	// BestAsk - Best ask price
	BestAsk ujson.Float64 `json:"bestAsk"`
	// BestAskSize - Best ask size
	BestAskSize ujson.Float64 `json:"bestAskSize"`
}

// TickerSpotItem - item in Get All Tickers response
// https://www.kucoin.com/docs-new/rest/spot-trading/market-data/get-all-tickers
type TickerSpotItem struct {
	// Symbol - Trading pair
	Symbol string `json:"symbol"`
	// SymbolName - Symbol display name
	SymbolName string `json:"symbolName"`
	// Buy - Best bid price
	Buy ujson.Float64 `json:"buy"`
	// BestBidSize - Best bid size
	BestBidSize ujson.Float64 `json:"bestBidSize"`
	// Sell - Best ask price
	Sell ujson.Float64 `json:"sell"`
	// BestAskSize - Best ask size
	BestAskSize ujson.Float64 `json:"bestAskSize"`
	// ChangeRate - 24h change rate
	ChangeRate ujson.Float64 `json:"changeRate"`
	// ChangePrice - 24h change price
	ChangePrice ujson.Float64 `json:"changePrice"`
	// High - 24h high
	High ujson.Float64 `json:"high"`
	// Low - 24h low
	Low ujson.Float64 `json:"low"`
	// Vol - 24h volume in base currency
	Vol ujson.Float64 `json:"vol"`
	// VolValue - 24h volume in quote currency
	VolValue ujson.Float64 `json:"volValue"`
	// Last - Last trade price
	Last ujson.Float64 `json:"last"`
	// AveragePrice - 24h weighted average price
	AveragePrice ujson.Float64 `json:"averagePrice"`
	// TakerFeeRate - Taker fee rate
	TakerFeeRate ujson.Float64 `json:"takerFeeRate"`
	// MakerFeeRate - Maker fee rate
	MakerFeeRate ujson.Float64 `json:"makerFeeRate"`
	// TakerCoefficient - Taker fee coefficient
	TakerCoefficient ujson.Float64 `json:"takerCoefficient"`
	// MakerCoefficient - Maker fee coefficient
	MakerCoefficient ujson.Float64 `json:"makerCoefficient"`
}

// TickersSpot - response wrapper for Get All Tickers
// https://www.kucoin.com/docs-new/rest/spot-trading/market-data/get-all-tickers
type TickersSpot struct {
	// Time - Server timestamp in milliseconds
	Time ujson.Int64 `json:"time"`
	// Ticker - Array of tickers
	Ticker []TickerSpotItem `json:"ticker"`
}

// GetSymbolsSpot retrieves all spot trading symbols
// GET /api/v2/symbols
// https://www.kucoin.com/docs-new/rest/spot-trading/market-data/get-all-symbols
type GetSymbolsSpot struct {
	// Market - Optional market filter (e.g., "USDS", "BTC")
	Market string `url:"market,omitempty"`
}

func (o GetSymbolsSpot) Do(c *Client) Response[[]SymbolSpot] {
	cc := c.Copy().WithBaseUrl(SpotBaseUrl).WithPath("api/v2")
	return GetPub(cc, "symbols", o, forward[[]SymbolSpot])
}

func (o *Client) GetSymbolsSpot() Response[[]SymbolSpot] {
	return GetSymbolsSpot{}.Do(o)
}

// GetSymbolSpot retrieves a single spot symbol by name
// GET /api/v2/symbols/{symbol}
// https://www.kucoin.com/docs-new/rest/spot-trading/market-data/get-symbol
type GetSymbolSpot struct {
	Symbol string
}

func (o GetSymbolSpot) Do(c *Client) Response[SymbolSpot] {
	cc := c.Copy().WithBaseUrl(SpotBaseUrl).WithPath("api/v2")
	return GetPub(cc, "symbols/"+o.Symbol, struct{}{}, forward[SymbolSpot])
}

func (o *Client) GetSymbolSpot(symbol string) Response[SymbolSpot] {
	return GetSymbolSpot{Symbol: symbol}.Do(o)
}

// GetTickersSpot retrieves all spot tickers
// GET /api/v1/market/allTickers
// https://www.kucoin.com/docs-new/rest/spot-trading/market-data/get-all-tickers
type GetTickersSpot struct{}

func (o GetTickersSpot) Do(c *Client) Response[TickersSpot] {
	cc := c.Copy().WithBaseUrl(SpotBaseUrl).WithPath("api/v1")
	return GetPub(cc, "market/allTickers", o, forward[TickersSpot])
}

func (o *Client) GetTickersSpot() Response[TickersSpot] {
	return GetTickersSpot{}.Do(o)
}

// GetTickerSpot retrieves Level 1 ticker for a symbol
// GET /api/v1/market/orderbook/level1?symbol=
// https://www.kucoin.com/docs-new/rest/spot-trading/market-data/get-ticker
type GetTickerSpot struct {
	Symbol string `url:"symbol"`
}

func (o GetTickerSpot) Do(c *Client) Response[TickerSpot] {
	cc := c.Copy().WithBaseUrl(SpotBaseUrl).WithPath("api/v1")
	return GetPub(cc, "market/orderbook/level1", o, forward[TickerSpot])
}

func (o *Client) GetTickerSpot(symbol string) Response[TickerSpot] {
	return GetTickerSpot{Symbol: symbol}.Do(o)
}
