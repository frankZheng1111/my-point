package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for i := 0; i != 10; i = i + 1 {
		fmt.Fprintf(os.Stdout, "result is %d\r", i)
		time.Sleep(time.Second * 1)
	}
	fmt.Println("over")
}
