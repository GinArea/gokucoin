package v1

import (
	"github.com/msw-x/moon/uhttp"
	"github.com/msw-x/moon/ujson"
)

// Place an Order
// https://docs.kucoin.com/futures/#place-an-order

type PlaceOrder struct {
	ClientOid     string
	Side          Side
	Symbol        string
	Type          OrderType
	Leverage      float64
	Remark        string
	Stop          StopType
	StopPriceType StopPriceType
	StopPrice     float64
	ReduceOnly    bool
	CloseOrder    bool
	ForceHold     bool
	Price         float64
	Size          float64
	TimeInForce   TimeInForce
	PostOnly      bool
	Hidden        bool
	Iceberg       bool
	VisibleSize   float64
}

type OrderId struct {
	OrderId string
}

func (o *Client) PlaceOrder(v PlaceOrder) Response[OrderId] {
	return v.Do(o)
}

func (o PlaceOrder) Do(c *Client) Response[OrderId] {
	return placeOrder[OrderId](o, c)
}

func placeOrder[T any](o PlaceOrder, c *Client) Response[T] {
	return Post[T](c, "orders", o, func(h uhttp.Responce) (r Response[T], er error) {
		if h.BodyExists() {
			raw := new(item[T])
			h.Json(raw)
			r.Time = getCurrentTime()
			r.Error = raw.Error()
			if r.Ok() {
				r.Data = []T{raw.Data}
			}
		}
		return
	})
}

// Get Fills (проверка статуса ордера)
// https://docs.kucoin.com/futures/#get-fills

type GetFills struct {
	OrderId     string    `url:",omitempty"`
	Symbol      string    `url:",omitempty"`
	Side        Side      `url:",omitempty"`
	Type        OrderType `url:",omitempty"`
	StartAt     int64     `url:",omitempty"`
	EndAt       int64     `url:",omitempty"`
	CurrentPage int64     `url:",omitempty"`
	PageSize    int64     `url:",omitempty"`
}

type Fills struct {
	Symbol         string
	TradeId        string
	OrderId        string
	Side           Side
	Liquidity      string
	ForceTaker     bool
	Price          ujson.Float64
	Size           ujson.Float64
	Value          ujson.Float64
	FeeRate        ujson.Float64
	FixFee         ujson.Float64
	FeeCurrency    string
	Stop           string
	Fee            ujson.Float64
	OrderType      OrderType
	TradeType      string
	CreatedAt      int64
	SettleCurrency string
	TradeTime      int64
	// в документации есть, но не передаются поля
	// OpenFeePay     ujson.Float64
	// CloseFeePay    ujson.Float64
}

func (o *Client) GetFills(v GetFills) Response[Fills] {
	return v.Do(o)
}

func (o GetFills) Do(c *Client) Response[Fills] {
	return getFills[Fills](o, c)
}

func getFills[T any](o GetFills, c *Client) Response[T] {
	return Get[T](c, "fills", o, func(h uhttp.Responce) (r Response[T], er error) {
		if h.BodyExists() {
			raw := new(nestedResponse[T])
			h.Json(raw)
			r.Time = getCurrentTime()
			r.Error = raw.Error()
			if r.Ok() {
				r.Data = raw.Data.Items
			}
		}
		return
	})
}
