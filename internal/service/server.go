package service

import (
	"net/http"

	"github.com/IvanKalug-QA/Go-URL-shorter/internal/handler"
)

func StartServer() {
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, handler.MainPage)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
