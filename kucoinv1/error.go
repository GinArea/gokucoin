package kucoinv1

import (
	"fmt"

	"golang.org/x/exp/slices"
)

type Error struct {
	Code string
	Text string
}

func (o *Error) Std() error {
	if o.Empty() {
		return nil
	}
	return o
}

func (o *Error) Empty() bool {
	return o.Code == "200000"
}

func (o *Error) Error() string {
	return fmt.Sprintf("code[%s]: %s", o.Code, o.Text)
}

func (o *Error) ApiKeyInvalid() bool {
	codes := []string{
		"400001", // Any of KC-API-KEY, KC-API-SIGN, KC-API-TIMESTAMP, KC-API-PASSPHRASE is missing
		"400002", // KC-API-TIMESTAMP Invalid
		"400003", // KC-API-KEY not exists
		"400004", // KC-API-PASSPHRASE error
		"400005", // Signature Error
		"400006", // The requested ip address is not in the api whitelist
		"400007", // Access Denied
	}
	return slices.Contains(codes, o.Code)
}

func (o *Error) TooManyVisits() bool {
	codes := []string{
		"100001", // The request is too frequent
		"429000", // Too Many Requests
	}
	return slices.Contains(codes, o.Code)
}

func (o *Error) UnmatchedIp() bool {
	return o.Code == "400006"
}

func (o *Error) InsufficientBalance() bool {
	codes := []string{
		"300000", // Balance insufficient
		"300003", // Insufficient balance
	}
	return slices.Contains(codes, o.Code)
}

func (o *Error) Timeout() bool {
	codes := []string{
		"100002", // System Error
		"500000", // Internal Server Error
	}
	return slices.Contains(codes, o.Code)
}

func (o *Error) OrderNotExist() bool {
	return o.Code == "400100" // Order not exists
}

func (o *Error) SymbolNotExists() bool {
	return o.Code == "100004" // Parameter error (used for invalid symbol)
}

func (o *Error) TradingDisabled() bool {
	codes := []string{
		"200002", // Trading is off
		"200003", // Limit order is disabled
		"200004", // Market order is disabled
	}
	return slices.Contains(codes, o.Code)
}

func (o *Error) OrderFrozen() bool {
	return o.Code == "411100" // Order frozen
}

func (o *Error) ReduceOnlyError() bool {
	return o.Code == "300014" // Reduce only order can not increase position
}

func (o *Error) InvalidOrderPrice() bool {
	return o.Code == "300011" // Unsupported order price
}

func (o *Error) InvalidOrderQuantity() bool {
	codes := []string{
		"300001", // Order quantity less than the minimum requirement
		"300002", // Order quantity greater than the maximum limit
		"300010", // Unsupported order quantity
	}
	return slices.Contains(codes, o.Code)
}

func (o *Error) InvalidLeverage() bool {
	return o.Code == "300012" // Unsupported leverage
}

func (o *Error) SignatureError() bool {
	return o.Code == "100005" // Signature error
}
