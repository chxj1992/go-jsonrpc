package main

import (
	"flag"
	"github.com/chxj1992/go-jsonrpc/jsonrpc"
	"github.com/chxj1992/go-jsonrpc/service"
	"log"
)

func main() {
	port := flag.String("p", "12345", "listening port")
	flag.Parse()

	handlers := []jsonrpc.RpcHandler{
		jsonrpc.RpcHandler{"/rpc", new(service.HelloService)},
		jsonrpc.RpcHandler{"/rpc", new(service.MathService)},
	};

	log.Println("Listen on port: " + *port)
	jsonrpc.Serve(*port, handlers)
}
