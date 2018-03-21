package main

import "fmt"
import "os"
import "log"

type ByteSize float64

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

// 每个源文件都可以通过定义自己的无参数 init 函数来设置一些必要的状态。 （其实每
// 个文件都可以拥有多个 init 函数。）而它的结束就意味着初始化结束： 只有该包中的所有变
// 量声明都通过它们的初始化器求值后 init 才会被调用， 而那些 init 只有在所有已导入的包都
// 被初始化后才会被求值。
func init() {
	fmt.Println("Start init1...")
	if user == "" {
		log.Fatal("$USER not set")
	}
	if home == "" {
		home = "/home/" + user
	}
	if gopath == "" {
		gopath = home + "/go"
	}
}

// 变量
// 变量的初始化与常量类似，但其初始值也可以是在运行时才被计算的一般表达式。
var (
	home   = os.Getenv("HOME")
	user   = os.Getenv("USER")
	gopath = os.Getenv("GOPATH")
)

func init() {
	fmt.Println("Start init2...")
}

func main() {
	fmt.Println("Start main func")
	fmt.Println(MB, GB, TB, PB, EB)
	fmt.Println(home, user, gopath)
}
