package main

import (
	"fmt"
	"runtime"
)

func main() {
	done := false

	runtime.GOMAXPROCS(3)
	fmt.Println(runtime.GOMAXPROCS(-1))

	go func() {
		done = true
	}()

	for !done {
		fmt.Println("notdone!")
		runtime.Gosched()
	}
	fmt.Println("done!")
}
