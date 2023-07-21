package kucoinv1

import (
	"encoding/json"

	"github.com/msw-x/moon/ujson"
)

type RawTopic Topic[json.RawMessage]

type Topic[T any] struct {
	Type    string
	Topic   string
	Subject string
	Data    T
}

func UnmarshalRawTopic[T any](raw RawTopic) (ret Topic[T], err error) {
	ret.Type = raw.Type
	ret.Topic = raw.Topic
	ret.Subject = raw.Subject
	err = json.Unmarshal(raw.Data, &ret.Data)
	return
}

// Position Change Events
// https://docs.kucoin.com/futures/#position-change-events
type PositionShot struct {
	Position
	ChangeReason string
}

// Trade Orders
// https://docs.kucoin.com/futures/#trade-orders
type OrderShot struct {
	OrderId      string
	Symbol       string
	Type         OrderStatusType
	Status       TradeOrderStatus
	MatchSize    ujson.Float64
	MatchPrice   ujson.Float64
	OrderType    OrderType
	Side         Side
	Price        ujson.Float64
	Size         ujson.Float64
	RemainSize   ujson.Float64
	FilledSize   ujson.Float64
	CanceledSize ujson.Float64
	TradeId      string
	ClientOid    string
	OrderTime    int64
	OldSize      ujson.Float64
	Liquidity    string
	Ts           ujson.Int64
}

// Account Balance Events
// https://docs.kucoin.com/futures/#account-balance-events
type WalletShot struct {
	Currency         string
	Timestamp        ujson.Int64
	OrderMargin      ujson.Float64
	WithdrawHold     ujson.Float64
	AvailableBalance ujson.StringFloat64
	HoldBalance      ujson.StringFloat64
}
