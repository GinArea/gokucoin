package kucoinv1

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
	TimeInForceGTC TimeInForce = "GTC" // Good Till Canceled
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
type OrderStatus string

const (
	OrderStatusOpen      OrderStatus = "open"
	OrderStatusDone      OrderStatus = "done"
	OrderStatusMatch     OrderStatus = "match"
	OrderStatusCancelled OrderStatus = "cancelled"
)

// PositionSide represents position side
type PositionSide string

const (
	PositionSideLong  PositionSide = "long"
	PositionSideShort PositionSide = "short"
	PositionSideBoth  PositionSide = "both"
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
