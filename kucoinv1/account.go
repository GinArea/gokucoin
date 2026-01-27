package kucoinv1

import "github.com/msw-x/moon/ujson"

// AccountSummary - response for GET /api/v2/user-info (Spot API v2)
// https://www.kucoin.com/docs-new/rest/account-info/account-funding/get-account-summary-info
// VIP-level only
type AccountSummary struct {
	// Level - User's account level or tier classification
	Level int `json:"level"`
	// SubQuantity - Total number of sub-accounts currently created
	SubQuantity int `json:"subQuantity"`
	// MaxSubQuantity - Maximum allowed sub-accounts for this account
	MaxSubQuantity int `json:"maxSubQuantity"`
	// MaxDefaultSubQuantity - Default maximum sub-account limit
	MaxDefaultSubQuantity int `json:"maxDefaultSubQuantity"`
	// SpotSubQuantity - Count of sub-accounts with spot trading enabled
	SpotSubQuantity int `json:"spotSubQuantity"`
	// MarginSubQuantity - Count of sub-accounts with margin trading enabled
	MarginSubQuantity int `json:"marginSubQuantity"`
	// FuturesSubQuantity - Count of sub-accounts with futures trading enabled
	FuturesSubQuantity int `json:"futuresSubQuantity"`
	// OptionSubQuantity - Count of sub-accounts with options trading enabled
	OptionSubQuantity int `json:"optionSubQuantity"`
	// MaxSpotSubQuantity - Maximum spot trading sub-accounts allowed
	MaxSpotSubQuantity int `json:"maxSpotSubQuantity"`
	// MaxMarginSubQuantity - Maximum margin trading sub-accounts allowed
	MaxMarginSubQuantity int `json:"maxMarginSubQuantity"`
	// MaxFuturesSubQuantity - Maximum futures trading sub-accounts allowed
	MaxFuturesSubQuantity int `json:"maxFuturesSubQuantity"`
	// MaxOptionSubQuantity - Maximum options trading sub-accounts allowed
	MaxOptionSubQuantity int `json:"maxOptionSubQuantity"`
}

type GetAccountSummary struct{}

func (o GetAccountSummary) Do(c *Client) Response[AccountSummary] {
	cc := c.Copy().WithBaseUrl(SpotBaseUrl).WithPath("api/v2")
	return Get(cc, "user-info", struct{}{}, forward[AccountSummary])
}

func (o *Client) GetAccountSummary() Response[AccountSummary] {
	return GetAccountSummary{}.Do(o)
}

// ApiKeyInfo - response for GET /api/v1/user/api-key (Spot API v1)
// https://www.kucoin.com/docs-new/rest/account-info/account-funding/get-apikey-info
type ApiKeyInfo struct {
	// Remark - User-defined label or note for the API key
	Remark string `json:"remark"`
	// ApiKey - The unique identifier for the API key
	ApiKey string `json:"apiKey"`
	// ApiVersion - Version number of the API key (e.g., 3)
	ApiVersion int `json:"apiVersion"`
	// Permission - Comma-separated list of granted permissions (e.g., " "General,Futures,Unified,Spot,Earn,InnerTransfer,Margin,LeadtradeFutures"")
	Permission string `json:"permission"`
	// IpWhitelist - IP addresses authorized to use this API key
	IpWhitelist string `json:"ipWhitelist"`
	// CreatedAt - Timestamp indicating when the API key was created
	CreatedAt int64 `json:"createdAt"`
	// Uid - User ID associated with the API key
	Uid int64 `json:"uid"`
	// IsMaster - Boolean indicating if this is a master account API key
	IsMaster bool `json:"isMaster"`
	// Region - Geographic region designation for the API key
	Region string `json:"region"`
	// KycStatus - KYC (Know Your Customer) verification status level
	KycStatus int `json:"kycStatus"`
	// SiteType - Site type identifier
	SiteType string `json:"siteType"`
}

type GetApiKeyInfo struct{}

func (o GetApiKeyInfo) Do(c *Client) Response[ApiKeyInfo] {
	cc := c.Copy().WithBaseUrl(SpotBaseUrl).WithPath("api/v1")
	return Get(cc, "user/api-key", struct{}{}, forward[ApiKeyInfo])
}

func (o *Client) GetApiKeyInfo() Response[ApiKeyInfo] {
	return GetApiKeyInfo{}.Do(o)
}

// ------- Balances ----------- //

