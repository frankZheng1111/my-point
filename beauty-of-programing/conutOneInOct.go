package main

import "fmt"

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
	lowerNum := 0   // 低位
	currNum := 0    // 当前位
	higherNum := 0  //高位
	var baseNum int //当前位数 1为个位，10为十位
	for baseNum = 1; baseNum <= num; baseNum = baseNum * 10 {
		lowerNum = num % baseNum
		higherNum = num / (baseNum * 10)
		currNum = (num - higherNum*baseNum*10 - lowerNum) / baseNum
		switch currNum {
		case 0:
			count += higherNum * baseNum
		case 1:
			count += higherNum*baseNum + lowerNum + 1
		default:
			count += (higherNum + 1) * baseNum
		}
	}
	return
}

func main() {
	n := 12345
	fmt.Printf("1到%v中1出现的次数(遍历): %v\n", n, Count1LessIntegerRaw(n))
	fmt.Printf("1到%v中1出现的次数(规则): %v\n", n, Count1LessIntegerByRule(n))
}
