package src

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
)

func Serve(port string, handlers []RpcHandler) {

	s := rpc.NewServer()
	s.RegisterCodec(json2.NewCodec(), "application/json")

	r := mux.NewRouter()
	for _, handler := range handlers {
		s.RegisterService(handler.Service, "")
		r.Handle(handler.Path, s)
	}

	http.ListenAndServe(":" + port, r)
}

type RpcHandler struct {
	Path    string
	Service interface{}
}

