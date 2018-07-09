package main

import (
	"errors"
	"log"
	"net/http"
	"net/rpc"
	"time"
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
	time.Sleep(5 * time.Second)
	*reply = args.A * args.B
	return nil
}
func (t *Arith) Divide(args *Args, quo *Quotient) error {
	time.Sleep(5 * time.Second)
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
	// HandleHTTP registers an HTTP handler for RPC messages to DefaultServer on DefaultRPCPath and a debugging handler on DefaultDebugPath.
	// 	DefaultRPCPath   = "/_goRPC_"
	// 	DefaultDebugPath = "/debug/rpc"
	// It is still necessary to invoke http.Serve(), typically in a go statement.
	// 自定义Server
	// func NewServer() *Server
	// NewServer returns a new Server.
	//
	// func (server *Server) Accept(lis net.Listener)
	// Accept accepts connections on the listener and serves requests for each incoming connection. Accept blocks until the listener returns a non-nil error. The caller typically invokes Accept in a go statement.
	//
	// func (server *Server) HandleHTTP(rpcPath, debugPath string)
	// HandleHTTP registers an HTTP handler for RPC messages on rpcPath, and a debugging handler on debugPath. It is still necessary to invoke http.Serve(), typically in a go statement.
	// 可以同时处理多个请求
	rpc.HandleHTTP()
	e := http.ListenAndServe(":1234", nil)
	if e != nil {
		log.Fatal("listen error:", e)
	} // else {
	// time.Sleep(3600 * time.Second)
	// }
}
