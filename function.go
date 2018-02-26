package main

import "fmt"

func main() {
  /* 定义局部变量 */
  a := 100
  b := 200

  /* 调用函数并返回最大值 */
  maxValue := max(a, b)

  fmt.Println( "最大值是 : %d", maxValue )
  c, d := swapNumber(a, b)
  fmt.Println( "输入a = %d, 输入b = %d, 输出 a = %d, b = %d", a, b, c, d )
}

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
  /* 定义局部变量 */
  var result int

  if (num1 > num2) {
    result = num1
  } else {
    result = num2
  }
  return result 
}

/* 函数返回两个数的最大值 */
func swapNumber(num1, num2 int) (int, int) {
  return num2, num1
}
