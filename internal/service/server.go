package service

import (
	"fmt"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"github.com/IvanKalug-QA/Go-URL-shorter/internal/handler"
)

func StartServer(addr string) {
	mux := router.New()
	mux.GET(`/`, handler.MainPage)
	mux.GET(`/{id}`, handler.MainPage)
	mux.POST(`/`, handler.MainPage)

	handlerChain := fasthttp.CompressHandler(mux.Handler)

	fmt.Println("Runnig server on", addr)
	err := fasthttp.ListenAndServe(addr, handlerChain)
	if err != nil {
		panic(err)
	}
}
