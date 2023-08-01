package kucoinv1

import (
	"fmt"

	"github.com/msw-x/moon/ujson"
	"golang.org/x/exp/slices"
)

// Error
// Список ошибок - https://docs.kucoin.com/futures/#requests
type Error struct {
	Code ujson.Int64
	Text string
}

func (o *Error) Std() error {
	if o.Empty() {
		return nil
	} else {
		return o
	}
}

func (o *Error) Empty() bool {
	return o.Code == 200000
}

func (o *Error) Error() string {
	return fmt.Sprintf("code[%d]: %s", o.Code.Value(), o.Text)
}

func (o *Error) SymbolIsNotWhitelisted() bool {
	//return o.Code == 10029
	return false
}

func (o *Error) RequestParameterError() bool {
	codes := []int64{
		100001, // There are invalid parameters
		100003, // Contract parameter invalid
		300016, // The leverage cannot be greater than xxx
		300018, // clientOid parameter repeated
	}

	return slices.Contains(codes, o.Code.Value())
}

func (o *Error) ApiKeyInvalid() bool {
	codes := []int64{
		400003, // KC-API-KEY not exists
		400004, // Invalid KC-API-PASSPHRASE
		400007, // Access Denied -- Your API key does not have sufficient permissions to access the URI
		411100, // User is frozen -- Please contact us via support center
	}
	return slices.Contains(codes, o.Code.Value())
}

func (o *Error) ApiKeyExpired() bool {
	// не существует такой ошибки
	return false
}

func (o *Error) TooManyVisits() bool {
	codes := []int64{
		1015,   // 	cloudflare frequency limit according to IP, block 30s
		200002, // rate limit of each private endpoint of kucoin, based on user uid+endpoint mode limit, block10s
	}
	return slices.Contains(codes, o.Code.Value())
}

func (o *Error) UnmatchedIp() bool {
	return o.Code == 400006 // The IP address is not in the API whitelist
}

func (o *Error) KycNeeded() bool {
	// не существует такой ошибки
	return false
}

func (o *Error) Timeout() bool {
	codes := []int64{
		429000, // kucoin stand-alone capacity limit. It can be understood that the server is overloaded.
	}
	return slices.Contains(codes, o.Code.Value())
}

func (o *Error) InsufficientBalance() bool {
	codes := []int64{
		300003, // Balance insufficient
	}
	return slices.Contains(codes, o.Code.Value())
}
