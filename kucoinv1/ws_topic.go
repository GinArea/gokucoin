package kucoinv1

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/msw-x/moon/ujson"
)

type RawTopic Topic[json.RawMessage]

type Topic[T any] struct {
	Type        string
	Topic       string
	Subject     string
	Id          string
	UserId      string
	ChannelType string
	Sn          ujson.Int64
	Data        T
}

func UnmarshalRawTopic[T any](raw RawTopic) (ret Topic[T], err error) {
	ret.Type = raw.Type
	ret.Topic = raw.Topic
	ret.Subject = raw.Subject
	ret.Id = raw.Id
	ret.UserId = raw.UserId
	ret.ChannelType = raw.ChannelType
	ret.Sn = raw.Sn
	err = json.Unmarshal(raw.Data, &ret.Data)
	return
}

// Orderbook - Level 5 (Futures)
// https://www.kucoin.com/docs-new/3470083w0 (futures)
// https://www.kucoin.com/docs-new/3470069w0 (spot)
type Orderbook struct {
	Ask       [][]ujson.Float64 `json:"asks"`
	Bid       [][]ujson.Float64 `json:"bids"`
	Timestamp int64             `json:"timestamp"`
	Ts        int64             `json:"ts"`
	Sequence  int64             `json:"sequence"`
}

// Position Change Events
// https://www.kucoin.com/docs-new/3470093w0
type PositionShot struct {
	Position
	ChangeReason string
}

// Orders (Futures)
// https://www.kucoin.com/docs-new/3470090w0
type OrderShot struct {
	Symbol       string
	OrderType    OrderType
	TradeType    string
	Side         Side
	CanceledSize ujson.Float64
	OrderId      string
	PositionSide PositionSide
	Liquidity    string
	MarginMode   MarginMode
	Type         OrderTypeWs
	FeeType      string
	OldSize      ujson.Float64
	OrderTime    int64
	Size         ujson.Float64
	FilledSize   ujson.Float64
	Price        ujson.Float64
	MatchPrice   ujson.Float64
	MatchSize    ujson.Float64
	RemainSize   ujson.Float64
	TradeId      string
	ClientOid    string
	Status       OrderStatus
	Ts           int64
}

// Account Balance Events (Futures)
// https://www.kucoin.com/docs-new/3470092w0
type WalletShot struct {
	Currency                 string
	Equity                   ujson.StringFloat64
	WalletBalance            ujson.StringFloat64
	AvailableBalance         ujson.StringFloat64
	HoldBalance              ujson.StringFloat64
	TotalCrossMargin         ujson.StringFloat64
	CrossPosMargin           ujson.StringFloat64
	CrossOrderMargin         ujson.StringFloat64
	CrossUnPnl               ujson.Float64
	IsolatedPosMargin        ujson.Float64
	IsolatedOrderMargin      ujson.Float64
	IsolatedFundingFeeMargin ujson.Float64
	IsolatedUnPnl            ujson.Float64
	Version                  string
	Timestamp                ujson.Int64
	WithdrawHold             ujson.Float64
}

// RelationContext - context for balance change event (Spot API)
type RelationContext struct {
	// Symbol - trading pair symbol
	Symbol string `json:"symbol"`
	// OrderId - order ID that triggered the balance change
	OrderId string `json:"orderId"`
	// TradeId - trade ID that triggered the balance change
	TradeId string `json:"tradeId"`
}

// OrderShotSpot - Spot Order Events (Spot API)
// https://www.kucoin.com/docs-new/3470073w0
type OrderShotSpot struct {
	// ClientOid - unique order id created by users
	ClientOid string `json:"clientOid"`
	// OrderId - order ID assigned by the system
	OrderId string `json:"orderId"`
	// Symbol - trading pair symbol (e.g., BTC-USDT)
	Symbol string `json:"symbol"`
	// Side - order side: "buy" or "sell"
	Side Side `json:"side"`
	// OrderType - order type: "limit" or "market"
	OrderType OrderType `json:"orderType"`
	// Type - event type: "received", "open", "update", "match", "filled", "canceled"
	Type string `json:"type"`
	// Status - order status: "new", "open", "match", "done"
	Status string `json:"status"`
	// Price - order price (for limit orders)
	Price ujson.Float64 `json:"price"`
	// OriginSize - original order size
	OriginSize ujson.Float64 `json:"originSize"`
	// Size - current order size
	Size ujson.Float64 `json:"size"`
	// FilledSize - executed size
	FilledSize ujson.Float64 `json:"filledSize"`
	// CanceledSize - cancelled size
	CanceledSize ujson.Float64 `json:"canceledSize"`
	// RemainSize - remaining size
	RemainSize ujson.Float64 `json:"remainSize"`
	// OldSize - previous order size (only in "update" events)
	OldSize ujson.Float64 `json:"oldSize"`
	// MatchPrice - match price (only in "match" events)
	MatchPrice ujson.Float64 `json:"matchPrice"`
	// MatchSize - match size (only in "match" events)
	MatchSize ujson.Float64 `json:"matchSize"`
	// Liquidity - liquidity type: "taker" or "maker" (only in "match" events)
	Liquidity string `json:"liquidity"`
	// FeeType - fee type: "takerFee" or "makerFee" (only in "match" events)
	FeeType string `json:"feeType"`
	// TradeId - trade ID (only in "match" events)
	TradeId string `json:"tradeId"`
	// OrderTime - order placement time in milliseconds
	OrderTime ujson.Int64 `json:"orderTime"`
	// Ts - event timestamp in nanoseconds
	Ts ujson.Int64 `json:"ts"`
	// RemainFunds - remaining funds (only in "filled"/"canceled" events)
	RemainFunds ujson.Float64 `json:"remainFunds"`
	// not documented
	OriginFunds   ujson.Float64 `json:"originFunds"`
	CanceledFunds ujson.Float64 `json:"canceledFunds"`
	Funds         ujson.Float64 `json:"funds"`
	Pt            ujson.Int64   `json:"pt"`
}

