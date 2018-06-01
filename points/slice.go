package main

import "fmt"

func FuncSlice(s []int, t int) {
	s[0]++
	s = append(s, t)
	s[0]++
}
func main() {
	a := []int{0, 1, 2, 3}
	FuncSlice(a, 4)
	fmt.Println(a) // {1,1,2,3}
}
