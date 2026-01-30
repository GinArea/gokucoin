package kucoinv1

import (
	"github.com/msw-x/moon/ulog"
	"github.com/msw-x/moon/uws"
)

type WsPublic struct {
	WsBase
}

func NewWsPublic() *WsPublic {
	o := new(WsPublic)
	o.init(nil, false)
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

func (o *WsPublic) OrderBook(symbol string) *Executor[Orderbook] {
	return NewExecutor[Orderbook]("/contractMarket/level2Depth5:"+symbol, o.subscriptions)
}
