package main

import "fmt"

func main() {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("c")
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
		}
		fmt.Println("d")
	}()
	f()
}

func f() {
	fmt.Println("a")
	causePanic()
	fmt.Println("b")
	fmt.Println("f")
}

func causePanic() {
	panic(55)
}

// panic 是用来表示非常严重的不可恢复的错误的。在Go语言中这是一个内置函数，接收一个interface{}类型的值（也就是任何值了）作为参数。panic的作用就像我们平常接触的异常。不过Go可没有try…catch，所以，panic一般会导致程序挂掉（除非recover）。所以，Go语言中的异常，那真的是异常了。你可以试试，调用panic看看，程序立马挂掉，然后Go运行时会打印出调用栈。
// 但是，关键的一点是，即使函数执行的时候panic了，函数不往下走了，运行时并不是立刻向上传递panic，而是到defer那，等defer的东西都跑完了，panic再向上传递。所以这时候 defer 有点类似 try-catch-finally 中的 finally。
// panic就是这么简单。抛出个真正意义上的异常。

// 上面说到，panic的函数并不会立刻返回，而是先defer，再返回。这时候（defer的时候），如果有办法将panic捕获到，并阻止panic传递，那就异常的处理机制就完善了。
//
// Go语言提供了recover内置函数，前面提到，一旦panic，逻辑就会走到defer那，那我们就在defer那等着，调用recover函数将会捕获到当前的panic（如果有的话），被捕获到的panic就不会向上传递了，于是，世界恢复了和平。你可以干你想干的事情了。
//
// 不过要注意的是，recover之后，逻辑并不会恢复到panic那个点去，函数还是会在defer之后返回。

// 关于何时返回error, 何时使用panic, 个人开发时倾向的做法是，在封装一个服务(函数，工具，api等等), 在同时满足以下所有条件时使用panic：
//     1. 当前服务申明的行为无法进行处理(或者说当前的服务完成其申明的行为)
//     2. 出现状况的原因与使用该服务的方式无关(即出现的该状况不是因为的错误的调用或者传入错误的参数引起的)
// 若不能满足条件则返回error

// 至于团队合作时, error和panic的使用时机需要和团队成员取得一致(因为未recover的话panic会使整个程序挂掉)
