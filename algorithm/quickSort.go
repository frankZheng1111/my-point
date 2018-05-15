package main

import (
  "fmt"
)

func quickSort(values []int) (sortValues []int) {
  if len(values) <= 1 { return values }
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
  if len(smallValues) > 0 { sortValues = append(sortValues, quickSort(smallValues)...) }
  sortValues = append(sortValues, baseVal)
  if len(bigValues) > 0 { sortValues = append(sortValues, quickSort(bigValues)...) }
  return
}

func main() {
  fmt.Println(quickSort([]int {3,2,3,4,3}))
}

