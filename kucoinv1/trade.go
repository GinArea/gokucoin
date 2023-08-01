package kucoinv1

import (
	"github.com/msw-x/moon/uhttp"
	"github.com/msw-x/moon/ujson"
)

// Place an Order
// https://docs.kucoin.com/futures/#place-an-order

type PlaceOrder struct {
	ClientOid     string        `json:",omitempty"`
	Side          Side          `json:",omitempty"`
	Symbol        string        `json:",omitempty"`
	Type          OrderType     `json:",omitempty"`
	Leverage      ujson.Float64 `json:",omitempty"`
	Remark        string        `json:",omitempty"`
	Stop          StopType      `json:",omitempty"`
	StopPriceType StopPriceType `json:",omitempty"`
	StopPrice     ujson.Float64 `json:",omitempty"`
	ReduceOnly    bool          `json:",omitempty"`
	CloseOrder    bool          `json:",omitempty"`
	ForceHold     bool          `json:",omitempty"`
	Price         ujson.Float64 `json:",omitempty"`
	Size          ujson.Float64 `json:",omitempty"`
	TimeInForce   TimeInForce   `json:",omitempty"`
	PostOnly      bool          `json:",omitempty"`
	Hidden        bool          `json:",omitempty"`
	Iceberg       bool          `json:",omitempty"`
	VisibleSize   ujson.Float64 `json:",omitempty"`
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
			r.Error = raw.Error()
			if r.Ok() {
				r.Data = raw.Data.Items
			}
		}
		return
	})
}

// Get Details of a Single Order - нет комиссии
// https://docs.kucoin.com/futures/#get-details-of-a-single-order

type GetDetailsOfSingleOrder struct {
	OrderId   string `url:",omitempty"`
	ClientOid string `url:",omitempty"`
}

type DetailsOfSingleOrder struct {
	Id             string
	Symbol         string
	Type           OrderType
	Side           Side
	Price          ujson.Float64
	Size           ujson.Float64
	Value          ujson.Float64
	DealValue      ujson.Float64
	DealSize       ujson.Float64
	Stp            string
	Stop           string
	StopPriceType  string
	StopTriggered  bool
	StopPrice      ujson.Float64
	TimeInForce    TimeInForce
	PostOnly       bool
	Hidden         bool
	Iceberg        bool
	Leverage       ujson.Float64
	ForceHold      bool
	CloseOrder     bool
	VisibleSize    ujson.Float64
	ClientOid      string
	Remark         string
	Tags           string
	IsActive       bool
	CancelExist    bool
	CreatedAt      ujson.TimeMs
	UpdatedAt      ujson.TimeMs
	EndAt          ujson.TimeMs
	OrderTime      int64
	SettleCurrency string
	Status         Status
	FilledValue    ujson.Float64
	FilledSize     ujson.Float64
	ReduceOnly     bool
}

func (o *Client) GetDetailsOfSingleOrder(v GetDetailsOfSingleOrder) Response[DetailsOfSingleOrder] {

	return v.Do(o)
}

func (o GetDetailsOfSingleOrder) Do(c *Client) Response[DetailsOfSingleOrder] {
	return getDetailsOfSingleOrder[DetailsOfSingleOrder](o, c)
}

func getDetailsOfSingleOrder[T any](o GetDetailsOfSingleOrder, c *Client) (res Response[T]) {

	mapper := func(h uhttp.Responce) (r Response[T], err error) {
		if h.BodyExists() {
			raw := new(item[T])
			h.Json(raw)
			r.Error = raw.Error()
			if r.Ok() {
				r.Data = []T{raw.Data}
			}
		}
		return
	}
	if o.OrderId != "" {
		orderId := o.OrderId
		query := GetDetailsOfSingleOrder{}
		if o.ClientOid != "" {
			// если заполнен ClientOid, то также передаем его в качестве параметра, чтобы запрос был вида
			// GET /api/v1/orders/{order-id}?clientOid={client-order-id}
			query = GetDetailsOfSingleOrder{
				ClientOid: o.ClientOid,
			}
		}
		// иначе запрос будет
		// GET /api/v1/orders/{order-id}
		return Get[T](c, "orders/"+orderId, query, mapper)
	} else {
		if o.ClientOid != "" {
			//GET /api/v1/orders/byClientOid?clientOid={client-order-id}
			return Get[T](c, "orders/byClientOid", o, mapper)
		}
	}
	return res
}
