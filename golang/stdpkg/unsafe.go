package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// unsafe.Pointer && uintptr类型
//（1）任何类型的指针都可以被转化为Pointer
//（2）Pointer可以被转化为任何类型的指针
//（3）uintptr可以被转化为Pointer
//（4）Pointer可以被转化为uintptr
//
func PointerAnduintptr() {
	v1 := uint(12)
	v2 := int(13)

	fmt.Println("type of uint(12): ", reflect.TypeOf(v1))   //uint
	fmt.Println("type of int(12): ", reflect.TypeOf(v2))    //int
	fmt.Println("type of &uint(12): ", reflect.TypeOf(&v1)) //*uint
	fmt.Println("type of &int(12): ", reflect.TypeOf(&v2))  //*int

	p := &v1

	//两个变量的类型不同,不能赋值
	// p = &v2 //cannot use &v2 (type *int) as type *uint in assignment

	p = (*uint)(unsafe.Pointer(&v2)) //使用unsafe.Pointer进行类型的转换

	fmt.Println("p = ", *p)
}

func main() {
	PointerAnduintptr()
}
