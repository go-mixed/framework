package http

import (
	"gopkg.in/go-mixed/framework.v1/contracts/http"
)

type RateLimiter struct {
	limiters map[string]func(ctx http.Context) []http.ILimit
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		limiters: make(map[string]func(ctx http.Context) []http.ILimit),
	}
}

func (r *RateLimiter) For(name string, callback func(ctx http.Context) http.ILimit) {
	r.limiters[name] = func(ctx http.Context) []http.ILimit {
		return []http.ILimit{callback(ctx)}
	}
}

func (r *RateLimiter) ForWithLimits(name string, callback func(ctx http.Context) []http.ILimit) {
	r.limiters[name] = callback
}

func (r *RateLimiter) Limiter(name string) func(ctx http.Context) []http.ILimit {
	return r.limiters[name]
}
