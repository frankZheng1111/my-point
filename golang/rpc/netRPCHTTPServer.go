package main

import (
	"errors"
	"log"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}

// 商: 除数结果
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
	quo.Rem = args.A % args.B
	return nil
}

// 实现RPC服务器:
func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	e := http.ListenAndServe(":1234", nil)
	if e != nil {
		log.Fatal("listen error:", e)
	} // else {
	// time.Sleep(3600 * time.Second)
	// }
}
