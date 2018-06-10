package main

import "fmt"
import "math"

func Count1InInteger(num int) (count int) {
	for num > 0 {
		if num%10 == 1 {
			count++
		}
		num /= 10
	}
	return
}

func Count1LessIntegerRaw(num int) (count int) {
	for i := 1; i <= num; i++ {
		count += Count1InInteger(i)
	}
	return
}

func Count1LessIntegerByRule(num int) (count int) {
	x := 0
	for num > 0 {
		count += num % 10 * (1 + x*int(math.Pow(10, float64(x-1))))
		num /= 10
		x++
	}
	return
}

func main() {
	n := 110
	fmt.Printf("1到%v中1出现的次数(遍历): %v\n", n, Count1LessIntegerRaw(n))
	fmt.Printf("1到%v中1出现的次数(规则): %v\n", n, Count1LessIntegerByRule(n))
}
