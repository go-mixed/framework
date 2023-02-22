package route

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gookit/color"

	httpcontract "gopkg.in/go-mixed/framework.v1/contracts/http"
	"gopkg.in/go-mixed/framework.v1/contracts/route"
	"gopkg.in/go-mixed/framework.v1/facades"
	goravelhttp "gopkg.in/go-mixed/framework.v1/http"
)

type Gin struct {
	route.Route
	instance *gin.Engine
}

func NewGin() *Gin {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	if debugLog := getDebugLog(); debugLog != nil {
		engine.Use(debugLog)
	}

	return &Gin{instance: engine, Route: NewGinGroup(
		engine.Group("/"),
		"",
		[]httpcontract.Middleware{},
		[]httpcontract.Middleware{goravelhttp.GinResponseMiddleware()},
	)}
}

func (r *Gin) Run(host ...string) error {
	if len(host) == 0 {
		defaultHost := facades.Config.GetString("route.host")
		if defaultHost == "" {
			return errors.New("host can't be empty")
		}

		host = append(host, defaultHost)
	}

	outputRoutes(r.instance.Routes())
	color.Greenln("[HTTP] Listening and serving HTTP on " + host[0])

	return r.instance.Run([]string{host[0]}...)
}

func (r *Gin) RunTLS(host ...string) error {
	if len(host) == 0 {
		defaultHost := facades.Config.GetString("route.tls.host")
		if defaultHost == "" {
			return errors.New("host can't be empty")
		}

		host = append(host, defaultHost)
	}

	certFile := facades.Config.GetString("route.tls.ssl.cert")
	keyFile := facades.Config.GetString("route.tls.ssl.key")

	return r.RunTLSWithCert(host[0], certFile, keyFile)
}

func (r *Gin) RunTLSWithCert(host, certFile, keyFile string) error {
	if host == "" {
		return errors.New("host can't be empty")
	}
	if certFile == "" || keyFile == "" {
		return errors.New("certificate can't be empty")
	}

	outputRoutes(r.instance.Routes())
	color.Greenln("[HTTPS] Listening and serving HTTPS on " + host)

	return r.instance.RunTLS(host, certFile, keyFile)
}

func (r *Gin) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	r.instance.ServeHTTP(writer, request)
}

func (r *Gin) GlobalMiddleware(middlewares ...httpcontract.Middleware) {
	if len(middlewares) > 0 {
		r.instance.Use(middlewaresToGinHandlers(middlewares)...)
	}
	r.Route = NewGinGroup(
		r.instance.Group("/"),
		"",
		[]httpcontract.Middleware{},
		[]httpcontract.Middleware{goravelhttp.GinResponseMiddleware()},
	)
}
