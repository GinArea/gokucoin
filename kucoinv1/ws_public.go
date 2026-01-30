package kucoinv1

import (
	"github.com/msw-x/moon/ulog"
	"github.com/msw-x/moon/uws"
)

type WsPublic struct {
	WsBase
	category Category
}

func NewWsPublic(category Category) *WsPublic {
	o := new(WsPublic)
	o.category = category
	o.init(nil, false, category)
	return o
}

// Builder methods (return *WsPublic for chaining)

func (o *WsPublic) WithLog(log *ulog.Log) *WsPublic {
	o.setLog(log)
	return o
}

func (o *WsPublic) WithProxy(proxy string) *WsPublic {
	o.setProxy(proxy)
	return o
}

func (o *WsPublic) WithLogRequest(enable bool) *WsPublic {
	o.setLogRequest(enable)
	return o
}

func (o *WsPublic) WithLogResponse(enable bool) *WsPublic {
	o.setLogResponse(enable)
	return o
}

func (o *WsPublic) WithOnDialError(f func(error) bool) *WsPublic {
	o.setOnDialError(f)
	return o
}

func (o *WsPublic) WithOnConnected(f func()) *WsPublic {
	o.setOnConnected(f)
	return o
}

func (o *WsPublic) WithOnDisconnected(f func()) *WsPublic {
	o.setOnDisconnected(f)
	return o
}

func (o *WsPublic) WithOnReady(f func()) *WsPublic {
	o.setOnReady(f)
	return o
}

// Transport returns WebSocket transport options (shadows WsBase for type consistency)
func (o *WsPublic) Transport() *uws.Options {
	return o.WsBase.Transport()
}

// Topic subscriptions

// OrderBook subscribes to Level 5 order book updates for a symbol
// Futures topic: /contractMarket/level2Depth5:{symbol}
// https://www.kucoin.com/docs-new/3470083w0 (futures)
// Spot topic: /spotMarket/level2Depth5:{symbol}
// https://www.kucoin.com/docs-new/3470069w0 (Spot)
func (o *WsPublic) OrderBook(symbol string) *Executor[Orderbook] {
	return NewExecutor[Orderbook](o.orderbookTopic(symbol), o.subscriptions)
}

// orderbookTopic returns the appropriate topic based on category
func (o *WsPublic) orderbookTopic(symbol string) string {
	if o.category == Spot {
		return "/spotMarket/level2Depth5:" + symbol
	}
	return "/contractMarket/level2Depth5:" + symbol
}
