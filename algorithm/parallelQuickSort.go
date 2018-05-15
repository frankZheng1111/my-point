package main

import (
  "fmt"
  "time"
)

func quickSort(values []int, resultChan chan []int) {
  time.Sleep(time.Second/2)
  if len(values) <= 1 {
    resultChan <- values
    return
  }
  baseVal := values[0]
  var smallValues []int
  var bigValues []int
  for num, value := range values {
    if num == 0 { continue }
    if value >= baseVal {
      bigValues = append(bigValues, value)
    } else {
      smallValues = append(smallValues, value)
    }
  }
  smallChan := make(chan []int, 1)
  bigChan := make(chan []int, 1)
  go quickSort(smallValues, smallChan)
  go quickSort(bigValues, bigChan)
  sortSmallValues := <-smallChan
  sortBigValues := <-bigChan
  sortValues := sortSmallValues
  sortValues = append(sortValues, baseVal)
  sortValues = append(sortValues, sortBigValues...)
  resultChan <- sortValues
}

func main() {
  resultChan := make(chan []int, 1)
  go quickSort([]int {3,2,3,4,3}, resultChan)
  fmt.Println(<-resultChan)
}

