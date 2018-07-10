package main

import (
	"fmt"
	"unsafe"
)

// 这里有一些关于unsafe.Pointer和uintptr的事实：
//
// uintptr是一个整数类型。
// 即使uintptr变量仍然有效，由uintptr变量表示的地址处的数据也可能被GC回收。
// unsafe.Pointer是一个指针类型。
// 但是unsafe.Pointer值不能被取消引用。
// 如果unsafe.Pointer变量仍然有效，则由unsafe.Pointer变量表示的地址处的数据不会被GC回收。
// unsafe.Pointer是一个通用的指针类型，就像* int等。

// 关于指针类型转换
//对于将 T1转换为unsafe.Pointer，然后转换为 T2，unsafe包docs说：

// 如果T2比T1大，并且两者共享等效内存布局，则该转换允许将一种类型的数据重新解释为另一类型的数据。
// 这种“等效内存布局”的定义是有一些模糊的。 看起来go团队故意如此。 这使得使用unsafe包更危险。
//
// 由于Go团队不愿意在这里做出准确的定义，本文也不尝试这样做。

// 核心
// 由于uintptr是一个整数类型，uintptr值可以进行算术运算。 所以通过使用uintptr和unsafe.Pointer，我们可以绕过限制，* T值不能在Golang中计算偏移量：

var data = struct {
	onebyte   byte
	boolean   bool
	integer64 int64
	str       string
}{str: "data-string"}

var size12 = struct {
	i8  int8
	i32 int32
	i16 int16
}{}

var size8 = struct {
	i8  int8
	i16 int16
	i32 int32
}{}

func VisitSliceWithAddress() {
	fmt.Println("\n Run VisitSliceWithAddress")
	a := [4]string{"00", "01", "02", "03"}
	// 集合内的地址是连续的 故获取第n个元素的地址(头)后，加上m个地址长度, 则是可以获得第n
	p1 := unsafe.Pointer(&a[1])
	fmt.Println("sizeof uintptr: ", unsafe.Sizeof(uintptr(p1))) // sizeof uintptr:  8
	p3 := unsafe.Pointer(uintptr(p1) + 2*unsafe.Sizeof(a[0]))

	*(*string)(p3) = "003"
	fmt.Println("slice a =", a) // slice a = [00 01 02 003]
}

func VisitStringLengthWithAddress() {
	fmt.Println("\n Run VisitStringLengthWithAddress")
	str := "abcdefABCDEF"
	// 集合内的地址是连续的 故获取第n个元素的地址(头)后，加上m个地址长度, 则是可以获得第n
	strUnsafePtr := unsafe.Pointer(&str)
	// string 类型占16个字节，内部包含一个指向数据的指针（8个字节）和一个 int 的长度（8个字节）
	len := unsafe.Pointer(uintptr(strUnsafePtr) + unsafe.Sizeof(str)/2)
	strData := unsafe.Pointer(uintptr(strUnsafePtr))
	fmt.Println("string len by address =", *(*int)(len))         //string len by address = 6
	fmt.Println("string data by address =", *(*[]byte)(strData)) //string data by address = [97 98 99 100 101 102 65 66 67 68 69 70]
	fmt.Println(string([]byte{96}))                              // `

	*(*int)(len) = 3
	fmt.Println("string str =", str) // string str = abc
}

func main() {
	VisitSliceWithAddress()
	VisitStringLengthWithAddress()
}
