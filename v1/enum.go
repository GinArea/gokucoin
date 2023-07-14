package v1

// contractStatuc (https://docs.kucoin.com/futures/#get-open-contract-list)
type ContractStatus string

const (
	Init         ContractStatus = "Init"
	Open         ContractStatus = "Open"
	BeingSettled ContractStatus = "BeingSettled"
	Settled      ContractStatus = "Settled"
	Paused       ContractStatus = "Paused"
	Closed       ContractStatus = "Closed"
	CancelOnly   ContractStatus = "CancelOnly"
)

type Side string

const (
	Buy  Side = "buy"
	Sell Side = "sell"
)

type OrderType string

const (
	Limit      OrderType = "limit"
	Market     OrderType = "market"
	LimitStop  OrderType = "limit_stop"
	MarketStop OrderType = "market_stop"
)

type StopType string

const (
	Up   StopType = "up"
	Down StopType = "down"
)

type Status string

const (
	OpenStatus Status = "open"
	DoneStatus Status = "done"
)

type StopPriceType string

const (
	Tp StopPriceType = "TP"
	Ip StopPriceType = "IP"
	Mp StopPriceType = "MP"
)

type TimeInForce string

const (
	GoodTillCancel    TimeInForce = "GTC"
	ImmediateOrCancel TimeInForce = "IOC"
)
