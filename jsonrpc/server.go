package jsonrpc

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
	"github.com/fvbock/endless"
)

func Serve(port string, handlers []RpcHandler) {
	r := mux.NewRouter()
	for _, handler := range handlers {
		s := rpc.NewServer()
		s.RegisterCodec(json2.NewCodec(), "application/json")
		s.RegisterService(handler.Service, "")
		r.Handle(handler.Path, s)
	}

	endless.ListenAndServe(":" + port, r)
}

type RpcHandler struct {
	Path    string
	Service interface{}
}

