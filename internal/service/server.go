package service

import (
	"fmt"
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
	mux.GET(`/`, controler.LoggingMiddleware(handler.MainPage))
	mux.GET(`/{id}`, controler.LoggingMiddleware(handler.MainPage))
	mux.POST(`/`, controler.LoggingMiddleware(handler.MainPage))

	handlerChain := fasthttp.CompressHandler(mux.Handler)

	fmt.Println("Runnig server on", addr)
	err := fasthttp.ListenAndServe(addr, handlerChain)
	if err != nil {
		panic(err)
	}
}
