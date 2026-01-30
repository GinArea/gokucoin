package kucoinv1

import (
	"github.com/msw-x/moon/ulog"
)

type WsResponse struct {
	Id      string `json:"id"`
	Type    string `json:"type"`
	Topic   string `json:"topic"`
	Subject string `json:"subject"`
}

//  need to add (???): "welcome", "ack", "message"

func (o WsResponse) Log(log *ulog.Log) {
	switch o.Type {
	case "ping":
	case "pong":
	case "subscribe":
	case "unsubscribe":
	case "error":
	default:
		log.Error("invalid response:", o.Type)
	}
}
