package main

import "fmt"

func main() {
	var data interface{} = "great"

	// 如果只收一个参数断言, 则断言失败时会panic
	// 必须精准断言, 若是int32 断言int8会失败
	if res, ok := data.(bool); ok {
		fmt.Println("[is an int] data =>", data, " res => ", res)
	} else {
		fmt.Println("[not an int] value =>", data, "res => ", res)
		//prints: [not an int] value => 0 (not "great")
	}
}
