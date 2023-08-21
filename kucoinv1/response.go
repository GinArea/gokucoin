package kucoinv1

import "github.com/msw-x/moon/ujson"

type Response[T any] struct {
	Data       T
	Limit      RateLimit
	Error      error
	StatusCode int
}

func (o *Response[T]) Ok() bool {
	return o.Error == nil
}

func (o *Response[T]) SetErrorIfNil(err error) {
	if o.Error == nil {
		o.Error = err
	}
}

type response[T any] struct {
	Code ujson.Int64
	Data T
	Msg  string
}

func (o *response[T]) Error() error {
	e := Error{
		Code: o.Code,
		Text: o.Msg,
	}
	return e.Std()
}
