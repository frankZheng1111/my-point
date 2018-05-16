package main

import "fmt"

func main() {
	// comment1
	fmt.Println("Hello, World!")
	/*
	 * comment2
	 * comment3
	 */
	fmt.Println("Hello, World! Again!")
	ret, err := f1(1)
	fmt.Printf("ret is %s, err is %s\n", ret, err)
	// 左边至少一个未定义的变量
	ret2, err := f1(0)
	fmt.Printf("ret2 is %s, err is %s\n", ret2, err)
}

func f1(a int) (result string, errorMsg string) {
	if a > 0 {
		return "ok", ""
	}
	return "fail", "a <= 0"
}
