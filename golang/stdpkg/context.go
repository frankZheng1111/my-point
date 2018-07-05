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
//
//   // Value方法获取该Context上绑定的值，是一个键值对
//   // 所以要通过一个Key才可以获取对应的值，这个值一般是线程安全的。
//   // 可以通过WithValue传递
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
	"sync"
	"time"
)

var key string = "key"
var wg sync.WaitGroup

func main() {
	// WithCancel函数，传递一个父Context作为参数，返回子Context，以及一个取消函数用来取消Context。
	// WithValue函数和取消Context无关，它是为了生成一个绑定了一个键值对数据的Context，这个绑定的数据可以通过Context.Value方法访问到，这是我们实际用经常要用到的技巧，一般我们想要通过上下文来传递数据时，可以通过这个方法，如我们需要tarce追踪系统调用栈的时候。
	//
	fmt.Println("Start show WithCancel & WithValue")
	ctx, cancel := context.WithCancel(context.Background())
	valueCtx := context.WithValue(ctx, key, "add value")

	wg.Add(1)

	go ContextWithCancelAndValue(valueCtx)
	time.Sleep(2 * time.Second) // 2秒后取消
	cancel()

	wg.Wait()

	fmt.Println("Start show WithTimeout(多少秒后超时)")

	wg.Add(1)

	ctx, cancel = context.WithTimeout(context.Background(), 4*time.Second)

	go ContextWithTimeout(ctx)

	wg.Wait()
}

func ContextWithTimeout(ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case t := <-time.After(1 * time.Second):
			fmt.Println("Doing some work ", t)

			// we received the signal of cancelation in this channel
		case <-ctx.Done():
			fmt.Println(ctx.Err()) //o: context deadline exceeded
			fmt.Println("Cancel the context ")
			return
		}
	}
}

func ContextWithCancelAndValue(ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done(): // struct{}
			//get value
			fmt.Println(ctx.Err()) // o: context canceled
			fmt.Println(ctx.Value(key), "is cancel")
			return
		default:
			//get value
			fmt.Println(ctx.Value(key), "int goroutine")
			time.Sleep(1 * time.Second)
		}
	}
}
