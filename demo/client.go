package main

import (
	"github.com/chxj1992/go-jsonrpc/jsonrpc"
	"fmt"
	"strconv"
)


type MathArgs struct {
	First  int
	Second int
}
type MathReply struct {
	Result int
}

func main() {
	client := jsonrpc.Client{"http://localhost:12345/rpc"}

	args := "World"
	var reply string

	client.Call("HelloService.Say", args, &reply)
	fmt.Println("message: " + reply)

	mathArgs := MathArgs{1, 2}
	//	mathArgs := map[string]int{"First":1, "Second":2}
	var mathReply MathReply

	client.Call("MathService.Add", mathArgs, &mathReply)
	fmt.Println("result: " + strconv.Itoa(mathReply.Result))
}


