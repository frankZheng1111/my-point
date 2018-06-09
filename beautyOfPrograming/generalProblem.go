// 将帅问题
// 书 p13
package main

// 坐标,用于描述将帅九宫格坐标
// 1 2 3
// 4 5 6
// 7 8 9

import "fmt"

// a 为 将；b为帅 输出所有不对将的位置
func GeneralProblem() {
	for a := 1; a <= 9; a++ {
		for b := 1; b <= 9; b++ {
			if a%3 != b%3 {
				fmt.Printf("将: %v, 帅: %v \n", a, b)
			}
		}
	}
}

func GeneralProblemWithOneByte() {
	var count byte
	for count = 0; count < 81; count++ {
		// 此为上述方法的展开
		// 等式两边的count/9 + 1与 count%9 + 1 两边分别遍历了0+1 ~ 8+1 与 0+1 ~ 8+1
		if (count/9+1)%3 != (count%9+1)%3 {
			fmt.Printf("将: %v, 帅: %v \n", count/9+1, count%9+1)
		}
	}
}

func main() {
	// GeneralProblem()
	GeneralProblemWithOneByte() // 54 种结果
}
