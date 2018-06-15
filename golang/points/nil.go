package main

// This is the second most common gotcha in Go because interfaces are not pointers even though they may look like pointers. Interface variables will be "nil" only when their type and value fields are "nil".

import "fmt"

func main() {
	var data *byte
	var in interface{}

	fmt.Println(data, " data == nil => ", data == nil) //prints: <nil> true
	fmt.Println(in, "in == nil => ", in == nil)        //prints: <nil> true

	in = data
	fmt.Println(in, in == nil) //prints: <nil> false
	//'data' is 'nil', but 'in' is not 'nil'
}
