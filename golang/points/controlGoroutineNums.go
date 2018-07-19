package main

//控制goroutine的数量

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type GPool struct {
	queue chan struct{}
	wg    sync.WaitGroup
}

func NewGPool(size int) *GPool {
	if size <= 0 {
		size = 1
	}
	return &GPool{
		queue: make(chan struct{}, size),
		wg:    sync.WaitGroup{},
	}
}

func (p *GPool) Add(delta int) {
	for i := 0; i < delta; i++ {
		p.queue <- struct{}{}
	}
	p.wg.Add(delta)
}

func (p *GPool) Done() {
	<-p.queue
	p.wg.Done()
}

func (p *GPool) Wait() {
	p.wg.Wait()
}

func main() {
	pool := NewGPool(2)
	fmt.Println("循环前: ", runtime.NumGoroutine())
	for i := 0; i < 10; i++ {
		pool.Add(1)
		go func() {
			time.Sleep(time.Second)
			fmt.Println(runtime.NumGoroutine())
			pool.Done()
		}()
	}
	pool.Wait()
	fmt.Println("结束前", runtime.NumGoroutine())
}
