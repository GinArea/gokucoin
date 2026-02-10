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

// Bar represents candlestick interval (futures)
// HTTP requrest - `granularity` field
// https://www.kucoin.com/docs-new/rest/futures-trading/market-data/get-klines
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

// KlineInterval - spot candlestick interval type
// HTTP request - `type` field
// https://www.kucoin.com/docs-new/rest/spot-trading/market-data/get-klines
type KlineSpotInt string

const (
	KlineSpotInt1min   KlineSpotInt = "1min"
	KlineSpotInt3min   KlineSpotInt = "3min"
	KlineSpotInt5min   KlineSpotInt = "5min"
	KlineSpotInt15min  KlineSpotInt = "15min"
	KlineSpotInt30min  KlineSpotInt = "30min"
	KlineSpotInt1hour  KlineSpotInt = "1hour"
	KlineSpotInt2hour  KlineSpotInt = "2hour"
	KlineSpotInt4hour  KlineSpotInt = "4hour"
	KlineSpotInt6hour  KlineSpotInt = "6hour"
	KlineSpotInt8hour  KlineSpotInt = "8hour"
	KlineSpotInt12hour KlineSpotInt = "12hour"
	KlineSpotInt1day   KlineSpotInt = "1day"
	KlineSpotInt1week  KlineSpotInt = "1week"
)

// CandleInterval represents WebSocket candle subscription interval (Spot)
// https://www.kucoin.com/docs-new/3470071w0

// https://www.kucoin.com/docs-new/3470086w0 (Futures) + 1month -6hour
type CandleInterval string

const (
	CandleInterval1m  CandleInterval = "1min"
	CandleInterval3m  CandleInterval = "3min"
	CandleInterval5m  CandleInterval = "5min"
	CandleInterval15m CandleInterval = "15min"
	CandleInterval30m CandleInterval = "30min"
	CandleInterval1H  CandleInterval = "1hour"
	CandleInterval2H  CandleInterval = "2hour"
	CandleInterval4H  CandleInterval = "4hour"
	CandleInterval8H  CandleInterval = "8hour"
	CandleInterval12H CandleInterval = "12hour"
	CandleInterval1D  CandleInterval = "1day"
	CandleInterval1W  CandleInterval = "1week"
)
