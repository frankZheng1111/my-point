package main

import "fmt"

// 正如 ByteSize 那样，我们可以为任何已命名的类型（除了指针或接口）定义方法； 接收者可
// 不必为结构体。
//
type ByteSlice []byte

// 在之前讨论切片时，我们编写了一个 Append 函数。 我们也可将其定义为切片的方法。为
// 此，我们首先要声明一个已命名的类型来绑定该方法， 然后使该方法的接收者成为该类型的
// 值。
func (slice ByteSlice) Append(data []byte) []byte {
	l := len(slice)
	if l+len(data) > cap(slice) { // 重新分配
		// 为了后面的增长，需分配两份。
		newSlice := make([]byte, (l+len(data))*2)
		// copy 函数是预声明的，且可用于任何切片类型。
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : l+len(data)]
	for i, c := range data {
		slice[l+i] = c
	}
	return slice
}

func (slice *ByteSlice) AppendSelf(data []byte) {
	*slice = (*slice).Append(data)
}

func main() {
	fmt.Println("Start main")
	byte1 := ByteSlice{0, 1, 2}
	fmt.Println("byte1: ", byte1)
	fmt.Println("byte1 Append result", byte1.Append([]byte{3, 4}))
	fmt.Println("byte1: after Append ", byte1)
	byte1.AppendSelf([]byte{3, 4})
	fmt.Println("byte1: after AppendSelf ", byte1)
}
