package v1

import (
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"
)

type Sign struct {
	key      string
	secret   string
	password string
}

func NewSign(key, secret, password string) *Sign {
	o := new(Sign)
	o.key = key
	o.secret = secret
	return o
}

func (o *Sign) HeaderGet(h http.Header, v url.Values) {
	o.header(h, encodeSortParams(v))
}

func (o *Sign) HeaderPost(h http.Header, body []byte) {
	//TODO
	o.header(h, string(body[:]))
}

func (o *Sign) header(h http.Header, s string) {
	ts := o.timestamp()
	// h.Set("X-BAPI-API-KEY", o.key)
	h.Set("X-BAPI-TIMESTAMP", ts)
	//h.Set("X-BAPI-SIGN", signHmac(ts+o.key+window+s, o.secret))
}

func (o *Sign) timestamp() string {
	i := nowUtcMs()
	return strconv.Itoa(i)
}

func nowUtcMs() int {
	return int(time.Now().UTC().UnixNano() / int64(time.Millisecond))
}

func encodeSortParams(src url.Values) (s string) {
	if len(src) == 0 {
		return
	}
	keys := make([]string, len(src))
	i := 0
	for k := range src {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for _, k := range keys {
		s += encodeParam(k, src.Get(k)) + "&"
	}
	s = s[0 : len(s)-1]
	return
}

func encodeParam(name, value string) string {
	params := url.Values{}
	params.Add(name, value)
	return params.Encode()
}
