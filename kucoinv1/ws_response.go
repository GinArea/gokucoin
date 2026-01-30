package kucoinv1

import (
	"github.com/msw-x/moon/ulog"
)

// WsResponse - WebSocket message response
type WsResponse struct {
	Id      string `json:"id"`
	Type    string `json:"type"`
	Topic   string `json:"topic"`
	Subject string `json:"subject"`
}

// IsTopic - returns true for data messages that should be routed to topic handler
func (o WsResponse) IsTopic() bool {
	return o.Type == "message"
}

// IsOperation - returns true for subscription operation responses
func (o WsResponse) IsOperation() bool {
	return o.Type == "subscribe" || o.Type == "unsubscribe"
}

// IsError - returns true for error messages
func (o WsResponse) IsError() bool {
	return o.Type == "error"
}

// IsWelcome - returns true for welcome message (connection confirmed by server)
func (o WsResponse) IsWelcome() bool {
	return o.Type == "welcome"
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
