package main

import (
	"github.com/chxj1992/go-jsonrpc/src"
	"github.com/chxj1992/go-jsonrpc/demo"
)

func main() {
	client := src.Client{"http://localhost:12345/hello"}

	var reply demo.HelloReply
	client.Call("HelloService.Say", demo.HelloArgs{"Andy"}, &reply)

	print(reply.Message)
}
