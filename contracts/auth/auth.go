package auth

import "gopkg.in/go-mixed/framework.v1/contracts/http"

//go:generate mockery --name=IAuth
type IAuth interface {
	Guard(name string) IAuth
	Parse(ctx http.Context, token string) error
	User(ctx http.Context, user any) error
	Login(ctx http.Context, user any) (token string, err error)
	LoginUsingID(ctx http.Context, id any) (token string, err error)
	Refresh(ctx http.Context) (token string, err error)
	Logout(ctx http.Context) error
}
