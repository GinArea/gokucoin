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

func (o WsResponse) Log(log *ulog.Log) {
	switch o.Type {
	case "ping", "pong":
		// Silent
	case "welcome":
		log.Info("connected")
	case "message":
		// Data, handled separately
	case "subscribe", "unsubscribe":
		log.Debug("subscription:", o.Topic)
	case "error":
		log.Error("error:", o.Topic)
	default:
		log.Warning("unknown type:", o.Type)
	}
}
