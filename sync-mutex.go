package main

import (
  "fmt"
  "sync"
  "time"
)

func main() {
  var a = 0
  var b = 0


  // 启动多个协程，需要足够大, 若没有锁则可能出现最大值不是循环次数的问题
  // 一个互斥锁只能同时被一个 goroutine 锁定，其它 goroutine 将阻塞直到互斥锁被解锁（重新争抢对互斥锁的锁定）。
  // 当代码执行到有锁的地方时，它获取不到互斥锁的锁定，会阻塞在那里，从而达到控制同步的目的。
  //
  var lock sync.Mutex
  for i := 0; i < 10000; i++ {
    go func(idx int) {
      lock.Lock()
      defer lock.Unlock()
      a += 1
      if (a > b) { b = a; }
    }(i)
  }

  // 等待 1s 结束主程序
  // 确保所有协程执行完
  time.Sleep(time.Second)
  fmt.Printf("max value b=%d\n", b)
}

