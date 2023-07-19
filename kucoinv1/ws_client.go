package kucoinv1

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
	"github.com/msw-x/moon/uhttp"
	"github.com/msw-x/moon/ulog"
	"github.com/msw-x/moon/uws"
)

type WsClient struct {
	c *uws.Client
	// onConnected    func()
	// onDisconnected func()
	onTopic func([]byte) error
}

func NewWsClient(sign *Sign) *WsClient {

	client := NewClient()

	var tokenInfo Response[WsTokenResponse]
	if sign.Key == "" || sign.Secret == "" || sign.Password == "" {
		tokenInfo = client.GetPublicWsToken()
	} else {
		client.WithAuth(sign.Key, sign.Secret, sign.Password)
		tokenInfo = client.GetPrivateWsToken()
	}

	if len(tokenInfo.Data) > 0 {
		token := tokenInfo.Data[0].Token
		endPoint := tokenInfo.Data[0].InstanceServers[0].Endpoint
		url := endPoint + "?token=" + token
		o := new(WsClient)
		o.c = uws.NewClient(url)
		return o
	}
	return nil
}

func (c *Client) GetPrivateWsToken() Response[WsTokenResponse] {
	return Post[WsTokenResponse](c, "bullet-private", nil, prepareTokeInfo)
}

func (c *Client) GetPublicWsToken() Response[WsTokenResponse] {
	return PostPub[WsTokenResponse](c, "bullet-public", nil, prepareTokeInfo)
}

func prepareTokeInfo(h uhttp.Responce) (r Response[WsTokenResponse], er error) {
	if h.BodyExists() {
		raw := new(item[WsTokenResponse])
		h.Json(raw)
		r.Time = getCurrentTime()
		r.Error = raw.Error()
		if r.Ok() {
			res := WsTokenResponse{
				Token:           raw.Data.Token,
				InstanceServers: raw.Data.InstanceServers,
			}
			r.Data = []WsTokenResponse{res}
		}
	}
	return
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

func (o *WsClient) ping() {
	o.Send(WsRequest{
		Id:   getRandomInt32(),
		Type: "ping",
	})
}

func (o *WsClient) onMessage(messageType int, data []byte) {
	log := o.c.Log()
	if messageType != websocket.TextMessage {
		log.Warning("invalid message type:", uws.MessageTypeString(messageType))
		return
	}
	var r WsResponse
	err := json.Unmarshal(data, &r)
	if err == nil {
		// fmt.Printf("Topic: %s; Type: %s \n", r.Topic, r.Type)
		// fmt.Printf("Full msg: %s \n", r)
		if o.onTopic != nil {
			err = o.onTopic(data)
		}
	}
	if err != nil {
		log.Error(err)
	}
}

func getRandomInt32() int32 {
	// Initialize the random number generator with a unique seed based on the current time
	rand.Seed(time.Now().UnixNano())
	return rand.Int31() // Generate a random int32 value
}
