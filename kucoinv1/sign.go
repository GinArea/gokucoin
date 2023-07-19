package kucoinv1

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"
)

const (
	partner  = "Ginareafutures"
	kcApiKey = "69b6522d-82a3-44e8-bbe5-32bb7658ba79"
)

type Sign struct {
	Key      string
	Secret   string
	Password string
}

func NewSign(key, secret, password string) *Sign {
	o := new(Sign)
	o.Key = key
	o.Secret = secret
	o.Password = password
	return o
}

func (o *Sign) HeaderGet(h http.Header, v url.Values, path string) {

	encodedParams := encodeSortParams(v)
	o.header(h, encodedParams, false, path, "GET")
}

func (o *Sign) HeaderPost(h http.Header, body []byte, path string) {
	o.header(h, string(body[:]), true, path, "POST")
}

func (o *Sign) header(h http.Header, s string, needPartnerHeader bool, path string, method string) {

	//необходимо получить preSignString в формате: timestamp + GET + url
	//например:
	// 1689235726523GET/api/v1/account-overview?currency=USDT

	ts := o.timestamp()
	preSignString := ts + method + "/" + ApiVersion + "/" + path
	if s != "" {
		var delimeter string
		if method == "GET" {
			delimeter = "?"
		}
		preSignString = preSignString + delimeter + s
	}
	fmt.Println(preSignString)

	kcApiSign := signHmac(preSignString, o.Secret)

	if needPartnerHeader {
		h.Set("KC-API-PARTNER", partner)
		h.Set("KC-API-PARTNER-SIGN", signHmac(ts+partner+o.Key, kcApiKey))
	}

	h.Set("KC-API-KEY", o.Key)
	h.Set("KC-API-SIGN", kcApiSign)
	h.Set("KC-API-TIMESTAMP", ts)
	h.Set("KC-API-PASSPHRASE", signHmac(o.Password, o.Secret))
	h.Set("KC-API-KEY-VERSION", "2")
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

func signHmac(s, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	io.WriteString(h, s)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
