package main

import (
	"context"
	"flag"
	"log"

	"github.com/smallnest/rpcx/client"
)

func main() {
	flag.Parse()

	d := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: "tcp@localhost:8972"}, {Key: "tcp@localhost:8973"}})
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &struct{ A, B int }{
		A: 10,
		B: 20,
	}

	for i := 0; i < 10; i++ {
		// 异步调用
		asyncReply := &struct{ C int }{}
		call, err := xclient.Go(context.Background(), "Mul", args, asyncReply, nil)
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}
		replyCall := <-call.Done
		if replyCall.Error != nil {
			log.Fatalf("failed to call: %v", replyCall.Error)
		} else {
			log.Printf("异步调用: %d * %d = %d, %v \n", args.A, args.B, asyncReply.C, replyCall)
		}
	}
}
