package main

import "fmt"
import "os"

type ByteSize float64

func main() {
	// 常量
	// Go 中的常量就是不变量。它们在编译时创建，即便它们可能是函数中定义的局部变量。 常量
	// 只能是数字、字符（符文）、字符串或布尔值。由于编译时的限制， 定义它们的表达式必须
	// 也是可被编译器求值的常量表达式。例如 1<<3 就是一个常量表达式，而
	// math.Sin(math.Pi/4) 则不是，因为对 math.Sin 的函数调用在运行时才会发生。
	const (
		// 通过赋予空白标识符忽略第一个值
		_           = iota // ignore first value by assigning to blank identifier
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

	// 变量
	// 变量的初始化与常量类似，但其初始值也可以是在运行时才被计算的一般表达式。
	var (
		home   = os.Getenv("HOME")
		user   = os.Getenv("USER")
		gopath = os.Getenv("GOPATH")
	)
	fmt.Println(home, user, gopath)
}
