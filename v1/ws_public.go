package v1

import (
	"github.com/msw-x/moon/ulog"
	"github.com/msw-x/moon/uws"
)

type WsPublic struct {
	c             *WsClient
	onConnected   func()
	subscriptions *Subscriptions
}

func NewWsPublic() *WsPublic {
	o := new(WsPublic)
	o.c = NewWsClient()
	o.subscriptions = NewSubscriptions(o)
	return o
}

func (o *WsPublic) Close() {
	o.c.Close()
}

func (o *WsPublic) Transport() *uws.Options {
	return o.c.Transport()
}

func (o *WsPublic) Connected() bool {
	return o.c.Connected()
}

func (o *WsPublic) Ready() bool {
	return o.Connected()
}

func (o *WsPublic) WithLog(log *ulog.Log) *WsPublic {
	o.c.WithLog(log)
	return o
}

// func (o *WsPublic) WithBaseUrl(url string) *WsPublic {
// 	o.c.WithBaseUrl(url)
// 	return o
// }

func (o *WsPublic) WithProxy(proxy string) *WsPublic {
	o.c.WithProxy(proxy)
	return o
}

func (o *WsPublic) WithLogRequest(enable bool) *WsPublic {
	o.c.WithLogRequest(enable)
	return o
}

func (o *WsPublic) WithLogResponse(enable bool) *WsPublic {
	o.c.WithLogResponse(enable)
	return o
}

func (o *WsPublic) WithOnDialError(f func(error) bool) *WsPublic {
	o.c.WithOnDialError(f)
	return o
}

func (o *WsPublic) WithOnConnected(f func()) *WsPublic {
	o.onConnected = f
	return o
}

func (o *WsPublic) WithOnDisconnected(f func()) *WsPublic {
	o.c.WithOnDisconnected(f)
	return o
}

func (o *WsPublic) subscribe(topic string) {
	o.c.Subscribe(topic)
}

func (o *WsPublic) unsubscribe(topic string) {
	o.c.Unsubscribe(topic)
}

func (o *WsPublic) Run() {
	o.c.WithOnConnected(func() {
		if o.onConnected != nil {
			o.onConnected()
		}
		o.subscriptions.subscribeAll()
	})
	o.c.WithOnResponse(o.onResponse)
	o.c.WithOnTopic(o.onTopic)
	o.c.Run()
}

func (o *WsPublic) onResponse(r WsResponse) error {
	r.Log(o.c.Log())
	return nil
}

func (o *WsPublic) onTopic(data []byte) error {
	return o.subscriptions.processTopic(data)
}

func (o *WsPublic) OrderBook(symbol string) *Executor[Orderbook] {
	return NewExecutor[Orderbook](getTopicName(symbol), o.subscriptions)
}

func getTopicName(symbol string) string {
	return "/contractMarket/level2Depth50:" + symbol
}