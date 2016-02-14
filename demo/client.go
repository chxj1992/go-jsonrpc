package main

import (
	"github.com/chxj1992/go-jsonrpc/src"
	"fmt"
)

type Args struct {
	Who string
}

type Reply struct {
	Message string
}

func main() {
	client := src.Client{"http://localhost:12345/rpc"}

	args := Args{"World"}
	var reply Reply

	client.Call("HelloService.Say", args, &reply)
	fmt.Println("message: " + reply.Message)
}


