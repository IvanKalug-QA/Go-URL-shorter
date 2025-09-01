package service

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"github.com/IvanKalug-QA/Go-URL-shorter/internal/handler"
)

func StartServer() {
	mux := router.New()
	mux.GET(`/`, handler.MainPage)
	mux.GET(`/{id}`, handler.MainPage)
	mux.POST(`/`, handler.MainPage)

	handlerChain := fasthttp.CompressHandler(mux.Handler)

	err := fasthttp.ListenAndServe(`:8080`, handlerChain)
	if err != nil {
		panic(err)
	}
}
