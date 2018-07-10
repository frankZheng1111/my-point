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
	p3 := unsafe.Pointer(uintptr(p1) + 2*unsafe.Sizeof(a[0]))

	*(*string)(p3) = "003"
	fmt.Println("slice a =", a) // slice a = [00 01 02 003]
}

func main() {
	VisitSliceWithAddress()
}
