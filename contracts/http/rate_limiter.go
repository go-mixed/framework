package http

type IRateLimiter interface {
	For(name string, callback func(ctx Context) ILimit)
	ForWithLimits(name string, callback func(ctx Context) []ILimit)
	Limiter(name string) func(ctx Context) []ILimit
}

type ILimit interface {
	By(key string) ILimit
	Response(func(ctx Context)) ILimit
}
