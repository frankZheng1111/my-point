package main

//https://gocn.vip/article/470

import (
	"log"
	"runtime"
	"sync"
	// "time"
	"fmt"
)

var done bool = false
var num int

func wait(wg *sync.WaitGroup, id int) {
	if id > 0 {
		log.Println("enter wait func", id)
	}
	num += 1 - id
	num += 2 - id
	if id > 0 {
		for i := 0; i < 60000; i++ {
			fmt.Printf("p")
		}
	}
	num += 3
	num += 4
	num += 5
	log.Println("num in ", id, " : ", num)
	// for !done {
	// 	// 个人理解: go关键字生成一个goroutine并竞争资源，若竞争到资源后再该关键字执行完之前不会让出资源
	// 	// 但是一些特定操作会临时让出资源
	// 	// log.Printf("notdone! %d\n", id)
	// 	// time.Sleep(time.Nanosecond)
	runtime.Gosched()
	// }
	wg.Done()
}

func main() {

	runtime.GOMAXPROCS(1)
	log.Println(runtime.GOMAXPROCS(-1))

	// go func() {
	// 	log.Println("setdone!") //
	// 	done = true
	// }()

	wg := sync.WaitGroup{}
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go wait(&wg, i)
	}

	// for !done {
	// 	runtime.Gosched() // 若在GOMAXPROCS=1的情况下移除此句，会无限pending
	// }

	wg.Wait()
	log.Println("done!")
}
