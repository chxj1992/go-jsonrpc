package main

import (
	"flag"
	"github.com/chxj1992/go-jsonrpc/src"
	"github.com/chxj1992/go-jsonrpc/service"
)

func main() {
	port := flag.String("p", "12345", "listening port")
	flag.Parse()

	handlers := []src.RpcHandler{
		src.RpcHandler{"/rpc", new(service.HelloService)},
		src.RpcHandler{"/rpc", new(service.MathService)},
	};

	src.Serve(*port, handlers)
}
