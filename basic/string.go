package main

import (
	"fmt"
)

func main() {
	var str string //声明一个字符串
	str = "string" //赋值

	ch := str[0] //获取第一个字符 string[index]获取的是字符byte
	fmt.Println("str[0] 输出", ch)
	fmt.Println("str[0] 输出", ch)

	strLen := len(str) //字符串的长度,len是内置函数 ,len=5
	fmt.Println("str 的长度是", strLen)

	str = "string字符串"

	for i := 0; i < len(str); i++ {
		fmt.Println(string(str[i]), ": ", str[i])
	}

	for i, s := range str {
		fmt.Println(i, "Unicode(", s, ") string=", string(s)) //一个汉字在UTF-8>中占3个字节
	}

	// rune 类型是 Unicode 字符类型，和 int32 类型等价，通常用于表示一个 Unicode 码点。rune 和 int32 可以互换使用。
	r := []rune(str)
	fmt.Println("rune=", r)
	for i := 0; i < len(r); i++ {
		fmt.Println("r[", i, "]=", r[i], "string=", string(r[i]))
	}
}
