package main

import "flag"

var port string

func ParseArgs() {
	flag.StringVar(&port, `a`, `:8000`, `address and port to run server`)
	flag.Parse()
}