// WalletShotSpot - Account Balance Events (Spot API)
// https://www.kucoin.com/docs-new/3470075w0
type WalletShotSpot struct {
	// AccountId - unique account identifier
	AccountId string `json:"accountId"`
	// Currency - the asset currency (e.g., USDT)
	Currency string `json:"currency"`
	// Total - total balance in the account
	Total ujson.Float64 `json:"total"`
	// Available - available balance for trading
	Available ujson.Float64 `json:"available"`
	// Hold - balance held for open orders
	Hold ujson.Float64 `json:"hold"`
	// AvailableChange - change in available balance (can be negative)
	AvailableChange ujson.Float64 `json:"availableChange"`
	// HoldChange - change in held balance
	HoldChange ujson.Float64 `json:"holdChange"`
	// RelationContext - context about what triggered the balance change
	RelationContext RelationContext `json:"relationContext"`
	// RelationEvent - event type that triggered the change (e.g., "trade.hold")
	RelationEvent string `json:"relationEvent"`
	// RelationEventId - unique ID for the relation event
	RelationEventId string `json:"relationEventId"`
	// Time - timestamp of the event in milliseconds
	Time ujson.Int64 `json:"time"`
}

// CandleShotSpot - Spot Candles/Klines (Spot API)
// https://www.kucoin.com/docs-new/3470071w0
type CandleShotSpot struct {
	// Symbol - trading pair symbol (e.g., BTC-USDT)
	Symbol string `json:"symbol"`
	// Candles - OHLCV data as string array: [time, open, close, high, low, volume, turnover]
	Candles []string `json:"candles"`
	// Time - event timestamp in nanoseconds
	Time int64 `json:"time"`
}

// Kline converts the string array candles data to a typed Kline struct
// Note: WS order is [ts, open, close, high, low, volume, turnover]
func (o *CandleShotSpot) Kline() (Kline, error) {
	if len(o.Candles) < 7 {
		return Kline{}, fmt.Errorf("candles: expected 7 elements, got %d", len(o.Candles))
	}
	ts, err := strconv.ParseInt(o.Candles[0], 10, 64)
	if err != nil {
		return Kline{}, fmt.Errorf("candles[0] timestamp: %w", err)
	}
	open, err := strconv.ParseFloat(o.Candles[1], 64)
	if err != nil {
		return Kline{}, fmt.Errorf("candles[1] open: %w", err)
	}
	close_, err := strconv.ParseFloat(o.Candles[2], 64)
	if err != nil {
		return Kline{}, fmt.Errorf("candles[2] close: %w", err)
	}
	high, err := strconv.ParseFloat(o.Candles[3], 64)
	if err != nil {
		return Kline{}, fmt.Errorf("candles[3] high: %w", err)
	}
	low, err := strconv.ParseFloat(o.Candles[4], 64)
	if err != nil {
		return Kline{}, fmt.Errorf("candles[4] low: %w", err)
	}
	volume, err := strconv.ParseFloat(o.Candles[5], 64)
	if err != nil {
		return Kline{}, fmt.Errorf("candles[5] volume: %w", err)
	}
	turnover, err := strconv.ParseFloat(o.Candles[6], 64)
	if err != nil {
		return Kline{}, fmt.Errorf("candles[6] turnover: %w", err)
	}
	return Kline{
		Ts:       ts,
		Open:     open,
		High:     high,
		Low:      low,
		Close:    close_,
		Volume:   volume,
		Turnover: turnover,
	}, nil
}
