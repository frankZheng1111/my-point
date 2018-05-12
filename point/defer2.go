package main

import "fmt"

func calc(index string, a, b int) int {
  ret := a + b
  fmt.Println(index, a, b, ret)
  return ret
}
func main() {
  a := 1                                             //line 1
  b := 2                                             //2
  defer calc("1", a, calc("10", a, b))  //3
  a = 0                                              //4
  defer calc("2", a, calc("20", a, b))  //5
  b = 1                                              //6
}

// output
// 10 1 2 3
// 20 0 2 2
// 2 0 2 2
// 1 1 3 4

// defer是在函数末尾的return前执行，先进后执行.
// 函数调用时 int 参数发生值拷贝。

// 不管代码顺序如何，defer calc func中参数b必须先计算，故会在运行到第三行时，执行calc("10",a,b)输出：10 1 2 3得到值3，将cal("1",1,3)存放到延后执执行函数队列中。
// 执行到第五行时，现行计算calc("20", a, b)即calc("20", 0, 2)输出：20 0 2 2得到值2,将cal("2",0,2)存放到延后执行函数队列中。
// 执行到末尾行，按队列先进后出原则依次执行：cal("2",0,2)、cal("1",1,3) ，依次输出：2 0 2 2、1 1 3 4 。
//
