package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
)

//http://colobu.com/2016/09/18/go-net-rpc-guide/
// 如果对象的方法要能远程访问，它们必须满足一定的条件，否则这个对象的方法回呗忽略。
// 这些条件是：
// 方法的类型是可输出的 (the method's type is exported)
// 方法本身也是可输出的 （the method is exported）
// 方法必须由两个参数，必须是输出类型或者是内建类型 (the method has two arguments, both exported or builtin types)
// 方法的第二个参数是指针类型 (the method's second argument is a pointer)
// 方法返回类型为 error (the method has return type error)

// 第一步你需要定义传入参数和返回参数的数据结构：
//
type Args struct {
	A, B int
}

// 商: 除数结果
type Quotient struct {
	Quo, Rem int
}

// 第二步定义一个服务对象，这个服务对象可以很简单， 比如类型是int或者是interface{},重要的是它输出的方法。
// 这里我们定义一个算术类型Arith，其实它是一个int类型，但是这个int的值我们在后面方法的实现中也没用到，所以它基本上就起一个辅助的作用。

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
	var ms = new(Arith)
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
		rpc.ServeConn(conn)
	}
	// time.Sleep(3600 * time.Second)
}
