// go run -tags zookeeper rpcxZookeeperClient.go
package main

import (
	"context"
	"flag"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/share"
	"log"
	// "time"
)

var (
	zkAddr   = flag.String("zkAddr", "localhost:2181", "zookeeper address")
	basePath = flag.String("base", "/rpcx_test", "prefix path")
)

func main() {
	flag.Parse()

	d := client.NewZookeeperDiscovery(*basePath, "Arith", []string{*zkAddr}, nil)
	option := client.DefaultOption
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, option)
	defer xclient.Close()

	args := &struct{ A, B int }{
		A: 10,
		B: 20,
	}

	// zookeeper 不影响调用
	// log.Println("====================")
	// time.Sleep(15 * time.Second)

	var replyStr string
	ctx := context.WithValue(context.Background(), share.ReqMetaDataKey, map[string]string{"reqMeta": "FromClient"})
	ctx = context.WithValue(ctx, share.ResMetaDataKey, make(map[string]string))
	err := xclient.Call(ctx, "PrintMetaData", "reqStr", &replyStr)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("同步调用: PrintMetaData: reply: %s, metaData: %s \n", replyStr, ctx.Value(share.ResMetaDataKey).(map[string]string)["resMeta"])

	// 同步调用
	reply := &struct{ C int }{}
	err = xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("同步调用: %d * %d = %d \n", args.A, args.B, reply.C)

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
