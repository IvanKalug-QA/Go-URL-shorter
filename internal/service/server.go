package service

import (
	"fmt"
	"strings"
	"os"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"github.com/IvanKalug-QA/Go-URL-shorter/internal/controller/logger"
	"github.com/IvanKalug-QA/Go-URL-shorter/internal/handler"
	"github.com/sirupsen/logrus"
)

func StartServer(addr string) {
	log := logrus.New()
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)

	controler := logger.NewBaseController(log)

	mux := router.New()
	mux.GET(`/`, controler.LoggingMiddleware(GzipMiddleware(handler.MainPage)))
  mux.GET(`/{id}`, controler.LoggingMiddleware(GzipMiddleware(handler.MainPage)))
	mux.POST(`/`, controler.LoggingMiddleware(GzipMiddleware(handler.MainPage)))

	handlerChain := fasthttp.CompressHandler(mux.Handler)

	fmt.Println("Runnig server on", addr)
	err := fasthttp.ListenAndServe(addr, handlerChain)
	if err != nil {
		panic(err)
	}
}

func GzipMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		contentEncoding := strings.ToLower(string(ctx.Request.Header.Peek("Content-Encoding")))
		if contentEncoding == "gzip" {
			body, err := ctx.Request.BodyGunzip()
			if err == nil {
				ctx.Request.SetBody(body)
				ctx.Request.Header.Del("Content-Encoding")
			}
		}
		acceptEncoding := strings.ToLower(string(ctx.Request.Header.Peek("Accept-Encoding")))
		supportGzip := strings.Contains(acceptEncoding, "gzip")
		next(ctx)
		if supportGzip && ShouldCompresGzip(ctx) {
			comressBody := fasthttp.AppendGzipBytes(nil, ctx.Response.Body())
			ctx.Response.SetBody(comressBody)
			ctx.Response.Header.Set("Content-Encoding", "gzip")
		}
	}
}

func ShouldCompresGzip(ctx *fasthttp.RequestCtx) bool {
	contentType := strings.ToLower(string(ctx.Response.Header.Peek("Content-Type")))
	contentTypes := []string{
		"application/json",
		"text/html",
	}
	for _, t := range contentTypes {
		if contentType == t {
			return true
		}
	}
	return false
}
