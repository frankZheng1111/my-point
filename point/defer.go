package main

import "fmt"

func main() {
	deferCall()
}

func deferCall() {
	defer func() { fmt.Println("defer1") }()
	defer func() { fmt.Println("defer2") }()
	defer func() { fmt.Println("defer3") }()
	panic("触发异常")
}

// 考察对defer的理解，defer函数属延迟执行，延迟到调用者函数执行 return 命令前被执行。多个defer之间按LIFO先进后出顺序执行。故考题中，在Panic触发时结束函数运行，在return前先依次打印:打印后、打印中、打印前 。最后由runtime运行时抛出打印panic异常信息。
//
