package main

import (
	"fmt"
)

type Phone interface {
	call()
}

type NokiaPhone struct {
	name string
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main() {
	var phone Phone

	phone = new(NokiaPhone)
	// fmt.Println(phone.name) // hone.name undefined (type Phone has no field or method name)
	phone.call()

	phone = &IPhone{}
	phone.call()

}
