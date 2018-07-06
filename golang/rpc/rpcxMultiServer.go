package main

import (
	"context"
	"fmt"
	"github.com/smallnest/rpcx/server"
)

type Args struct {
	A int `msg:"a"`
	B int `msg:"b"`
}

type Reply struct {
	C int `msg:"c"`
}

type Arith int

func (t *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	fmt.Println("this is Arith")
	reply.C = args.A * args.B
	fmt.Println(reply.C, " = ", args.A, " * ", args.B)
	return nil
}

type Arith2 int

func (t *Arith2) Mul(ctx context.Context, args *Args, reply *Reply) error {
	fmt.Println("this is Arith2")
	reply.C = args.A * args.B
	fmt.Println(reply.C, " = ", args.A, " * ", args.B)
	return nil
}

func main() {
	go func() {
		server := server.NewServer()
		server.RegisterName("Arith", new(Arith), "")
		server.Serve("tcp", "127.0.0.1:8972")
	}()

	go func() {
		server := server.NewServer()
		// group is a meta data. If you set group metadata for some services, only clients in this group can access those services.
		server.RegisterName("Arith", new(Arith2), "group=test")
		server.Serve("tcp", "127.0.0.1:8973")
	}()

	select {}
}
