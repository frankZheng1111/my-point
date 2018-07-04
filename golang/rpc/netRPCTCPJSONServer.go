package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Args struct {
	A, B int
}

// 商: 除数结果
type Quotient struct {
	Quo, Rem int
}

type Arith struct {
	setting int
}

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	args.A = 30 // 修改不会映射到原参数
	return nil
}
func (t *Arith) Divide(args *Args, quo *Quotient) error {
	fmt.Println("setting: ", t.setting)
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

// 实现RPC服务器:
func main() {
	var ms = new(Arith)
	ms.setting = 10
	rpc.Register(ms)
	var address, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:1234")
	listener, err := net.ListenTCP("tcp", address)
	if err != nil {
		fmt.Println("启动失败！", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		fmt.Println("接收到一个调用请求...")
		jsonrpc.ServeConn(conn)
	}
	// time.Sleep(3600 * time.Second)
}
