package console

type PolicyStubs struct {
}

func (receiver PolicyStubs) Policy() string {
	return `package policies

import (
	"context"
	
	"gopkg.in/go-mixed/framework/contracts/auth/access"
)

type DummyPolicy struct {
}

func NewDummyPolicy() *DummyPolicy {
	return &DummyPolicy{}
}

func (r *DummyPolicy) Create(ctx context.Context, arguments map[string]any) *access.Response {
	return nil
}
`
}
