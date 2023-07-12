package v1

import (
	"fmt"
	"testing"
)

func TestClient_GetOpenContracts(t *testing.T) {
	type args struct {
		v GetOpenContracts
	}
	tests := []struct {
		name   string
		client *Client
		args   args
	}{
		{
			name:   "GetOpenContractsTest",
			client: NewClient(),
			args: args{
				v: GetOpenContracts{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.client.GetOpenContracts(tt.args.v)
			fmt.Printf("%v", got)
		})
	}
}

func TestClient_GetTicker(t *testing.T) {
	type args struct {
		v GetTicker
	}
	tests := []struct {
		name   string
		client *Client
		args   args
	}{
		{
			name:   "GetTickerTest",
			client: NewClient(),
			args: args{
				v: GetTicker{
					Symbol: "XBTUSDM",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.client.GetTicker(tt.args.v)
			fmt.Printf("%v", got)
		})
	}
}
