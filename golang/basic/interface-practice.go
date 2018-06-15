package main

import "fmt"

type S struct {
	i int
}

func (p *S) Get() int {
	return p.i
}
func (p *S) Put(v int) {
	p.i = v
}

type I interface {
	Get() int
	Put(int)
}

func iSet888(p I) {
	fmt.Println(p.Get())
	p.Put(888)
}

func f2(p interface{}) {
	fmt.Println("")
	switch t := p.(type) {
	case int:
		fmt.Println("this is int number")
	case I:
		fmt.Println("I.i:", t.Get())
	default:
		fmt.Println("unknow type")
	}
}

// 如果你不想使用switch，可以用下面的笨方法。
// func g(something interface{}){
//   if t,ok := something.(I); ok{
//     fmt.Println("I:",t.Get())
//   }else if t,ok := something.(int); ok{
//     fmt.Println("int:",t)
//   }else{
//     fmt.Println("not found:",something)
//   }
// }

//指针修改原数据
func dd(a *S) {
	a.Put(999)
	fmt.Println(a.Get(), "in dd func")
}

//临时数据
func aa(a S) {
	a.Put(2222)
	fmt.Println(a.Get(), "in aa func")
}

func main() {
	var s S
	// 首先s是用S结构体创建的，S有Get() Put()两个方法。所以s可以执行Put()
	s.Put(333)
	fmt.Println("s.i = ", s.Get())
	// 因为S实现了I类型的接口，换句话说，S实现了I interface类型定义好的方法，那么I定义也就有了Get方法。
	iSet888(&s)
	fmt.Println(s.Get())
	// interface{}空接口可以是任何类型，我们可以在逻辑用断言的方式区别他是什么类型，然后根据类型做相应的处理。对应到上面的代码,
	// 我给你给他传任何值，f2因为是空间口都会接收进来。
	// 后面的 t := p.(type)是断言，所谓的断言就是区分他的type类型。
	f2(&s)
	dd(&s)
	fmt.Println(s.Get())
	aa(s)
	fmt.Println(s.Get())
}
