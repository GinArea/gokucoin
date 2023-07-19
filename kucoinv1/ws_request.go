package kucoinv1

type WsRequest struct {
	Id             int32  `json:"id,omitempty"`
	Type           string `json:"type,omitempty"`
	Topic          string `json:"topic,omitempty"`
	PrivateChannel bool   `json:"privateChannel,omitempty"`
	Response       bool   `json:"response,omitempty"`
}
