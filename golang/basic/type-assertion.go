package main

import "fmt"

func main() {
	var data interface{} = "great"

	if res, ok := data.(bool); ok {
		fmt.Println("[is an int] data =>", data, " res => ", res)
	} else {
		fmt.Println("[not an int] value =>", data, "res => ", res)
		//prints: [not an int] value => 0 (not "great")
	}
}
