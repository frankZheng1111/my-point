package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// unsafe包用于Go编译器，而不是Go运行时。
// 使用unsafe作为程序包名称只是让你在使用此包是更加小心。
// 使用unsafe.Pointer并不总是一个坏主意，有时我们必须使用它。
// Golang的类型系统是为了安全和效率而设计的。 但是在Go类型系统中，安全性比效率更重要。 通常Go是高效的，但有时安全真的会导致Go程序效率低下。 unsafe包用于有经验的程序员通过安全地绕过Go类型系统的安全性来消除这些低效。
// unsafe包可能被滥用并且是危险的。

// unsafe.Pointer && uintptr类型
//（1）任何类型的指针都可以被转化为Pointer
//（2）Pointer可以被转化为任何类型的指针
//（3）uintptr可以被转化为Pointer
//（4）Pointer可以被转化为uintptr
//
//直到现在（Go1.7），unsafe包含以下资源：

// 三个函数在编译时即可求值, 故能赋值给常量：
// func Alignof（variable ArbitraryType）uintptr
// func Offsetof（selector ArbitraryType）uintptr
// func Sizeof（variable ArbitraryType）uintptr
// （BTW，unsafe包中的函数中非唯一调用将在编译时求值。当传递给len和cap的参数是一个数组值时，内置函数和cap函数的调用也可以在编译时被求值。）
func PkgFunction() {
	fmt.Println("\nRUN PkgFunction;")
	var boolean bool
	var integer64 int64
	const BOOL_TYPE_SIZE = unsafe.Sizeof(boolean)
	const INT64_TYPE_SIZE = unsafe.Sizeof(integer64)
	fmt.Println("sizeof bool type: ", BOOL_TYPE_SIZE)   // 1个字节
	fmt.Println("sizeof int64 type: ", INT64_TYPE_SIZE) // 8个字节
}

func PointerAnduintptr() {
	fmt.Println("\nRUN PointerAnduintptr;")
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
	PkgFunction()
	PointerAnduintptr()
}