// AccountOverview - response for GET /api/v1/account-overview (Futures API v1)
// https://www.kucoin.com/docs-new/rest/account-info/account-funding/get-account-futures
// Futures balance for each coin per single request
type AccountOverview struct {
	// AccountEquity - Total account equity = marginBalance + unrealisedPNL
	AccountEquity ujson.Float64 `json:"accountEquity"`
	// UnrealisedPNL - Unrealized profit and loss from open positions
	UnrealisedPNL ujson.Float64 `json:"unrealisedPNL"`
	// MarginBalance - Available margin balance in the account
	MarginBalance ujson.Float64 `json:"marginBalance"`
	// PositionMargin - Margin allocated to current open positions
	PositionMargin ujson.Float64 `json:"positionMargin"`
	// OrderMargin - Margin reserved for pending orders
	OrderMargin ujson.Float64 `json:"orderMargin"`
	// FrozenFunds - Funds temporarily locked or unavailable
	FrozenFunds ujson.Float64 `json:"frozenFunds"`
	// AvailableBalance - Total balance available for trading or withdrawal
	AvailableBalance ujson.Float64 `json:"availableBalance"`
	// AvailableMargin - Margin available for new positions or orders
	AvailableMargin ujson.Float64 `json:"availableMargin"`
	// Currency - Account denomination currency (e.g., USDT)
	Currency string `json:"currency"`
	// RiskRatio - Account risk ratio indicator
	RiskRatio ujson.Float64 `json:"riskRatio"`
	// MaxWithdrawAmount - Maximum amount that can be withdrawn currently
	MaxWithdrawAmount ujson.Float64 `json:"maxWithdrawAmount"`
}

type GetAccountOverview struct {
	Currency string `url:"currency,omitempty"` // Default XBT
}

func (o GetAccountOverview) Do(c *Client) Response[AccountOverview] {
	return Get(c, "account-overview", o, forward[AccountOverview])
}

func (o *Client) GetAccountOverview(currency string) Response[AccountOverview] {
	return GetAccountOverview{Currency: currency}.Do(o)
}

// Account - item in GET /api/v1/accounts response (Spot API v1)
// https://www.kucoin.com/docs-new/rest/account-info/account-funding/get-account-list-spot
// Spot balance
//
// This method can be used to obtain balances of all coins at once for
// spot trading only. This is the balance for the Trading account.
// It cannot be used for inverse ones.

// c.GetAccounts()               // All spot accounts
//
// c.GetAccountDetail(accountId) // Single spot account

type Account struct {
	// Id - Unique identifier for the account
	Id string `json:"id"`
	// Currency - The asset type held in the account (e.g., USDT)
	Currency string `json:"currency"`
	// Type - Account classification: "main" for primary holding or "trade" for active trading
	Type string `json:"type"`
	// Balance - Total funds in the account, including both available and held amounts
	Balance ujson.Float64 `json:"balance"`
	// Available - Funds accessible for immediate trading or withdrawal
	Available ujson.Float64 `json:"available"`
	// Holds - Amount currently locked or reserved for open orders
	Holds ujson.Float64 `json:"holds"`
}

type GetAccounts struct {
	Currency string `url:"currency,omitempty"`
	Type     string `url:"type,omitempty"` // main, trade, margin, trade_hf
}

func (o GetAccounts) Do(c *Client) Response[[]Account] {
	cc := c.Copy().WithBaseUrl(SpotBaseUrl).WithPath("api/v1")
	return Get(cc, "accounts", o, forward[[]Account])
}

func (o *Client) GetAccounts() Response[[]Account] {
	return GetAccounts{}.Do(o)
}

// AccountDetail - response for GET /api/v1/accounts/{accountId} (Spot API)
// Get information for a single spot account. Use this endpoint when you know the accountId.
// https://www.kucoin.com/docs-new/rest/account-info/account-funding/get-account-detail-spot
// Useless
type AccountDetail struct {
	// Currency - The asset type held in the account (e.g., USDT)
	Currency string `json:"currency"`
	// Balance - The total amount of the currency currently in the account
	Balance ujson.Float64 `json:"balance"`
	// Available - The portion of the balance accessible for trading or withdrawal
	Available ujson.Float64 `json:"available"`
	// Holds - Amount temporarily locked or reserved for pending orders
	Holds ujson.Float64 `json:"holds"`
}

// GetAccountDetail - Spot API v1
// GET /api/v1/accounts/{accountId}
type GetAccountDetail struct {
	AccountId string
}

func (o GetAccountDetail) Do(c *Client) Response[AccountDetail] {
	cc := c.Copy().WithBaseUrl(SpotBaseUrl).WithPath("api/v1")
	return Get(cc, "accounts/"+o.AccountId, struct{}{}, forward[AccountDetail])
}

