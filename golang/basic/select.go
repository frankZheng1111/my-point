package main

import (
	"fmt"
	"time"
)

func main() {
	var c0 chan int
	c0 = make(chan int)
	// chan 若未初始化, 则会一直阻塞
	go func() {
		_ = <-c0
	}()
	c0 <- 2
	var c1 = make(chan int, 1) // 单个元素不超过65535个字节
	var c2 = make(chan int, 1)
	var c3 = make(chan int, 2)
	var i1 int
	var i3 int
	i1 = 1
	i2 := 2
	i3 = 3
	c3 <- i3
	c1 <- i1
	// close(c3)
	_, ok := <-c3
	fmt.Println(ok) // true
	_, ok = <-c3
	fmt.Println(ok) // false
	select {
	case a := <-time.After(5 * time.Second):
		fmt.Println("After 5 second", a)
	case i1 = <-c1:
		fmt.Println("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Println("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Println("received ", i3, " from c3\n")
		} else {
			fmt.Println("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}
