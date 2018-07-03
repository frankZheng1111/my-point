package main

// 小米三面Q1：
// 写个写方法和读方法, 实现读写锁
// 不用标准带的读写锁

import (
	"fmt"
	"runtime"
	"sync"
)

var count int

// 同时只能有一个 goroutine 能够获得写锁定。
// 同时可以有任意多个 gorouinte 获得读锁定。
// 同时只能存在写锁定或读锁定（读和写互斥）。
//
// 也就是说，当有一个 goroutine 获得写锁定，其它无论是读锁定还是写锁定都将阻塞直到写解锁；当有一个 goroutine 获得读锁定，其它读锁定任然可以继续；当有一个或任意多个读锁定，写锁定将等待所有读锁定解锁之后才能够进行写锁定。所以说这里的读锁定（RLock）目的其实是告诉写锁定：有很多人正在读取数据，你给我站一边去，等它们读（读解锁）完你再来写（写锁定）。
//
type CustomRWLock struct {
	BaseReadLock  sync.Mutex
	BaseWriteLock sync.Mutex
	Reader        int
	Writer        int
}

func (locker *CustomRWLock) RLock() {
	fmt.Println("取读锁")
	for {
		if locker.Writer <= 0 {
			break
		} else {
			runtime.Gosched()
		}
	}
	locker.BaseReadLock.Lock()
	defer locker.BaseReadLock.Unlock()
	locker.Reader++
	fmt.Println("取读锁成功")
}

func (locker *CustomRWLock) RUnlock() {
	fmt.Println("解读锁")
	locker.BaseReadLock.Lock()
	defer locker.BaseReadLock.Unlock()
	if locker.Reader <= 0 {
		panic("missing lock")
	}
	locker.Reader--
	fmt.Println("读锁成功")
}

func (locker *CustomRWLock) Lock() {
	fmt.Println("取写锁")
	for {
		if locker.Reader <= 0 && locker.Writer <= 0 {
			break
		} else {
			runtime.Gosched()
		}
	}
	locker.BaseWriteLock.Lock()
	defer locker.BaseWriteLock.Unlock()
	locker.Writer++
	fmt.Println("取写锁成功")
}

func (locker *CustomRWLock) Unlock() {
	fmt.Println("解写锁")
	locker.BaseWriteLock.Lock()
	defer locker.BaseWriteLock.Unlock()
	if locker.Writer <= 0 {
		panic("missing lock")
	}
	locker.Writer--
	fmt.Println("解写锁成功")
}

var wg sync.WaitGroup
var rw *CustomRWLock = new(CustomRWLock)

func main() {
	runtime.GOMAXPROCS(2)
	wg = sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 5; i++ {
		go read(i)
	}
	for i := 0; i < 5; i++ {
		go write(i)
	}
	wg.Wait()
}

func read(n int) {
	rw.RLock()
	fmt.Printf("goroutine %d 进入读操作...\n", n)
	v := count
	fmt.Printf("goroutine %d 读取结束，值为：%d\n", n, v)
	rw.RUnlock()
	wg.Done()
}

func write(n int) {
	rw.Lock()
	fmt.Printf("goroutine %d 进入写操作...\n", n)
	count++
	v := count
	fmt.Printf("goroutine %d 写入结束，新值为：%d\n", n, v)
	rw.Unlock()
	wg.Done()
}
