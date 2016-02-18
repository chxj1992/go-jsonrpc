package main

import (
	"github.com/chxj1992/go-jsonrpc/jsonrpc"
	"testing"
	"github.com/stretchr/testify/assert"
)


type MathArgs struct {
	First  int
	Second int
}
type MathReply struct {
	Result int
}


func TestHello(t *testing.T) {
	client := jsonrpc.Client{"http://localhost:12345/hello"}

	args := "World"
	var reply string

	client.Call("HelloService.Say", args, &reply)
	assert.Equal(t, reply, "Hello, World!")
}

func TestMath(t *testing.T) {
	client := jsonrpc.Client{"http://localhost:12345/math"}

	mathArgs := MathArgs{1, 2}
	//	mathArgs := map[string]int{"First":1, "Second":2}
	var mathReply MathReply

	client.Call("MathService.Add", mathArgs, &mathReply)

	assert.Equal(t, mathReply.Result, 3)
}

