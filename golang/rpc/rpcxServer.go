package main

import (
	"context"
	"fmt"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/share"
	"time"
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
	reply.C = args.A * args.B
	fmt.Println(reply.C, " = ", args.A, " * ", args.B)
	return nil
}

func (t *Arith) PrintMetaData(ctx context.Context, args string, reply *string) error {
	*reply = "success"
	fmt.Println("req.Meta: ", ctx.Value(share.ReqMetaDataKey).(map[string]string)["reqMeta"])
	resMeta := ctx.Value(share.ResMetaDataKey).(map[string]string)
	resMeta["resMeta"] = "FromServer"
	time.Sleep(3 * time.Second)
	return nil
}

func (t *Arith) Error(ctx context.Context, args *Args, reply *Reply) error {
	panic("ERROR")
}

func main() {
	server := server.NewServer()
	server.RegisterName("Arith", new(Arith), "")
	server.Serve("tcp", "127.0.0.1:8972")
}
