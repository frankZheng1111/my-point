package main

// https://halfrost.com/go_slice/

import "fmt"

func main() {

	var nilSlice []int           // nil
	fmt.Println(nilSlice == nil) //true
	// fmt.Println(len(nil))        // panic: use of untyped nil
	fmt.Println(len(nilSlice), cap(nilSlice)) // 0, 0
	nilSlice = append(nilSlice, 1)            // 不会panic
	fmt.Println(len(nilSlice), cap(nilSlice)) // 1, 1
	// makeSlice := make([]int)//  panic: missing len argument to make([]int)
	makeSlice := make([]int, 1)
	fmt.Println(len(makeSlice), cap(makeSlice)) // 1, 1
	initSlice := []int{}
	fmt.Println(initSlice == nil)               // false
	fmt.Println(len(initSlice), cap(initSlice)) // 0, 0

	/* 创建切片 */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)

	// 切片保存了对底层数组的引用，若你将某个切片赋予另一个切片，它们会引用同一个数组。
	// 若某个函数将一个切片作为参数传入，则它对该切片元素的修改对调用者而言同样可见， 这
	// 可以理解为传递了底层数组的指针。
	updateSlice(number3)
	fmt.Println("src number3[1]=", number3[1])

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func updateSlice(x []int) {
	x[1] = 1000
}
