package server

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

// func (t *T) MethodName(argType T1, replyType *T2) error

// 方法的返回值如果不为空， 那么它作为一个字符串返回给调用者

//参数
type Args struct {
	A, B int
}

//sang
type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {

	*reply = args.A * args.B

	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {

	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Quo = args.A % args.B

	return nil
}

func Server() {
	arith := new(Arith)
	rpc.Register(arith)

	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)

	}
	go http.Serve(l, nil)

}
