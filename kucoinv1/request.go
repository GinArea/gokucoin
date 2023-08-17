package kucoinv1

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/msw-x/moon/ufmt"
	"github.com/msw-x/moon/uhttp"
)

func req[T any](c *Client, method string, path string, request any, sign bool, transform func(uhttp.Responce) (Response[T], error)) (r Response[T]) {
	var perf *uhttp.Performer
	switch method {
	case http.MethodGet:
		perf = c.c.Get(path).Params(request)
	case http.MethodPost:
		perf = c.c.Post(path).Json(request)
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
		if h.StatusCode == http.StatusOK || h.StatusCode == http.StatusTooManyRequests || h.StatusCode == http.StatusInternalServerError {
			r, _ = transform(h)
		} else {
			r.Error = errors.New(ufmt.Join(h.Status, h.Text()))
		}
		if sign {
			r.SetErrorIfNil(h.HeaderTo(&r.Limit))
		}
		r.StatusCode = h.StatusCode
	} else {
		r.Error = h.Error
	}
	return
}

func request[T any](c *Client, method string, path string, request any, sign bool, transform func(uhttp.Responce) (Response[T], error)) (r Response[T]) {
	var attempt int
	for {
		r = req(c, method, path, request, sign, transform)
		if r.StatusCode != http.StatusOK && c.onTransportError != nil {
			if c.onTransportError(r.Error, r.StatusCode, attempt) {
				attempt++
				continue
			}
		}
		break
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

func PostPub[T any](c *Client, path string, req any, transform func(uhttp.Responce) (Response[T], error)) Response[T] {
	return request[T](c, http.MethodPost, path, req, false, transform)
}
