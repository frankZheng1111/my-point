package main

//https://gocn.vip/article/470
// https://github.com/k2huang/blogpost/blob/master/golang/%E5%B9%B6%E5%8F%91%E7%BC%96%E7%A8%8B/%E5%B9%B6%E5%8F%91%E6%9C%BA%E5%88%B6/Go%E5%B9%B6%E5%8F%91%E6%9C%BA%E5%88%B6.md
// https://juejin.im/entry/58d4ba87570c350058ca5ad8

import (
	"fmt"
	"log"
	"runtime"
	"sync"
)

var done bool = false
var num int

// 个人理解: go关键字生成一个goroutine并竞争资源，若竞争到资源后再该关键字后的函数执行完之前不会让出资源
// 一些特殊情况
// Go的运行时调度器中也有类似内核调度器的抢占机制，但并不能保证抢占能成功(调用函数的时候会检查flag)，因为Go运行时系统并没有内核调度器的中断能力，它只能通过向运行时间过长的G中设置抢占flag的方法温柔的让运行的G(goroutine)自己主动让出M(machine,指cpu)的执行权。
// 说到这里就不得不提一下Goroutine在运行过程中可以动态扩展自己线程栈的能力，可以从初始的2KB大小扩展到最大1G（64bit系统上），因此在每次调用函数之前需要先计算该函数调用需要的栈空间大小，然后按需扩展（超过最大值将导致运行时异常）。Go抢占式调度的机制就是利用在判断要不要扩栈的时候顺便查看以下自己的抢占flag，决定是否继续执行，还是让出自己。
// 运行时系统的监控线程会计时并设置抢占flag到运行时间过长的G，然后G在有函数调用的时候会检查该抢占flag，如果已设置就将自己放入全局队列，这样该M上关联的其他G就有机会执行了。但如果正在执行的G是个很耗时的操作且没有任何函数调用(如只是for循环中的计算操作)，即使抢占flag已经被设置，该G还是将一直霸占着当前M直到执行完自己的任务
func wait(wg *sync.WaitGroup, id int) {
	if id > 0 {
		log.Println("enter wait func", id)
	}
	num += 1 - id
	num += 2 - id
	if id > 0 {
		// runtime.Gosched()
		for i := 0; i < 60000; i++ {
			fmt.Printf("p")
		}
	}
	num += 3
	num += 4
	num += 5
	log.Println("num in ", id, " : ", num)
	wg.Done()
}

func main() {

	runtime.GOMAXPROCS(1)
	log.Println(runtime.GOMAXPROCS(-1))

	wg := sync.WaitGroup{}
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go wait(&wg, i)
	}
	wg.Wait()

	go func() {
		log.Println("setdone!")
		done = true
	}()

	for !done {
		// log.Println("notdone!") //在仅有此句的情况下若执行时间过长会被调度器设置flag并主动让出自己
		runtime.Gosched() // 若在GOMAXPROCS=1的情况下移除此句，会无限pending
	}

	log.Println("done!")
}
