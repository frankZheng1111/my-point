package main

import "fmt"

type ByteSize float64

func main() {
  const (
    // 通过赋予空白标识符忽略第一个值
    _ = iota // ignore first value by assigning to blank identifier
    KB ByteSize = 1 << (10 * iota)
    MB
    GB
    TB
    PB
    EB
    // ZB
    // YB
  )
  fmt.Println(MB, GB, TB, PB, EB)
}
