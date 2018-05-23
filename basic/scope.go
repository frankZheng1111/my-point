package main

import (
	"fmt"
	"time"
)

var mainProfile string

func main() {
	condition := "main"
	if condition := "if"; 3 == (2 + 1) {
		fmt.Println("conditionPrint: ", condition)
	}
	fmt.Println("condition(main)", condition)

	data := []string{"one", "two", "three"}

	for _, v := range data {
		vcopy := v //
		go func() {
			fmt.Println("vcopy:", vcopy)
			fmt.Println("vcopy address:", &vcopy)
			fmt.Println("v: ", v)
			fmt.Println("v address ", &v)
		}()
	}

	time.Sleep(3 * time.Second)
	//goroutines vcopy print: one, two, three & diff address
	//goroutines v print: three, three, three & same address
}

// type field struct {
//   name string
// }
//
// func (p *field) print() {
//   fmt.Println(p.name)
// }
//
// func main() {
//   data := []field{{"one"},{"two"},{"three"}}
//
//   for _,v := range data {
//     v := v
//     go v.print()
//   }
//
//   time.Sleep(3 * time.Second)
//   //goroutines print: one, two, three
// }
