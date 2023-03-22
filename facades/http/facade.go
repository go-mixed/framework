package http

import (
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/contracts/http"
)

func getHttp() http.IRateLimiter {
	return container.MustMakeAs("http", http.IRateLimiter(nil))
}

func For(name string, callback func(ctx http.Context) http.ILimit) {
	getHttp().For(name, callback)
}
func ForWithLimits(name string, callback func(ctx http.Context) []http.ILimit) {
	getHttp().ForWithLimits(name, callback)
}
func Limiter(name string) func(ctx http.Context) []http.ILimit {
	return getHttp().Limiter(name)
}
