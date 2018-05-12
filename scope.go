package main

import "fmt"

var mainProfile string

func main() {
	subProfile := "sub"
	if condition := "condition"; 3 == (2 + 1) {
		fmt.Println("conditionPrint: ", condition)
	}
	fmt.Println("subPrint ", subProfile)
}
