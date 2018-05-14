package main

import (
	"fmt"
	"time"
)

// 不定参数，语法糖，只能作为参数最后一个
func sum(resultChan chan int, values ...int) {
	sum := 0
	for _, value := range values {
		time.Sleep(time.Second)
		sum += value
	}
	resultChan <- sum //将计算结果·发送到channel中
}

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(values[2:4]) // （2<= index < 4）
	resultChan := make(chan int, 2)
	go sum(resultChan, values[:len(values)/2]...)
	go sum(resultChan, values[len(values)/2:]...)
	sum1, sum2 := <-resultChan, <-resultChan // 接收结果
	fmt.Println("Sum1:", sum1)
	fmt.Println("Sum2:", sum2)
	fmt.Println("Sum:", sum1+sum2)
}

// ch := make(chan interface{}) 和 ch := make(chan interface{},1)是不一样的
// 无缓冲的 不仅仅是只能向 ch 通道放 一个值 而是一直要有人接收，那么ch <- elem才会继续下去，要不然就一直阻塞着，也就是说有接收者才去放，没有接收者就阻塞。
// 而缓冲为1则即使没有接收者也不会阻塞，因为缓冲大小是1只有当 放第二个值的时候 第一个还没被人拿走，这时候才会阻塞 
