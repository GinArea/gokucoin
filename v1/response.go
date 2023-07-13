package v1

type Response[T any] struct {
	Time  uint64
	Data  []T
	Limit RateLimit
	Error error
}

type response[T any] struct {
	Code string
	Data []T
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
	return getError(o.Code, "")
}

func (o *item[T]) Error() error {
	return getError(o.Code, o.Msg)
}

func getError(code string, msg string) error {
	e := Error{
		Code: code,
		Text: msg,
	}
	return e.Std()
}
