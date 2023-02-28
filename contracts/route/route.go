package route

import (
	"net/http"

	httpcontract "gopkg.in/go-mixed/framework.v1/contracts/http"
)

type GroupFunc func(routes IRoute)

//go:generate mockery --name=Engine
type IRouteEngine interface {
	IRoute
	Run(host ...string) error
	RunTLS(host ...string) error
	RunTLSWithCert(host, certFile, keyFile string) error
	ServeHTTP(writer http.ResponseWriter, request *http.Request)
	GlobalMiddleware(middlewares ...httpcontract.Middleware)
}

//go:generate mockery --name=Route
type IRoute interface {
	Group(handler GroupFunc)
	Prefix(addr string) IRoute
	Middleware(middlewares ...httpcontract.Middleware) IRoute

	Any(relativePath string, handler httpcontract.HandlerFunc)
	Get(relativePath string, handler httpcontract.HandlerFunc)
	Post(relativePath string, handler httpcontract.HandlerFunc)
	Delete(relativePath string, handler httpcontract.HandlerFunc)
	Patch(relativePath string, handler httpcontract.HandlerFunc)
	Put(relativePath string, handler httpcontract.HandlerFunc)
	Options(relativePath string, handler httpcontract.HandlerFunc)

	Static(relativePath, root string)
	StaticFile(relativePath, filepath string)
	StaticFS(relativePath string, fs http.FileSystem)
}
