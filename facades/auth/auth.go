package auth

import (
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/contracts/auth"
	"gopkg.in/go-mixed/framework.v1/contracts/http"
)

func getAuth() auth.IAuth {
	return container.MustMake[auth.IAuth]("auth")
}

func Guard(name string) auth.IAuth {
	return getAuth().Guard(name)
}

func Parse(ctx http.Context, token string) error {
	return getAuth().Parse(ctx, token)
}

func User(ctx http.Context, user any) error {
	return getAuth().User(ctx, user)
}

func Login(ctx http.Context, user any) (token string, err error) {
	return getAuth().Login(ctx, user)
}

func LoginUsingID(ctx http.Context, id any) (token string, err error) {
	return getAuth().LoginUsingID(ctx, id)
}

func Refresh(ctx http.Context) (token string, err error) {
	return getAuth().Refresh(ctx)
}

func Logout(ctx http.Context) error {
	return getAuth().Logout(ctx)
}
