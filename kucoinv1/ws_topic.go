package kucoinv1

import (
	"encoding/json"

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
	err = json.Unmarshal(raw.Data, &ret.Data)
	return
}

// Orderbook - Level 5 (Futures)
// https://www.kucoin.com/docs-new/3470083w0
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
