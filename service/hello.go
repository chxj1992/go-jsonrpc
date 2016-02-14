package service

import "net/http"

type HelloService struct{}

func (p *HelloService) Say(r *http.Request, args *string, reply *string) error {
	*reply = "Hello, " + *args + "!"
	return nil
}