package kucoinv1

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Sign struct {
	Key      string
	Secret   string
	Password string
}

func NewSign(key, secret, password string) *Sign {
	return &Sign{Key: key, Secret: secret, Password: password}
}

// encryptPassphrase encrypts passphrase for API version 2
func (o *Sign) encryptPassphrase() string {
	return signHmac(o.Password, o.Secret)
}

func (o *Sign) timestamp() string {
	return strconv.FormatInt(time.Now().UnixMilli(), 10)
}

func (o *Sign) HeaderGet(h http.Header, v url.Values, path string) {
	o.header(h, v.Encode(), path, "GET")
}

func (o *Sign) HeaderPost(h http.Header, body []byte, path string) {
	o.header(h, string(body), path, "POST")
}

func (o *Sign) header(h http.Header, data string, path string, method string) {
	ts := o.timestamp()

	// Pre-sign: timestamp + method + /api/v1/path + body
	preSign := ts + method + "/" + ApiVersion + "/" + path
	if data != "" {
		if method == "GET" {
			preSign += "?" + data
		} else {
			preSign += data
		}
	}

	h.Set("KC-API-KEY", o.Key)
	h.Set("KC-API-SIGN", signHmac(preSign, o.Secret))
	h.Set("KC-API-TIMESTAMP", ts)
	h.Set("KC-API-PASSPHRASE", o.encryptPassphrase())
	h.Set("KC-API-KEY-VERSION", "2")

	// Broker headers
	partnerSign := signHmac(ts+BrokerPartner+o.Key, BrokerKey)
	h.Set("KC-API-PARTNER", BrokerPartner)
	h.Set("KC-API-PARTNER-SIGN", partnerSign)
	h.Set("KC-BROKER-NAME", BrokerName)
	h.Set("KC-API-PARTNER-VERIFY", "true")
}

// WebSocket generates signature for WS authentication
// Used when requesting /api/v1/bullet-private token
func (o *Sign) WebSocket(ts string) string {
	preSign := ts + "GET" + "/" + ApiVersion + "/bullet-private"
	return signHmac(preSign, o.Secret)
}

func signHmac(message, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	io.WriteString(h, message)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
