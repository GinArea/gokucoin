package kucoinv1

import (
	"github.com/msw-x/moon/ulog"
	"github.com/msw-x/moon/uws"
)

type WsPrivate struct {
	WsBase
	category Category
}

func NewWsPrivate(key, secret, password string, category Category) *WsPrivate {
	o := new(WsPrivate)
	o.category = category
	o.init(NewSign(key, secret, password), true, category)
	return o
}

// Builder methods (return *WsPrivate for chaining)

func (o *WsPrivate) WithLog(log *ulog.Log) *WsPrivate {
	o.setLog(log)
	return o
}

func (o *WsPrivate) WithProxy(proxy string) *WsPrivate {
	o.setProxy(proxy)
	return o
}

func (o *WsPrivate) WithLogRequest(enable bool) *WsPrivate {
	o.setLogRequest(enable)
	return o
}

func (o *WsPrivate) WithLogResponse(enable bool) *WsPrivate {
	o.setLogResponse(enable)
	return o
}

func (o *WsPrivate) WithOnDialError(f func(error) bool) *WsPrivate {
	o.setOnDialError(f)
	return o
}

func (o *WsPrivate) WithOnConnected(f func()) *WsPrivate {
	o.setOnConnected(f)
	return o
}

func (o *WsPrivate) WithOnDisconnected(f func()) *WsPrivate {
	o.setOnDisconnected(f)
	return o
}

func (o *WsPrivate) WithOnReady(f func()) *WsPrivate {
	o.setOnReady(f)
	return o
}

// Transport returns WebSocket transport options (shadows WsBase for type consistency)
func (o *WsPrivate) Transport() *uws.Options {
	return o.WsBase.Transport()
}

// Topic subscriptions

func (o *WsPrivate) Positions() *Executor[PositionShot] {
	return NewExecutor[PositionShot]("/contract/positionAll", o.subscriptions)
}

func (o *WsPrivate) OrdersFutures() *Executor[OrderShot] {
	return NewExecutor[OrderShot]("/contractMarket/tradeOrders", o.subscriptions)
}

func (o *WsPrivate) OrdersSpot() *Executor[OrderShotSpot] {
	return NewExecutor[OrderShotSpot]("/spotMarket/tradeOrdersV2", o.subscriptions)
}

func (o *WsPrivate) WalletFutures() *Executor[WalletShot] {
	return NewExecutor[WalletShot]("/contractAccount/wallet", o.subscriptions)
}

func (o *WsPrivate) WalletSpot() *Executor[WalletShotSpot] {
	return NewExecutor[WalletShotSpot]("/account/balance", o.subscriptions)
}