func (o *Client) GetAccountDetail(accountId string) Response[AccountDetail] {
	return GetAccountDetail{AccountId: accountId}.Do(o)
}

// Invitee - item in GET /api/v2/affiliate/queryInvitees response (Spot API v2)
// https://www.kucoin.com/docs-new/3474055e0
type Invitee struct {
	// Uid - User ID of invitee
	Uid string `json:"uid"`
	// NickName - User's nickname/email (partially masked)
	NickName string `json:"nickName"`
	// ReferralCode - Referral code used for registration
	ReferralCode string `json:"referralCode"`
	// Country - User's country location
	Country string `json:"country"`
	// RegistrationTime - Registration timestamp (milliseconds)
	RegistrationTime int64 `json:"registrationTime"`
	// CompletedKyc - KYC verification completed
	CompletedKyc bool `json:"completedKyc"`
	// CompletedFirstDeposit - First deposit completed
	CompletedFirstDeposit bool `json:"completedFirstDeposit"`
	// CompletedFirstTrade - First trade completed
	CompletedFirstTrade bool `json:"completedFirstTrade"`
	// Past7dFees - Trading fees from past 7 days
	Past7dFees string `json:"past7dFees"`
	// Past7dCommission - Commission earned past 7 days
	Past7dCommission string `json:"past7dCommission"`
	// TotalCommission - Total lifetime commission
	TotalCommission string `json:"totalCommission"`
	// MyCommissionRate - Current commission rate percentage
	MyCommissionRate string `json:"myCommissionRate"`
	// CashbackRate - Cashback rate percentage
	CashbackRate string `json:"cashbackRate"`
	// Currency - Settlement currency (e.g., USDT)
	Currency string `json:"currency"`
}

// GetInvitees - Spot API v2
// GET /api/v2/affiliate/queryInvitees
// https://www.kucoin.com/docs-new/3474055e0
type GetInvitees struct {
	// UserType - Filter by user type
	UserType string `url:"userType,omitempty"`
	// ReferralCode - Filter by specific referral code
	ReferralCode string `url:"referralCode,omitempty"`
	// Uid - Filter by specific user ID
	Uid string `url:"uid,omitempty"`
	// RegistrationStartAt - Filter by registration start timestamp (milliseconds)
	RegistrationStartAt int64 `url:"registrationStartAt,omitempty"`
	// RegistrationEndAt - Filter by registration end timestamp (milliseconds)
	RegistrationEndAt int64 `url:"registrationEndAt,omitempty"`
	// Page - Current page number for pagination
	Page int `url:"page,omitempty"`
	// PageSize - Number of results per page
	PageSize int `url:"pageSize,omitempty"`
}

func (o GetInvitees) Do(c *Client) Response[Paginated[Invitee]] {
	cc := c.Copy().WithBaseUrl(SpotBaseUrl).WithPath("api/v2")
	return Get(cc, "affiliate/queryInvitees", o, forward[Paginated[Invitee]])
}

func (o *Client) GetInvitees() Response[Paginated[Invitee]] {
	return GetInvitees{}.Do(o)
}

// DepositAddress - item in GET /api/v3/deposit-addresses response (Spot API v3)
// no direct v3 correct documentation
// https://www.kucoin.com/docs-new/abandoned-endpoints/account-funding/get-deposit-addresses-v1
type DepositAddress struct {
	// Address - The deposit address
	Address string `json:"address"`
	// Memo - Memo or tag required for certain blockchains (e.g., XRP, EOS)
	Memo string `json:"memo"`
	// Remark - User-defined remark for the address
	Remark string `json:"remark"`
	// ChainId - Blockchain network identifier (e.g., "trx", "arbitrum")
	ChainId string `json:"chainId"`
	// To - Account type destination: "TRADE" or "MAIN"
	To string `json:"to"`
	// ExpirationDate - Address validity timestamp (0 if no expiration)
	ExpirationDate int64 `json:"expirationDate"`
	// Currency - Cryptocurrency symbol (e.g., "USDT")
	Currency string `json:"currency"`
	// ContractAddress - Smart contract address on the blockchain
	ContractAddress string `json:"contractAddress"`
	// ChainName - Human-readable blockchain name (e.g., "TRC20", "ARBITRUM")
	ChainName string `json:"chainName"`
}

type GetDepositAddresses struct {
	// Currency - Cryptocurrency symbol (e.g., USDT). Required by API despite docs saying optional.
	Currency string `url:"currency,omitempty"`
	// Chain - Blockchain network identifier (e.g., trx, arbitrum)
	Chain string `url:"chain,omitempty"`
}

