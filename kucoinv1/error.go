package kucoinv1

import (
	"fmt"

	"golang.org/x/exp/slices"
)

// Error
type Error struct {
	Code string
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
	return o.Code == "200000"
}

func (o *Error) Error() string {
	return fmt.Sprintf("code[%s]: %s", o.Code, o.Text)
}

func (o *Error) SymbolIsNotWhitelisted() bool {
	//return o.Code == 10029
	return false
}

func (o *Error) RequestParameterError() bool {
	//return o.Code == 10001
	return false
}

func (o *Error) ApiKeyInvalid() bool {
	codes := []string{
		// 10003, // API key is invalid
		// 10004, // Error sign, please check your signature generation algorithm.
		// 10005, // Permission denied, please check your API key permissions
	}
	return slices.Contains(codes, o.Code)
}

func (o *Error) ApiKeyExpired() bool {
	// return o.Code == 33004
	return false
}

func (o *Error) TooManyVisits() bool {
	// return o.Code == 10006
	return false
}

func (o *Error) UnmatchedIp() bool {
	// return o.Code == 10010
	return false
}

func (o *Error) KycNeeded() bool {
	codes := []string{
		// 10024,  // Compliance rules triggered
		// 131004, // KYC needed
	}
	return slices.Contains(codes, o.Code)
}

func (o *Error) Timeout() bool {
	codes := []string{
		// 10000,  // Server Timeout
		// 170007, // Timeout waiting for response from backend server.
		// 170146, // Order creation timeout
		// 170147, // Order cancellation timeout
		// 177002, // Timeout
	}
	return slices.Contains(codes, o.Code)
}

func (o *Error) InsufficientBalance() bool {
	codes := []string{
		// 110004, // Wallet balance is insufficient
	}
	return slices.Contains(codes, o.Code)
}
