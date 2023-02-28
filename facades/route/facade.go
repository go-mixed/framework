package route

import (
	"gopkg.in/go-mixed/framework.v1/container"
	httpcontract "gopkg.in/go-mixed/framework.v1/contracts/http"
	"gopkg.in/go-mixed/framework.v1/contracts/route"
	"net/http"
)

func getRoute() route.IRouteEngine {
	return container.MustMake[route.IRouteEngine]("route")
}

func Run(host ...string) error {
	return getRoute().Run(host...)
}
func RunTLS(host ...string) error {
	return getRoute().RunTLS(host...)
}
func RunTLSWithCert(host, certFile, keyFile string) error {
	return getRoute().RunTLSWithCert(host, certFile, keyFile)
}
func ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	getRoute().ServeHTTP(writer, request)
}
func GlobalMiddleware(middlewares ...httpcontract.Middleware) {
	getRoute().GlobalMiddleware(middlewares...)
}
func Group(handler route.GroupFunc) {
	getRoute().Group(handler)
}
func Prefix(addr string) route.IRoute {
	return getRoute().Prefix(addr)
}
func Middleware(middlewares ...httpcontract.Middleware) route.IRoute {
	return getRoute().Middleware(middlewares...)
}
func Any(relativePath string, handler httpcontract.HandlerFunc) {
	getRoute().Any(relativePath, handler)
}
func Get(relativePath string, handler httpcontract.HandlerFunc) {
	getRoute().Get(relativePath, handler)
}
func Post(relativePath string, handler httpcontract.HandlerFunc) {
	getRoute().Post(relativePath, handler)
}
func Delete(relativePath string, handler httpcontract.HandlerFunc) {
	getRoute().Delete(relativePath, handler)
}
func Patch(relativePath string, handler httpcontract.HandlerFunc) {
	getRoute().Patch(relativePath, handler)
}
func Put(relativePath string, handler httpcontract.HandlerFunc) {
	getRoute().Put(relativePath, handler)
}
func Options(relativePath string, handler httpcontract.HandlerFunc) {
	getRoute().Options(relativePath, handler)
}
func Static(relativePath, root string) {
	getRoute().Static(relativePath, root)
}
func StaticFile(relativePath, filepath string) {
	getRoute().StaticFile(relativePath, filepath)
}
func StaticFS(relativePath string, fs http.FileSystem) {
	getRoute().StaticFS(relativePath, fs)
}
