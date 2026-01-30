package kucoinv1

import (
	"github.com/msw-x/moon/ulog"
	"github.com/msw-x/moon/uws"
)

// WsBase - common WebSocket functionality for public and private clients
type WsBase struct {
	c              *WsClient
	ready          bool
	privateChannel bool
	onReady        func()
	onConnected    func()
	onDisconnected func()
	onDialError    func(error) bool
	subscriptions  *Subscriptions
}

func (o *WsBase) init(s *Sign, privateChannel bool, category Category) {
	o.c = NewWsClient(s).WithCategory(category)
	o.privateChannel = privateChannel
	o.subscriptions = NewSubscriptions(o)
}

func (o *WsBase) Close() {
	o.c.Close()
}

func (o *WsBase) Transport() *uws.Options {
	return o.c.Transport()
}

func (o *WsBase) Connected() bool {
	return o.c.Connected()
}

func (o *WsBase) Ready() bool {
	return o.ready
}

func (o *WsBase) Run() {
	o.c.WithOnConnected(func() {
		if o.onConnected != nil {
			o.onConnected()
		}
	})
	o.c.WithOnDisconnected(func() {
		o.ready = false
		if o.onDisconnected != nil {
			o.onDisconnected()
		}
	})
	o.c.WithOnDialError(o.handleDialError)
	o.c.WithOnResponse(o.onResponse)
	o.c.WithOnTopic(o.onTopic)
	o.c.Run()
}

func (o *WsBase) subscribe(topic string) {
	o.c.Subscribe(topic, o.privateChannel)
}

func (o *WsBase) unsubscribe(topic string) {
	o.c.Unsubscribe(topic, o.privateChannel)
}

func (o *WsBase) onResponse(r WsResponse) {
	if r.IsWelcome() {
		o.subscriptions.subscribeAll()
		o.ready = true
		if o.onReady != nil {
			o.onReady()
		}
	}
}

func (o *WsBase) onTopic(data []byte) error {
	return o.subscriptions.processTopic(data)
}

func (o *WsBase) handleDialError(err error) bool {
	o.ready = false
	// TODO: add KuCoin-specific error handling if needed
	// Example: check for 401/403 and call o.c.c.Cancel() to stop reconnecting
	if o.onDialError != nil {
		return o.onDialError(err)
	}
	// continue reconnect
	return false
}

// Setters for callbacks (used by With* methods in derived types)

func (o *WsBase) setLog(log *ulog.Log) {
	o.c.WithLog(log)
}

func (o *WsBase) setProxy(proxy string) {
	o.c.WithProxy(proxy)
}

func (o *WsBase) setLogRequest(enable bool) {
	o.c.WithLogRequest(enable)
}

func (o *WsBase) setLogResponse(enable bool) {
	o.c.WithLogResponse(enable)
}

func (o *WsBase) setOnDialError(f func(error) bool) {
	o.onDialError = f
}

func (o *WsBase) setOnConnected(f func()) {
	o.onConnected = f
}

func (o *WsBase) setOnDisconnected(f func()) {
	o.onDisconnected = f
}

func (o *WsBase) setOnReady(f func()) {
	o.onReady = f
}
