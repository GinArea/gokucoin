package v1

type Response[T any] struct {
	Time     uint64
	Data     []T
	Limit    RateLimit
	Error    error
	NetError bool
}

type nestedResponse[T any] struct {
	Code string
	Data struct {
		CurrentPage int64
		PageSize    int64
		Items       []T
	}
	Msg string
}

type response[T any] struct {
	Code string
	Data []T
	Msg  string
}

type item[T any] struct {
	Code string
	Msg  string
	Data T
}

func (o *Response[T]) Ok() bool {
	return o.Error == nil
}

func (o *Response[T]) SetErrorIfNil(err error) {
	if o.Error == nil {
		o.Error = err
	}
}

func (o *response[T]) Error() error {
	return getError(o.Code, o.Msg)
}

func (o *item[T]) Error() error {
	return getError(o.Code, o.Msg)
}

func (o *nestedResponse[T]) Error() error {
	return getError(o.Code, o.Msg)
}

func getError(code string, msg string) error {
	e := Error{
		Code: code,
		Text: msg,
	}
	return e.Std()
}

type WsTokenResponse struct {
	Token           string
	InstanceServers []InstanceServers
}

type InstanceServers struct {
	Endpoint     string
	Encrypt      bool
	Protocol     string
	PingInterval int
	PingTimeout  int
}
