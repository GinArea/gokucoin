package kucoinv1

// RateLimit - rate limiting information from KuCoin API response headers
// https://www.kucoin.com/docs-new/rate-limit
//
// Rate limits use a 30-second rolling window. VIP0 users have 2,000-4,000
// requests per 30s (weight-based). HTTP 429 is returned when exceeded.
//
// Note: These headers are only returned for private (authenticated) endpoints.
// Public endpoints use IP-based rate limiting without these headers.
type RateLimit struct {
	// Limit - total resource pool quota available for this 30-second window
	Limit int `http:"gw-ratelimit-limit"`
	// Remaining - remaining quota in current cycle
	Remaining int `http:"gw-ratelimit-remaining"`
	// Reset - countdown in milliseconds until quota resets
	Reset int64 `http:"gw-ratelimit-reset"`
}
