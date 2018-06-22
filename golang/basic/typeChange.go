/*
提要:

* 设有类型为T, 变量为x, T(x) 即将返回值为x类型为T的值
...
以下情况可以进行类型转换

c1: x的值本身就可以赋值给类型T的变量
c2: 两者有相同的底层类型(如int8 int16 int32 int64)
c3: x类型和T 都是未命名(未命名的指的是未原生命名，int，int32属于named type, []int, type special， struct 属于unnamedType, 只有unnamed类型可以作为方法的接受者)的指针类型，它们的指针指向的对象类型的底层类型一致
c4: 整型和浮点型可以互相转换
c5: x的类型和T都是复数类型
c6: x是一个整数(按Ascll码转)或(byte或Rune)的slice可转成字符串
c7: 上一条逆向转
...
*/
package main

import "fmt"

func case1() {
	// case1
	var sum int = 1000
	var count int = 3
	var fValue float32
	fValue = float32(sum) / float32(count)
	// 10为宽度，即在字符串中的占位宽度，1为小数点后位数
	fmt.Printf("fValue的值为: %10.1f\n", fValue)
}

func case6() {
	var num int = 90
	var numStr string
	numStr = string(num)
	fmt.Println("numStr转换成功", numStr)
}

type T1 *int
type T2 *int

func case3() {
	var num int = 1
	var numPtr1 T1 = &num
	var numPtr2 T2
	numPtr2 = T2(numPtr1)
	fmt.Println("numPtr转换成功", numPtr2)
}

func main() {
	case1()
	case3()
	case6()
}

/* 原文
A non-constant value x can be converted to type T in any of these cases:

x is assignable to T.
x's type and T have identical underlying types.
x's type and T are unnamed pointer types and their pointer base types have identical underlying types.
x's type and T are both integer or floating point types.
x's type and T are both complex types.
x is an integer or a slice of bytes or runes and T is a string type.
x is a string and T is a slice of bytes or runes.
*/
