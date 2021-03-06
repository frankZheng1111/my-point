package main

import "fmt"

func main() {
	var n = [10]int{999, 998}      /* n 是一个长度为 10 的数组 */
	fmt.Println(n, len(n), cap(n)) // [999 998 0 0 0 0 0 0 0 0] 10 10

	/* 为数组 n 初始化元素 */
	for i := 3; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}

	/* 输出每个数组元素的值 */
	for j := 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}
