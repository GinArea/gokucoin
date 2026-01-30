package kucoinv1

import (
	"github.com/msw-x/moon/ulog"
	"github.com/msw-x/moon/uws"
)

type WsPrivate struct {
	c              *WsClient
	onConnected    func()
	onDisconnected func()
	subscriptions  *Subscriptions
}

func NewWsPrivate(key, secret, password string) *WsPrivate {
	o := new(WsPrivate)
	o.c = NewWsClient(NewSign(key, secret, password))
	o.subscriptions = NewSubscriptions(o)
	return o
}

func (o *WsPrivate) Close() {
	o.c.Close()
}

func (o *WsPrivate) Transport() *uws.Options {
	return o.c.Transport()
}

func (o *WsPrivate) WithLog(log *ulog.Log) *WsPrivate {
	o.c.WithLog(log)
	return o
}

func (o *WsPrivate) WithProxy(proxy string) *WsPrivate {
	o.c.WithProxy(proxy)
	return o
}

func (o *WsPrivate) WithLogRequest(enable bool) *WsPrivate {
	o.c.WithLogRequest(enable)
	return o
}

func (o *WsPrivate) WithLogResponse(enable bool) *WsPrivate {
	o.c.WithLogResponse(enable)
	return o
}

func (o *WsPrivate) WithOnDialError(f func(error) bool) *WsPrivate {
	o.c.WithOnDialError(f)
	return o
}

func (o *WsPrivate) WithOnConnected(f func()) *WsPrivate {
	o.onConnected = f
	return o
}

func (o *WsPrivate) WithOnDisconnected(f func()) *WsPrivate {
	o.onDisconnected = f
	return o
}

func (o *WsPrivate) Run() {
	o.c.WithOnConnected(func() {
		if o.onConnected != nil {
			o.onConnected()
		}
		o.subscriptions.subscribeAll()
	})
	o.c.WithOnTopic(o.onTopic)
	o.c.Run()
}

func (o *WsPrivate) Connected() bool {
	return o.c.Connected()
}

func (o *WsPrivate) Ready() bool {
	return o.Connected()
}

func (o *WsPrivate) subscribe(topic string) {
	o.c.Subscribe(topic, true)
}

func (o *WsPrivate) unsubscribe(topic string) {
	o.c.Unsubscribe(topic, true)
}

func (o *WsPrivate) onTopic(data []byte) error {
	return o.subscriptions.processTopic(data)
}

func (o *WsPrivate) Positions() *Executor[PositionShot] {
	return NewExecutor[PositionShot]("/contract/positionAll", o.subscriptions)
}

func (o *WsPrivate) Order(symbol string) *Executor[OrderShot] {
	return NewExecutor[OrderShot]("/contractMarket/tradeOrders:"+symbol, o.subscriptions)
}

func (o *WsPrivate) Orders() *Executor[OrderShot] {
	return NewExecutor[OrderShot]("/contractMarket/tradeOrders", o.subscriptions)
}

func (o *WsPrivate) Wallet() *Executor[WalletShot] {
	return NewExecutor[WalletShot]("/contractAccount/wallet", o.subscriptions)
}