func (o GetDepositAddresses) Do(c *Client) Response[[]DepositAddress] {
	cc := c.Copy().WithBaseUrl(SpotBaseUrl).WithPath("api/v3")
	return Get(cc, "deposit-addresses", o, forward[[]DepositAddress])
}

func (o *Client) GetDepositAddresses(currency, chain string) Response[[]DepositAddress] {
	return GetDepositAddresses{Currency: currency, Chain: chain}.Do(o)
}

// BrokerUser - item in GET /api/v2/broker/queryUser response (Broker API v2)
// https://www.kucoin.com/docs-new/rest/broker/api-broker/get-user-list
type BrokerUser struct {
	// Uid - User identifier
	Uid string `json:"uid"`
	// NickName - User nickname (partially masked)
	NickName string `json:"nickName"`
	// RegistrationTime - Registration timestamp (milliseconds)
	RegistrationTime int64 `json:"registrationTime"`
	// RateUp - Whether commission rate has been increased
	RateUp bool `json:"rateUp"`
	// TotalCommission - Total commission earned from this user
	TotalCommission string `json:"totalCommission"`
	// TotalTradingVolume - Total trading volume generated by user
	TotalTradingVolume string `json:"totalTradingVolume"`
	// TotalFee - Total trading fees paid by user
	TotalFee string `json:"totalFee"`
	// InvitedByMe - Whether user was directly invited by broker
	InvitedByMe bool `json:"invitedByMe"`
	// Rcode - Referral code used by user (nullable)
	Rcode *string `json:"rcode"`
	// Tags - User tags
	Tags string `json:"tags"`
	// SpotTradingVolumeWithTag - Spot trading volume with tag
	SpotTradingVolumeWithTag string `json:"spotTradingVolumeWithTag"`
	// FutureTradingVolumeWithTag - Futures trading volume with tag
	FutureTradingVolumeWithTag string `json:"futureTradingVolumeWithTag"`
	// TradingFeeWithTag - Trading fees with tag
	TradingFeeWithTag string `json:"tradingFeeWithTag"`
	// CommissionWithTag - Commission earned with tag
	CommissionWithTag string `json:"commissionWithTag"`
	// SpotTradingVolumeWithoutTag - Spot trading volume without tag
	SpotTradingVolumeWithoutTag string `json:"spotTradingVolumeWithoutTag"`
	// FutureTradingVolumeWithoutTag - Futures trading volume without tag
	FutureTradingVolumeWithoutTag string `json:"futureTradingVolumeWithoutTag"`
	// FuturesTradingVolumeWithTag - Futures trading volume with tag (alternate naming)
	FuturesTradingVolumeWithTag string `json:"futuresTradingVolumeWithTag"`
	// FuturesTradingVolumeWithoutTag - Futures trading volume without tag (alternate naming)
	FuturesTradingVolumeWithoutTag string `json:"futuresTradingVolumeWithoutTag"`
	// TradingFeeWithoutTag - Trading fees without tag
	TradingFeeWithoutTag string `json:"tradingFeeWithoutTag"`
	// CommissionWithoutTag - Commission earned without tag
	CommissionWithoutTag string `json:"commissionWithoutTag"`
	// Currency - Settlement currency (e.g., USDT)
	Currency string `json:"currency"`
}

// GetBrokerUsers - Broker API v2
// GET /api/v2/broker/queryUser
// https://www.kucoin.com/docs-new/rest/broker/api-broker/get-user-list
type GetBrokerUsers struct {
	// TradeType - Filter by trading type (e.g., "all")
	TradeType string `url:"tradeType,omitempty"`
	// Uid - Filter by user ID
	Uid string `url:"uid,omitempty"`
	// Rcode - Filter by referral code
	Rcode string `url:"rcode,omitempty"`
	// Tag - Filter by tag
	Tag string `url:"tag,omitempty"`
	// StartAt - Start timestamp (milliseconds)
	StartAt int64 `url:"startAt,omitempty"`
	// EndAt - End timestamp (milliseconds)
	EndAt int64 `url:"endAt,omitempty"`
	// Page - Page number (default: 1)
	Page int `url:"page,omitempty"`
	// PageSize - Records per page (default: 10)
	PageSize int `url:"pageSize,omitempty"`
}

func (o GetBrokerUsers) Do(c *Client) Response[Paginated[BrokerUser]] {
	cc := c.Copy().WithBaseUrl(SpotBaseUrl).WithPath("api/v2")
	return Get(cc, "broker/queryUser", o, forward[Paginated[BrokerUser]])
}

func (o *Client) GetBrokerUsers() Response[Paginated[BrokerUser]] {
	return GetBrokerUsers{}.Do(o)
}
