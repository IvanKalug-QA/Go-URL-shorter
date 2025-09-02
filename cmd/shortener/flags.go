package main

import (
	"flag"
	"os"
)

var port string

func ParseArgs() {
	flag.StringVar(&port, `a`, `:8000`, `address and port to run server`)
	flag.Parse()

	if AddrPort := os.Getenv("SERVER_ADDRES"); AddrPort != "" {
		port = AddrPort
	}
}
