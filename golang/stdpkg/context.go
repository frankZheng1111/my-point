package main

// https://juejin.im/post/5a6873fef265da3e317e55b6

/*
golang在1.6.2的时候还没有自己的context，在1.7的版本中就把golang.org/x/net/context包被加入到了官方的库中。golang 的 Context包，是专门用来简化对于处理单个请求的多个goroutine之间与请求域的数据、取消信号、截止时间等相关操作，这些操作可能涉及多个 API 调用。
比如有一个网络请求Request，每个Request都需要开启一个goroutine做一些事情，这些goroutine又可能会开启其他的goroutine。这样的话， 我们就可以通过Context，来跟踪这些goroutine，并且通过Context来控制他们的目的，这就是Go语言为我们提供的Context，中文可以称之为“上下文”。
*/

// type Context interface {
// 	Deadline() (deadline time.Time, ok bool)
// 	Done() <-chan struct{}
// 	Err() error
// 	Value(key interface{}) interface{}
// }

// 不要把Context放在结构体中，要以参数的方式传递，parent Context一般为Background
// 应该要把Context作为第一个参数传递给入口请求和出口请求链路上的每一个函数，放在第一位，变量名建议都统一，如ctx。
// 给一个函数方法传递Context的时候，不要传递nil，否则在tarce追踪的时候，就会断了连接
// Context的Value相关方法应该传递必须的数据，不要什么数据都使用这个传递
// Context是线程安全的，可以放心的在多个goroutine中传递
// 可以把一个 Context 对象传递给任意个数的 gorotuine，对它执行 取消 操作时，所有 goroutine 都会接收到取消信号。

//通过 context.WithValue 来传值

import (
	"context"
	"fmt"
	"time"
)

var key = "key"

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	//通过 context.WithValue 来传值
	valueCtx := context.WithValue(ctx, key, "add value")
	go watch(valueCtx)
	time.Sleep(10 * time.Second)
	cancel()
	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context) {
	for {
		select {
		case err := <-ctx.Done():
			//get value
			fmt.Println(ctx.Err(), err)
			fmt.Println(ctx.Value(key), "is cancel")
			return
		default:
			//get value
			fmt.Println(ctx.Value(key), "int goroutine")
			time.Sleep(2 * time.Second)
		}
	}
}
