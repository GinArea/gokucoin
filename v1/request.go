package v1

import (
	"fmt"
	"net/http"

	"github.com/msw-x/moon/uhttp"
)

func request[T any](c *Client, method string, path string, request any, sign bool, transform func(uhttp.Responce) (Response[T], error)) (r Response[T]) {
	var perf *uhttp.Performer
	switch method {
	case http.MethodGet:
		perf = c.c.Get(path).Params(request)
	case http.MethodPost:
		perf = c.c.Post(path).Json(request)
		fmt.Println(string(perf.Request.Body))
	default:
		r.Error = fmt.Errorf("forbidden method: %s", method)
		return
	}
	if sign && c.s != nil {
		if perf.Request.Header == nil {
			perf.Request.Header = make(http.Header)
		}
		switch method {
		case http.MethodGet:
			c.s.HeaderGet(perf.Request.Header, perf.Request.Params, path)
		case http.MethodPost:
			c.s.HeaderPost(perf.Request.Header, perf.Request.Body, path)
		}
	}
	h := perf.Do()
	if h.Error == nil {
		r, _ = transform(h)
		if sign {
			r.SetErrorIfNil(h.HeaderTo(&r.Limit))
		}
	} else {
		r.Error = h.Error
	}
	return
}

func GetPub[T any](c *Client, path string, req any, transform func(uhttp.Responce) (Response[T], error)) Response[T] {
	return request[T](c, http.MethodGet, path, req, false, transform)
}

func Get[T any](c *Client, path string, req any, transform func(uhttp.Responce) (Response[T], error)) Response[T] {
	return request[T](c, http.MethodGet, path, req, true, transform)
}

func Post[T any](c *Client, path string, req any, transform func(uhttp.Responce) (Response[T], error)) Response[T] {
	return request[T](c, http.MethodPost, path, req, true, transform)
}
