package http

import (
	"bytes"
	"net/http"

	"github.com/gin-gonic/gin"

	httpcontract "gopkg.in/go-mixed/framework.v1/contracts/http"
)

type GinResponse struct {
	instance *gin.Context
	origin   httpcontract.ResponseOrigin
}

func NewGinResponse(instance *gin.Context, origin httpcontract.ResponseOrigin) *GinResponse {
	return &GinResponse{instance, origin}
}

func (r *GinResponse) Data(code int, contentType string, data []byte) {
	r.instance.Data(code, contentType, data)
}

func (r *GinResponse) Download(filepath, filename string) {
	r.instance.FileAttachment(filepath, filename)
}

func (r *GinResponse) File(filepath string) {
	r.instance.File(filepath)
}

func (r *GinResponse) Header(key, value string) httpcontract.Response {
	r.instance.Header(key, value)

	return r
}

func (r *GinResponse) SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool) httpcontract.Response {
	r.instance.SetCookie(name, value, maxAge, path, domain, secure, httpOnly)

	return r
}

func (r *GinResponse) Json(code int, obj any) {
	r.instance.JSON(code, obj)
}

func (r *GinResponse) Origin() httpcontract.ResponseOrigin {
	return r.origin
}

func (r *GinResponse) Redirect(code int, location string) {
	r.instance.Redirect(code, location)
}

func (r *GinResponse) String(code int, format string, values ...any) {
	r.instance.String(code, format, values...)
}

func (r *GinResponse) Success() httpcontract.ResponseSuccess {
	return NewGinSuccess(r.instance)
}

func (r *GinResponse) Writer() http.ResponseWriter {
	return r.instance.Writer
}

type GinSuccess struct {
	instance *gin.Context
}

func NewGinSuccess(instance *gin.Context) httpcontract.ResponseSuccess {
	return &GinSuccess{instance}
}

type JsonRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
	Err  error  `json:"err,omitempty"`
}

func (r *GinSuccess) Data(contentType string, data []byte) {
	r.instance.Data(http.StatusOK, contentType, data)
}

func (r *GinSuccess) Json(code int, msg string, data any) {
	r.instance.JSON(http.StatusOK, &JsonRes{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func (r *GinSuccess) String(format string, values ...any) {
	r.instance.String(http.StatusOK, format, values...)
}

func (r *GinResponse) Error() httpcontract.ResponseError {
	return NewGinError(r.instance)
}

type GinError struct {
	instance *gin.Context
}

func NewGinError(instance *gin.Context) httpcontract.ResponseError {
	return &GinError{instance}
}

func (r *GinError) Data(contentType string, data []byte) {
	r.instance.Data(http.StatusInternalServerError, contentType, data)
}

func (r *GinError) Json(code int, msg string, err error) {
	if msg == "" {
		msg = err.Error()
	}
	r.instance.JSON(http.StatusInternalServerError, &JsonRes{
		Code: code,
		Msg:  msg,
		Err:  err,
	})
}

func (r *GinError) String(format string, values ...any) {
	r.instance.String(http.StatusInternalServerError, format, values...)
}

func GinResponseMiddleware() httpcontract.Middleware {
	return func(ctx httpcontract.Context) {
		blw := &BodyWriter{body: bytes.NewBufferString("")}
		switch ctx.(type) {
		case *GinContext:
			blw.ResponseWriter = ctx.(*GinContext).Instance().Writer
			ctx.(*GinContext).Instance().Writer = blw
		}

		ctx.WithValue("responseOrigin", blw)
		ctx.Request().Next()
	}
}

type BodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *BodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)

	return w.ResponseWriter.Write(b)
}

func (w *BodyWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)

	return w.ResponseWriter.WriteString(s)
}

func (w *BodyWriter) Body() *bytes.Buffer {
	return w.body
}
