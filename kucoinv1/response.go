package kucoinv1

type Response[T any] struct {
	Data       T
	Limit      RateLimit
	Error      error
	StatusCode int
	NetError   bool
}

type response[T any] struct {
	Code string
	Data T
	Msg  string
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
		Text: o.Msg,
	}
	return e.Std()
}

// Paginated - generic paginated response wrapper
type Paginated[T any] struct {
	CurrentPage int `json:"currentPage"`
	PageSize    int `json:"pageSize"`
	TotalNum    int `json:"totalNum"`
	TotalPage   int `json:"totalPage"`
	Items       []T `json:"items"`
}
