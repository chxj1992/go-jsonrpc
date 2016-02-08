package main

import (
	"flag"
	"github.com/chxj1992/go-jsonrpc/src"
	"github.com/chxj1992/go-jsonrpc/demo"
)

func main() {
	port := flag.String("p", "12345", "listening port")
	flag.Parse()

	handlers := []src.RpcHandler{
		src.RpcHandler{"/hello", new(demo.HelloService)},
		src.RpcHandler{"/math", new(demo.MathService)},
	};

	src.Serve(*port, handlers)
}
