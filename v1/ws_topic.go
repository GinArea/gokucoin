package v1

import "encoding/json"

type RawTopic Topic[json.RawMessage]

type Topic[T any] struct {
	Type    string
	Topic   string
	Subject string
	Data    T
	// PublicHeader
	// PrivateHeader
}

//TODO после интеграции с приватными топиками рассмотреть деление

// type PublicHeader struct {
// 	Type string
// }

// type PrivateHeader struct {
// 	CreationTime int64
// 	Id           string
// }

func UnmarshalRawTopic[T any](raw RawTopic) (ret Topic[T], err error) {
	ret.Type = raw.Type
	ret.Topic = raw.Topic
	ret.Subject = raw.Subject
	// ret.PublicHeader = raw.PublicHeader
	// ret.PrivateHeader = raw.PrivateHeader
	err = json.Unmarshal(raw.Data, &ret.Data)
	return
}
