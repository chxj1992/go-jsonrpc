package main

import (
	"flag"
	"github.com/chxj1992/go-jsonrpc/jsonrpc"
	"github.com/chxj1992/go-jsonrpc/service"
)

func main() {
	port := flag.String("p", "12345", "listening port")
	flag.Parse()

	handlers := []jsonrpc.RpcHandler{
		jsonrpc.RpcHandler{"/hello", new(service.HelloService)},
		jsonrpc.RpcHandler{"/math", new(service.MathService)},
	};

	jsonrpc.Serve(*port, handlers)
}
