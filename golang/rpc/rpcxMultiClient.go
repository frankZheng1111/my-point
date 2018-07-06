package main

import (
	"context"
	"flag"
	"log"

	"github.com/smallnest/rpcx/client"
)

func main() {
	flag.Parse()

	d := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: "tcp@localhost:8972"}, {Key: "tcp@localhost:8973", Value: "group=test"}})
	// Discovery can find the group. Client can use `option.Group` to set group.
	// If you have not set `option.Group`, clients can access any services whether services set group or not.
	option := client.DefaultOption
	// option.Group = "test" // will only access Arith2 after setting group
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, option)
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
