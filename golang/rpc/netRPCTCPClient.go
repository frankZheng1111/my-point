package main

import (
	"fmt"
	"net/rpc"
	// "time"
)

type Quotient struct {
	Quo, Rem int
}

func main() {
	var client, err = rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("连接不到服务器：", err)
	}
	var args = struct{ A, B int }{40, 3}
	var result int
	fmt.Println("开始调用！")
	err = client.Call("Arith.Multiply", args, &result) // 同步调用
	if err != nil {
		fmt.Println("调用失败！", err)
	}
	fmt.Println("Arith: ", args.A, " * ", args.B, " = ", result)

	quotient := new(Quotient)
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	// time.Sleep(time.Second * 1)
	replyCall := <-divCall.Done //异步调用
	fmt.Println("Arith: ", args.A, " / ", args.B, " = ", quotient, replyCall)
}
