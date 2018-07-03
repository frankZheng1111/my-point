package main

import (
	"fmt"
	"runtime"
	"sync"
)

var done bool = false

func wait(wg *sync.WaitGroup, id int) {
	for !done {
		fmt.Println("notdone!", id) // 个人理解，每执行一个函数后分配资源
		// runtime.Gosched()
	}
	wg.Done()
}

func main() {

	runtime.GOMAXPROCS(1)
	fmt.Println(runtime.GOMAXPROCS(-1))

	go func() {
		done = true
	}()

	wg := sync.WaitGroup{}
	wg.Add(4)
	for i := 0; i < 4; i++ {
		go wait(&wg, i)
	}

	// for !done {
	// 	runtime.Gosched() // 若在GOMAXPROCS=1的情况下移除此句，会无限pending
	// }

	wg.Wait()
	fmt.Println("done!")
}
