package access

import "context"

//go:generate mockery --name=IGate
type IGate interface {
	WithContext(ctx context.Context) IGate
	Allows(ability string, arguments map[string]any) bool
	Denies(ability string, arguments map[string]any) bool
	Inspect(ability string, arguments map[string]any) IResponse
	Define(ability string, callback func(ctx context.Context, arguments map[string]any) IResponse)
	Any(abilities []string, arguments map[string]any) bool
	None(abilities []string, arguments map[string]any) bool
	Before(callback func(ctx context.Context, ability string, arguments map[string]any) IResponse)
	After(callback func(ctx context.Context, ability string, arguments map[string]any, result IResponse) IResponse)
}

type IResponse interface {
	Allowed() bool
	Message() string
}
