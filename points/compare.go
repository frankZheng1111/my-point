package main

import (
	"fmt"
	"reflect"
)

// You can use the equality operator, ==, to compare struct variables if each structure field can be compared with the equality operator.

type data struct {
	num     int
	fp      float32
	complex complex64
	str     string
	char    rune
	yes     bool
	events  <-chan string
	handler interface{}
	ref     *byte
	raw     [10]byte
}

type udata struct {
	num    int               //ok
	checks [10]func() bool   //not comparable
	doit   func() bool       //not comparable
	m      map[string]string //not comparable
	bytes  []byte            //not comparable
}

func main() {
	v1 := data{}
	v2 := data{}
	fmt.Println("v1 == v2: ", v1 == v2) // prints: v1 == v2: true

	v3 := udata{}
	v4 := udata{}
	// fmt.Println("v3 == v4:", v3 == v4) // invalid operation: v3 == v4 (struct containing [10]func() bool cannot be compared)

	v3.m = map[string]string{"one": "a", "two": "b"}
	v4.m = map[string]string{"two": "b", "one": "a"}
	fmt.Println("v3 == v4(with map): ", reflect.DeepEqual(v3, v4)) // true

	v4.bytes = []byte{'a', 'b'}
	v3.bytes = []byte{'b', 'a'}

	fmt.Println("v3 == v4(with slice): ", reflect.DeepEqual(v3, v4)) // false
}
