package kucoinv1

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
	"github.com/msw-x/moon/ulog"
	"github.com/msw-x/moon/uws"
	"golang.org/x/exp/slices"
)

type WsClient struct {
	c       *uws.Client
	s       *Sign
	onTopic func([]byte) error
}

func NewWsClient(s *Sign) *WsClient {
	o := new(WsClient)
	o.s = s
	o.c = uws.NewClient("ws")
	o.c.WithOnPreDial(o.getUrl)
	return o
}

func (o *WsClient) Close() {
	o.c.Close()
}

func (o *WsClient) Log() *ulog.Log {
	return o.c.Log()
}

func (o *WsClient) Transport() *uws.Options {
	return &o.c.Options
}

func (o *WsClient) WithLog(log *ulog.Log) *WsClient {
	o.c.WithLog(log)
	return o
}

func (o *WsClient) WithProxy(proxy string) *WsClient {
	o.c.WithProxy(proxy)
	return o
}

func (o *WsClient) WithLogRequest(enable bool) *WsClient {
	o.Transport().LogSent.Size = enable
	o.Transport().LogSent.Data = enable
	return o
}

func (o *WsClient) WithLogResponse(enable bool) *WsClient {
	o.Transport().LogRecv.Size = enable
	o.Transport().LogRecv.Data = enable
	return o
}

func (o *WsClient) WithOnDialError(f func(error) bool) *WsClient {
	o.c.WithOnDialError(f)
	return o
}

func (o *WsClient) WithOnConnected(f func()) *WsClient {
	o.c.WithOnConnected(f)
	return o
}

func (o *WsClient) WithOnDisconnected(f func()) *WsClient {
	o.c.WithOnDisconnected(f)
	return o
}

func (o *WsClient) WithOnTopic(f func([]byte) error) *WsClient {
	o.onTopic = f
	return o
}

func (o *WsClient) Run() {
	o.c.WithOnPing(o.ping)
	o.c.WithOnMessage(o.onMessage)
	o.c.Run()
}

func (o *WsClient) Connected() bool {
	return o.c.Connected()
}

func (o *WsClient) Send(r WsRequest) {
	o.c.SendJson(r)
}

func (o *WsClient) Subscribe(s string, isPrivateChannel bool) {
	o.Send(WsRequest{
		Id:             getRandomInt32(),
		Type:           "subscribe",
		Topic:          s,
		PrivateChannel: isPrivateChannel,
		Response:       false,
	})
}

func (o *WsClient) Unsubscribe(s string, isPrivateChannel bool) {
	o.Send(WsRequest{
		Id:             getRandomInt32(),
		Type:           "unsubscribe",
		Topic:          s,
		PrivateChannel: isPrivateChannel,
		Response:       false,
	})
}

func (o *WsClient) getUrl(string) string {
	client := NewClient()
	var r Response[WsToken]
	if o.s == nil {
		r = client.GetPublicWsToken()
	} else {
		client.WithAuth(o.s.Key, o.s.Secret, o.s.Password)
		r = client.GetPrivateWsToken()
	}
	if r.Ok() {
		if len(r.Data.InstanceServers) > 0 {
			token := r.Data.Token
			endPoint := r.Data.InstanceServers[0].Endpoint
			return endPoint + "?token=" + token
		} else {
			o.Log().Error("instance servers is empty")
		}
	} else {
		o.Log().Error("token is missing:", r.Error)
	}
	return ""
}

func (o *WsClient) ping() {
	o.Send(WsRequest{
		Id:   getRandomInt32(),
		Type: "ping",
	})
}

func (o *WsClient) onMessage(messageType int, data []byte) {
	if messageType != websocket.TextMessage {
		o.c.Log().Warning("invalid message type:", uws.MessageTypeString(messageType))
		return
	}
	// fmt.Printf("%s\n", data)
	var r WsResponse
	err := json.Unmarshal(data, &r)
	if err == nil {
		skipTypes := []string{"welcome", "pong"}
		if o.onTopic != nil && !slices.Contains(skipTypes, r.Type) {
			err = o.onTopic(data)
		}
	}
	if err != nil {
		o.c.Log().Error(err)
	}
}

// one time
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func getRandomInt32() int32 {
	return rnd.Int31() // Generate a random int32 value
}
