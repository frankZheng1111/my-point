package main

import "fmt"

func main() {
	var a byte
	a = 11
	{
		fmt.Println(a) // 11
		a := 12
		fmt.Println(a) // 12
	}
	fmt.Println(a) // 11
}
