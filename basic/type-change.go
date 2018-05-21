package main

import "fmt"

/*
Conversions are expressions of the form T(x) where T is a type and x is an expression that can be converted to type T.

...

A non-constant value x can be converted to type T in any of these cases:

x is assignable to T.
x's type and T have identical underlying types.
x's type and T are unnamed pointer types and their pointer base types have identical underlying types.
x's type and T are both integer or floating point types.
x's type and T are both complex types.
x is an integer or a slice of bytes or runes and T is a string type.
x is a string and T is a slice of bytes or runes.
*/
func main() {
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum) / float32(count)
	fmt.Printf("mean 的值为: %f\n", mean)
}
