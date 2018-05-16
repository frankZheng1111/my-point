package main

import (
  "fmt"
  "math/rand"
  "sync"
  "time"
)

func main() {
  numbers1 := []int{6, 2, 7, 7, 3, 8, 9, 10,45,31,34,567,1,0,7,23,1,45,33,6}
  QuickSortIteratively(numbers1)
  fmt.Println(numbers1)

  numbers2 := []int{6, 2, 7, 7, 3, 8, 9, 10,45,31,34,567,1,0,7,23,1,45,33,6}
  QuickSortRecursively(numbers2)
  fmt.Println(numbers2)
}
// 将切片中某一段的值(head <= index <= tail)按照某一个基准点，左右划分，并排序，并返回基准点的index
func partition (values []int, head int, tail int, resultChan chan int) {
  time.Sleep(time.Second/2) // 用于区分
  baseValIndex := rand.Intn(tail - head) + head
  // fmt.Println("head", head)
  // fmt.Println("tail", tail)
  baseVal, i := values[baseValIndex], head // i为遍历游标

  for head < tail {
    if (i == baseValIndex) {
      i ++
      continue
    }
    if values[i] > baseVal {
      // 把比标准值大的位数放在最后, 同时左移尾标标记该外侧值经过确认
      values[i], values[tail] = values[tail], values[i]
      tail--
    } else {
      // 把比标准值大的位数放在最后, 同时右移头标标记该外侧值经过确认
      values[i], values[head] = values[head], values[i]
      head++
      i++
    }
  }
  resultChan <- head
}

// 非递归实现
// 将需要排序的区域通过栈记录
// 达到类似递归效果
func QuickSortIteratively(values []int) {
  length := len(values)
  if length <= 1 { return }
  var indexsStack = []struct{ head, tail int }{ { 0, length - 1 } }
  wg := sync.WaitGroup{}
  for len(indexsStack) > 0 {
    wg.Add(len(indexsStack))
    // fmt.Println("add", indexsStack)
    for len(indexsStack) > 0 {
      var indexs struct{ head, tail int }
      indexs, indexsStack = indexsStack[len(indexsStack) - 1], indexsStack[:(len(indexsStack) - 1)]
      go func (indexs struct{ head, tail int }) {
        resultChan := make(chan int, 1)
        partition(values, indexs.head, indexs.tail, resultChan)
        partitionIndex := <-resultChan
        if partitionIndex - indexs.head > 1 {
          indexsStack = append(indexsStack, struct{ head, tail int }{ indexs.head, partitionIndex - 1 })
          // fmt.Println("left")
        }
        if indexs.tail - partitionIndex > 1 {
          indexsStack = append(indexsStack, struct{ head, tail int }{ partitionIndex + 1, indexs.tail })
          // fmt.Println("right")
        }
        wg.Done()
      }(indexs)
    }
    wg.Wait()
  }
}

// 快排递归实现
func QuickSortRecursively(values []int) {
  wg := sync.WaitGroup{}
  wg.Add(1)
  var recursiveFunc func(values []int)
  recursiveFunc = func(values []int) {
    length := len(values)
    if length <= 1 {
      wg.Done()
      return
    }
    head, tail := 0, length-1 // 头尾坐标外侧的值都代表已经过标准值的划分
    partitionChan := make(chan int, 1)
    partition(values, head, tail, partitionChan)
    head = <-partitionChan
    wg.Add(2)
    go recursiveFunc(values[:head])
    go recursiveFunc(values[head+1:])
    wg.Done()
  }

  go recursiveFunc(values)
  wg.Wait()

}
