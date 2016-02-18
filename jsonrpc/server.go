package jsonrpc

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
	"log"
)

func Serve(port string, handlers []RpcHandler) {
	r := mux.NewRouter()
	for _, handler := range handlers {
		s := rpc.NewServer()
		s.RegisterCodec(json2.NewCodec(), "application/json")
		s.RegisterService(handler.Service, "")
		r.Handle(handler.Path, s)
	}

	log.Println("Listen on port: " + port)
	http.ListenAndServe(":" + port, r)
}

type RpcHandler struct {
	Path    string
	Service interface{}
}

