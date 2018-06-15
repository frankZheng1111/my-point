package main

import (
	"fmt"
	"runtime"
	"sync"
)

func onceFunc() {
	fmt.Println("YOu will see this only once")
}

func main() {
	runtime.GOMAXPROCS(1)
	// wg := sync.WaitGroup{} // 初始化方式都行
	var wg sync.WaitGroup
	wg.Add(20)
	once := sync.Once{}
	for i := 0; i < 10; i++ {
		go func() {
			// 一旦一个Once对象的Do方法被调用，那么接下来对该Once对象Do方法的调用都将不会执行。
			once.Do(onceFunc)
			fmt.Println("i: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {

			fmt.Println("i2: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// 将GOMAXPROCS设置为1，将影响goroutine的并发，后续代码中的go func()相当于串行执行。
//
// 两个for循环内部go func 调用参数i的方式是不同的，后者闭包，导致结果完全不同。这也是新手容易遇到的坑。
//
// 第一个go func中i是外部for的一个变量，地址不变化。遍历完成后，最终i=10。故go func执行时，i的值始终是10（10次遍历很快完成）。
//
// 第二个go func中i是函数参数，与外部for中的i完全是两个变量。尾部(i)将发生值拷贝，go func内部指向值拷贝地址。
