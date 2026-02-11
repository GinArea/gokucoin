package kucoinv1

import (
	"fmt"
	"strings"

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
	}
	return slices.Contains(codes, o.Code)
}

func (o *Error) TooManyVisits() bool {
	codes := []string{
		"429000", // Too Many Requests
	}
	return slices.Contains(codes, o.Code)
}

func (o *Error) UnmatchedIp() bool {
	return o.Code == "400006" // 	The requested ip address is not on the api whitelist
}

func (o *Error) InsufficientBalance() bool {
	codes := []string{
		"102421", // spot: Insufficient account balance
		"200004", // spot: Balance insufficient!
		"200005", // Insufficient balance (insufficient balance when modifying risk limit)
		"300003", // Balance not enough, please first deposit at least 2 USDT before you go to battle
		"330008", // Order quantity is too high, insufficient available margin.
		"400100", // account.available.amount -- Insufficient balance
	}
	return slices.Contains(codes, o.Code)
}

func (o *Error) Timeout() bool {
	codes := []string{
		"100002", // System Error
		"500000", // Internal Server Error
		"1015",   // cloudflare frequency limit according to IP, block 30s
	}
	return slices.Contains(codes, o.Code)
}

func (o *Error) SymbolNotExists() bool {
	codes := []string{
		"600202", // Request failed because the symbol ... was not available
		"600203", // Symbol XXX-XXX cant be traded -- The symbol is not enabled for trading, such as downtime for upgrades, etc.
		"900001", // spot: symbol does not exist
	}
	return slices.Contains(codes, o.Code)

}

func (o *Error) ReduceOnlyError() bool {
	codes := []string{
		"300014", // The position is being liquidated, unable to place/cancel the order. Please try again later.
		"300009", // No open positions to close
	}
	return slices.Contains(codes, o.Code)
}

func (o *Error) InvalidOrderQuantity() bool {
	codes := []string{
		"400760", // spot SELL: The order funds should more than ... USDT
		"600100", // spot: Order size increment invalid
		"600101", // spot BUY: The order funds should more than ... USDT
	}
	return slices.Contains(codes, o.Code)
}

func (o *Error) UnknownCurrency() bool {
	return o.Code == "111001" // Currency does not exist
}

func (o *Error) ClientOIdIsDuplicate() bool {
	// Classic account mode
	// Spot + Futures allows duplicating ClientOId
	return false
}

func (o *Error) IncorrectTradingMode() bool {
	codes := []string{
		"330012", // Your order is in One-Way Mode, while your account is set to Hedge Mode. Update your settings so they match and try again
		"330013", // Your order is in Hedge Mode, while your account is set to One-Way Mode. Update your settings so they match and try again
	}
	return slices.Contains(codes, o.Code)
}

func (o *Error) AccessDenied() bool {
	codes := []string{
		"400007", // Access Denied
		"40010",  // Unavailable to place orders. Your identity information/IP/phone number shows you're in a country/region that is restricted from this service.
		"411100", // User is frozen -- Please contact us via support center
	}
	return slices.Contains(codes, o.Code)
}

func (o *Error) OrderDoesNotExist() bool {
	// 100001 - common error, need text analyze
	return strings.Contains(o.Text, "orderNotExist") && o.Code == "100001"
}
