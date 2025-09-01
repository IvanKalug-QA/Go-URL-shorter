package main

import (
	"github.com/IvanKalug-QA/Go-URL-shorter/internal/service"
)

func main() {
	ParseArgs()
	service.StartServer(port)
}
