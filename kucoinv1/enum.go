package kucoinv1

type Category string

const (
	Spot    Category = "SPOT"
	Futures Category = "FUTURES"
)

// Side represents order side
type Side string

const (
	SideBuy  Side = "buy"
	SideSell Side = "sell"
)

// OrderType represents order type
type OrderType string

const (
	OrderTypeLimit  OrderType = "limit"
	OrderTypeMarket OrderType = "market"
)

// TimeInForce represents time in force
type TimeInForce string

const (
	TimeInForceGTC TimeInForce = "GTC" // Good Till Canceled - deafault value
	TimeInForceIOC TimeInForce = "IOC" // Immediate Or Cancel
	TimeInForceFOK TimeInForce = "FOK" // Fill Or Kill
)

// MarginMode represents margin mode for positions
type MarginMode string

const (
	MarginModeIsolated MarginMode = "ISOLATED"
	MarginModeCross    MarginMode = "CROSS"
)

// StopPriceType represents stop order price type
type StopPriceType string

const (
	StopPriceTypeTrade StopPriceType = "TP" // Trade price
	StopPriceTypeIndex StopPriceType = "IP" // Index price
	StopPriceTypeMark  StopPriceType = "MP" // Mark price
)

// OrderStatus represents order status
/*
	1. The HTTP futures details only available:
	- open
	- done

	2. The HTTP spot details response does not include orderStatus

	3. The WS futures order details only available:
	- open
	- match
	- done

	4. The WS spot order details only available:
	- new
	- open
	- match
	- done
*/

type OrderStatus string

const (
	OrderStatusNew   OrderStatus = "new"
	OrderStatusOpen  OrderStatus = "open"
	OrderStatusDone  OrderStatus = "done"
	OrderStatusMatch OrderStatus = "match"
)

type OrderTypeWs string

const (
	OrderTypeWsOpen     OrderTypeWs = "open"
	OrderTypeWsMatch    OrderTypeWs = "match"
	OrderTypeWsUpdate   OrderTypeWs = "update"
	OrderTypeWsFilled   OrderTypeWs = "filled"
	OrderTypeWsCanceled OrderTypeWs = "canceled"
	OrderTypeWsReceived OrderTypeWs = "received" // only for Spot
)

// PositionSide represents position side
type PositionSide string

const (
	PositionSideLong  PositionSide = "LONG"
	PositionSideShort PositionSide = "SHORT"
	PositionSideBoth  PositionSide = "BOTH"
)

// Bar represents candlestick interval
type Bar int

const (
	Bar1m  Bar = 1
	Bar5m  Bar = 5
	Bar15m Bar = 15
	Bar30m Bar = 30
	Bar1H  Bar = 60
	Bar2H  Bar = 120
	Bar4H  Bar = 240
	Bar8H  Bar = 480
	Bar12H Bar = 720
	Bar1D  Bar = 1440
	Bar1W  Bar = 10080
)

// CandleInterval represents WebSocket candle subscription interval (Spot)
// https://www.kucoin.com/docs-new/3470071w0
type CandleInterval string

const (
	CandleInterval1m  CandleInterval = "1min"
	CandleInterval3m  CandleInterval = "3min"
	CandleInterval15m CandleInterval = "15min"
	CandleInterval30m CandleInterval = "30min"
	CandleInterval1H  CandleInterval = "1hour"
	CandleInterval2H  CandleInterval = "2hour"
	CandleInterval4H  CandleInterval = "4hour"
	CandleInterval6H  CandleInterval = "6hour"
	CandleInterval8H  CandleInterval = "8hour"
	CandleInterval12H CandleInterval = "12hour"
	CandleInterval1D  CandleInterval = "1day"
	CandleInterval1W  CandleInterval = "1week"
)
