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

func (o *Response[T]) Ok() bool {
	return o.Error == nil
}

func (o *Response[T]) SetErrorIfNil(err error) {
	if o.Error == nil {
		o.Error = err
	}
}

func (o *response[T]) Error() error {
	e := Error{
		Code: o.Code,
		Text: "", //TODO
	}
	return e.Std()
}
