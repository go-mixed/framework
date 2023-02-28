package middleware

import (
	"fmt"

	"gopkg.in/go-mixed/framework.v1/contracts/http"
	ihttp "gopkg.in/go-mixed/framework.v1/facades/http"
	httplimit "gopkg.in/go-mixed/framework.v1/http/limit"
)

func Throttle(name string) http.Middleware {
	return func(ctx http.Context) {
		if limiter := ihttp.Limiter(name); limiter != nil {
			if limits := limiter(ctx); len(limits) > 0 {
				for _, limit := range limits {
					if instance, ok := limit.(*httplimit.Limit); ok {
						fmt.Println(instance.Key, instance.MaxAttempts, instance.DecayMinutes)
						if instance.ResponseCallback != nil {
							instance.ResponseCallback(ctx)
						}
						// TODO Determine whether to pass the ILimit check
					}
				}
			}
		}

		ctx.Request().Next()

		// TODO calculate remaining attempts(if needed)
	}
}
