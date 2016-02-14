package service

import "net/http"

type MathArgs struct {
	First int
	Second int
}
type MathReply struct {
	Result int
}
type MathService struct{}

func (p *MathService) Add(r *http.Request, args *MathArgs, reply *MathReply) error {
	reply.Result =  args.First + args.Second
	return nil
}