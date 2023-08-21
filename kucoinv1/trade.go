package kucoinv1

import (
	"github.com/msw-x/moon/ujson"
)

// Place an Order
// https://docs.kucoin.com/futures/#place-an-order
type PlaceOrder struct {
	ClientOid     string
	Side          Side
	Symbol        string
	Type          OrderType     `json:",omitempty"`
	Leverage      ujson.Float64 `json:",omitempty"`
	Remark        string        `json:",omitempty"`
	Stop          StopType      `json:",omitempty"`
	StopPriceType StopPriceType `json:",omitempty"`
	StopPrice     ujson.Float64 `json:",omitempty"`
	ReduceOnly    bool          `json:",omitempty"`
	CloseOrder    bool          `json:",omitempty"`
	ForceHold     bool          `json:",omitempty"`
	Price         ujson.Float64 `json:",omitempty"`
	Size          ujson.Int64   `json:",omitempty"`
	TimeInForce   TimeInForce   `json:",omitempty"`
	PostOnly      bool          `json:",omitempty"`
	Hidden        bool          `json:",omitempty"`
	Iceberg       bool          `json:",omitempty"`
	VisibleSize   ujson.Int64   `json:",omitempty"`
}

type OrderId struct {
	OrderId string
}

func (o PlaceOrder) Do(c *Client) Response[OrderId] {
	return Post(c, "orders", o, forward[OrderId])
}

func (o *Client) PlaceOrder(v PlaceOrder) Response[OrderId] {
	return v.Do(o)
}

// Get Fills
// https://docs.kucoin.com/futures/#get-fills
type GetFills struct {
	OrderId     string    `url:",omitempty"`
	Symbol      string    `url:",omitempty"`
	Side        Side      `url:",omitempty"`
	Type        OrderType `url:",omitempty"`
	StartAt     int64     `url:",omitempty"`
	EndAt       int64     `url:",omitempty"`
	CurrentPage int64     `url:",omitempty"`
	PageSize    int64     `url:",omitempty"`
}

type Fills struct {
	Symbol         string
	TradeId        string
	OrderId        string
	Side           Side
	Liquidity      string
	ForceTaker     bool
	Price          ujson.Float64
	Size           ujson.Float64
	Value          ujson.Float64
	FeeRate        ujson.Float64
	FixFee         ujson.Float64
	FeeCurrency    string
	Stop           string
	Fee            ujson.Float64
	OrderType      OrderType
	TradeType      string
	CreatedAt      int64
	SettleCurrency string
	TradeTime      int64
	OpenFeePay     ujson.Float64
	CloseFeePay    ujson.Float64
}

func (o GetFills) Do(c *Client) Response[[]Fills] {
	type result struct {
		CurrentPage int
		PageSize    int
		TotalNum    int
		TotalPage   int
		Items       []Fills
	}
	return Get(c, "fills", o, func(r result) ([]Fills, error) {
		return r.Items, nil
	})
}

func (o *Client) GetFills(v GetFills) Response[[]Fills] {
	return v.Do(o)
}

// Add Margin Manually
// https://docs.kucoin.com/futures/#add-margin-manually
type AddMarginManually struct {
	Symbol string
	Margin float64
	BizNo  string
}

type AddingMarginResponse struct {
	Id                string
	Symbol            string
	AutoDeposit       bool
	MaintMarginReq    ujson.Float64
	RiskLimit         ujson.Float64
	RealLeverage      ujson.Float64
	CrossMode         bool
	DelevPercentage   ujson.Float64
	OpeningTimestamp  ujson.Int64
	CurrentTimestamp  ujson.Int64
	CurrentQty        ujson.Int64
	CurrentCost       ujson.Float64
	CurrentComm       ujson.Float64
	UnrealisedCost    ujson.Float64
	RealisedGrossCost ujson.Float64
	RealisedCost      ujson.Float64
	IsOpen            bool
	MarkPrice         ujson.Float64
	MarkValue         ujson.Float64
	PostCost          ujson.Float64
	PosCross          ujson.Float64
	PosInit           ujson.Float64
	PosComm           ujson.Float64
	PosLoss           ujson.Float64
	PosMargin         ujson.Float64
	PosMaint          ujson.Float64
	MaintMargin       ujson.Float64
	RealisedGrossPnl  ujson.Float64
	RealisedPnl       ujson.Float64
	UnrealisedPnl     ujson.Float64
	UnrealisedPnlPcnt ujson.Float64
	UnrealisedRoePcnt ujson.Float64
	AvgEntryPrice     ujson.Float64
	LiquidationPrice  ujson.Float64
	BankruptPrice     ujson.Float64
	UserId            ujson.Int64
	SettleCurrency    string
}

func (o AddMarginManually) Do(c *Client) Response[AddingMarginResponse] {
	return Post(c, "position/margin/deposit-margin", o, forward[AddingMarginResponse])
}

func (o *Client) AddMarginManually(v AddMarginManually) Response[AddingMarginResponse] {
	return v.Do(o)
}

// Get Details of a Single Order
// https://docs.kucoin.com/futures/#get-details-of-a-single-order
type GetDetailsOfSingleOrder struct {
	OrderId   string `url:",omitempty"`
	ClientOid string `url:",omitempty"`
}

type DetailsOfSingleOrder struct {
	Id             string
	Symbol         string
	Type           OrderType
	Side           Side
	Price          ujson.Float64
	Size           ujson.Float64
	Value          ujson.Float64
	DealValue      ujson.Float64
	DealSize       ujson.Float64
	Stp            string
	Stop           string
	StopPriceType  string
	StopTriggered  bool
	StopPrice      ujson.Float64
	TimeInForce    TimeInForce
	PostOnly       bool
	Hidden         bool
	Iceberg        bool
	Leverage       ujson.Float64
	ForceHold      bool
	CloseOrder     bool
	VisibleSize    ujson.Float64
	ClientOid      string
	Remark         string
	Tags           string
	IsActive       bool
	CancelExist    bool
	CreatedAt      ujson.TimeMs
	UpdatedAt      ujson.TimeMs
	EndAt          ujson.TimeMs
	OrderTime      int64
	SettleCurrency string
	Status         Status
	FilledValue    ujson.Float64
	FilledSize     ujson.Float64
	ReduceOnly     bool
}

func (o GetDetailsOfSingleOrder) Do(c *Client) Response[DetailsOfSingleOrder] {
	var url string
	if o.OrderId == "" {
		url = "orders/byClientOid"
	} else {
		url = "orders/" + o.OrderId
	}
	return Get(c, url, o, forward[DetailsOfSingleOrder])
}

func (o *Client) GetDetailsOfSingleOrder(v GetDetailsOfSingleOrder) Response[DetailsOfSingleOrder] {
	return v.Do(o)
}
