package main

import "fmt"

func main() {
	type Student struct {
		Name string
		Age  int
	}
	students := []Student{
		{Name: "a", Age: 5},
		{Name: "d", Age: 51},
		{Name: "b", Age: 2},
		{Name: "c", Age: 6},
	}
	errorMap := make(map[string]*Student)
	correctMap := make(map[string]*Student)
	for i, student := range students {
		fmt.Printf("当前循环student的地址是: %p\n", &student)
		// 错误写法, range中是students元素的副本使用的是同一个地址
		errorMap[student.Name] = &student

		//正确写法
		correctMap[students[i].Name] = &students[i]
	}

	fmt.Println("错误结果: ")
	for k, v := range errorMap {
		fmt.Println(k, "=>", v.Name)
	}

	fmt.Println("正确结果1: ")
	for k, v := range correctMap {
		fmt.Println(k, "=>", v.Name)
	}

	// 正确写法2
	studentPtrs := []*Student{
		{Name: "a", Age: 5},
		{Name: "d", Age: 51},
		{Name: "b", Age: 2},
		{Name: "c", Age: 6},
	}

	correctMap2 := make(map[string]*Student)
	for _, student := range studentPtrs {
		correctMap2[student.Name] = student
	}
	fmt.Println("正确结果2: ")
	for k, v := range correctMap2 {
		fmt.Println(k, "=>", v.Name)
	}
}
