package kucoinv1

type RateLimit struct {
	Remaining       int   `http:"X-RateLimit-Remaining"`
	RequestedTokens int   `http:"X-RateLimit-Requested-Tokens"`
	BurstCapacity   int64 `http:"X-RateLimit-Burst-Capacity"`
	ReplenishRate   int64 `http:"X-RateLimit-Replenish-Rate"`
}
