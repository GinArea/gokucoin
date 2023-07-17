package v1

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
	case "ping":
	case "pong":
	case "subscribe":
		// if o.Success {
		// 	log.Info("subscribe: success")
		// } else {
		// 	log.Error("subscribe:", o.Message)
		// }
	case "unsubscribe":
		// log.Info("unsubscribe:", ufmt.SuccessFailure(o.Success))
	case "error":
		// log.Error("error", o.Id)
	default:
		log.Error("invalid response:", o.Type)
	}
}
