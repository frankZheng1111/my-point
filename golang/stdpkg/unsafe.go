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

// 合法用例2: 调用sync/atomic包中指针相关的函数
// sync / atomic包中的以下函数的大多数参数和结果类型都是unsafe.Pointer或*unsafe.Pointer：
//
// func CompareAndSwapPointer（addr * unsafe.Pointer，old，new unsafe.Pointer）（swapped bool）
// func LoadPointer（addr * unsafe.Pointer）（val unsafe.Pointer）
// func StorePointer（addr * unsafe.Pointer，val unsafe.Pointer）
// func SwapPointer（addr * unsafe.Pointer，new unsafe.Pointer）（old unsafe.Pointer）
// atomic包可以完成一些原子级操作比如对比后交换值之类的,个人理解就是相当于封装了锁操作，可能有性能上的优化

var data = struct {
	onebyte   byte
	boolean   bool
	integer64 int64
	str       string
}{str: ""}

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

func PkgFunctionSizeof() {
	// bool 类型虽然只有一位，但也需要占用1个字节，因为计算机是以字节为单位
	// 64为的机器，一个 int 占8个字节
	// string 类型占16个字节，内部包含一个指向数据的指针（8个字节）和一个 int 的长度（8个字节）
	// slice 类型占24个字节，内部包含一个指向数据的指针（8个字节）和一个 int 的长度（8个字节）和一个 int 的容量（8个字节）
	// map 类型占8个字节，是一个指向 map 结构的指针
	// chan 类型占4个字节
	// 可以用 struct{} 表示空类型，这个类型不占用任何空间，用这个作为 map 的 value，可以讲 map 当做 set 来用
	fmt.Println("\nRUN PkgFunctionSizeof;")
	const STRUCT_TYPE_SIZE = unsafe.Sizeof(data)
	const BOOL_TYPE_SIZE = unsafe.Sizeof(data.boolean)
	const BYTE_TYPE_SIZE = unsafe.Sizeof(data.onebyte)
	const INT64_TYPE_SIZE = unsafe.Sizeof(data.integer64)
	const STRING_TYPE_SIZE = unsafe.Sizeof(data.str)
	fmt.Println("sizeof byte type: ", BYTE_TYPE_SIZE)     // 1个字节
	fmt.Println("sizeof bool type: ", BOOL_TYPE_SIZE)     // 1个字节
	fmt.Println("sizeof int64 type: ", INT64_TYPE_SIZE)   // 8个字节
	fmt.Println("sizeof string type: ", STRING_TYPE_SIZE) // 16个字节
	fmt.Println("sizeof struct type: ", STRUCT_TYPE_SIZE) // 32个字节 = |x(byte)x(bool)------|xxxxxxxx(int64)|xxxxxxxx(str.ptr)|xxxxxxxx(str.len)|
}

// Alignof 函数返回对应参数的类型对齐所需要的"倍数"
// 什么是对齐: 结构体中的各个字段在内存中并不是紧凑排列的，而是按照字节对齐的，比如 int 占8个字节，那么就只能写在地址为8的倍数的地址处，至于为什么要字节对齐，主要是为了效率考虑
// 在struct中，它的对齐值是它的成员中的最大对齐值。
func PkgFunctionAlignof() {
	fmt.Println("\nRUN PkgFunctionAlignof;")
	const STRUCT_TYPE_ALIGN = unsafe.Alignof(data)
	const BOOL_TYPE_ALIGN = unsafe.Alignof(data.boolean)
	const BYTE_TYPE_ALIGN = unsafe.Alignof(data.onebyte)
	const INT64_TYPE_ALIGN = unsafe.Alignof(data.integer64)
	const STRING_TYPE_ALIGN = unsafe.Alignof(data.str)
	fmt.Println("alignof byte type: ", BYTE_TYPE_ALIGN)     // 1
	fmt.Println("alignof bool type: ", BOOL_TYPE_ALIGN)     // 1
	fmt.Println("alignof int64 type: ", INT64_TYPE_ALIGN)   // 8
	fmt.Println("alignof string type: ", STRING_TYPE_ALIGN) // 8
	fmt.Println("alignof struct type: ", STRUCT_TYPE_ALIGN) // 8

	fmt.Println("size12 内存分布 int8|int32|int16 = x---|xxxx|xx--", unsafe.Sizeof(size12)) // 12
	fmt.Println("size8 内存分布 int8|int16|int32 = x-xx|xxxx", unsafe.Sizeof(size8))        // 8
}

// unsafe.Offsetof 函數的參數必鬚是一個字段 x.f, 然後返迴 f 字段相對於 x 起始地址的偏移量, 包括可能的空洞.
// unsafe.Offsetof 函数的参数必须是一个结构体的字段，然后返回该字段相对于该结构体起始地址的偏移量
func PkgFunctionOffestof() {
	fmt.Println("\nRUN PkgFunctionOffsetof;")
	const BOOL_TYPE_OFFSET = unsafe.Offsetof(data.boolean)
	const BYTE_TYPE_OFFSET = unsafe.Offsetof(data.onebyte)
	const INT64_TYPE_OFFSET = unsafe.Offsetof(data.integer64)
	const STRING_TYPE_OFFEST = unsafe.Offsetof(data.str)
	fmt.Println("Offsetof byte type: ", BYTE_TYPE_OFFSET)     // 0
	fmt.Println("Offsetof bool type: ", BOOL_TYPE_OFFSET)     // 1
	fmt.Println("Offsetof int64 type: ", INT64_TYPE_OFFSET)   // 8
	fmt.Println("Offsetof string type: ", STRING_TYPE_OFFEST) // 16
	fmt.Println("size12 内存分布 int8|int32|int16 = x---|xxxx|xx--",
		unsafe.Offsetof(size12.i8),  //0
		unsafe.Offsetof(size12.i32), //4
		unsafe.Offsetof(size12.i16), //8
	)
	fmt.Println("size8 内存分布 int8|int16|int32 = x-xx|xxxx",
		unsafe.Offsetof(size8.i8),  //0
		unsafe.Offsetof(size8.i16), //2
		unsafe.Offsetof(size8.i32), //4
	) // 8

}

// 类型Pointer * ArbitraryType
// 这里，ArbitraryType不是一个真正的类型，它只是一个占位符。
// type ArbitraryType int
// ArbitraryType仅用于文档目的，实际上并不是不安全包的一部分。 它表示任意Go表达式的类型。
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
	PkgFunctionSizeof()
	PkgFunctionAlignof()
	PkgFunctionOffestof()
	PointerAnduintptr()
}
