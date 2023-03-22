package auth

import (
	"context"
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/contracts/auth/access"
)

func getGate() access.IGate {
	return container.MustMakeAs("gate", access.IGate(nil))
}

func WithContext(ctx context.Context) access.IGate {
	return getGate().WithContext(ctx)
}

func Allows(ability string, arguments map[string]any) bool {
	return getGate().Allows(ability, arguments)
}

func Denies(ability string, arguments map[string]any) bool {
	return getGate().Denies(ability, arguments)
}

func Inspect(ability string, arguments map[string]any) access.IResponse {
	return getGate().Inspect(ability, arguments)
}

func Define(ability string, callback func(ctx context.Context, arguments map[string]any) access.IResponse) {
	getGate().Define(ability, callback)
}

func Any(abilities []string, arguments map[string]any) bool {
	return getGate().Any(abilities, arguments)
}

func None(abilities []string, arguments map[string]any) bool {
	return getGate().None(abilities, arguments)
}

func Before(callback func(ctx context.Context, ability string, arguments map[string]any) access.IResponse) {
	getGate().Before(callback)
}

func After(callback func(ctx context.Context, ability string, arguments map[string]any, result access.IResponse) access.IResponse) {
	getGate().After(callback)
}
