package console

type Stubs struct {
}

func (r Stubs) Request() string {
	return `package requests

import (
	"gopkg.in/go-mixed/framework/contracts/http"
	"gopkg.in/go-mixed/framework/contracts/validation"
)

type DummyRequest struct {
	DummyField
}

func (r *DummyRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *DummyRequest) Rules() map[string]string {
	return map[string]string{}
}

func (r *DummyRequest) Messages() map[string]string {
	return map[string]string{}
}

func (r *DummyRequest) Attributes() map[string]string {
	return map[string]string{}
}

func (r *DummyRequest) PrepareForValidation(data validation.Data) error {
	return nil
}
`
}

func (r Stubs) Controller() string {
	return `package controllers

import (
	"gopkg.in/go-mixed/framework/contracts/http"
)

type DummyController struct {
	//Dependent services
}

func NewDummyController() *DummyController {
	return &DummyController{
		//Inject services
	}
}

func (r *DummyController) Index(ctx http.Context) {
}	
`
}

func (r Stubs) Middleware() string {
	return `package middleware

import (
	contractshttp "gopkg.in/go-mixed/framework/contracts/http"
)

func DummyMiddleware() contractshttp.Middleware {
	return func(ctx contractshttp.Context) {
		ctx.Request().Next()
	}
}
`
}
