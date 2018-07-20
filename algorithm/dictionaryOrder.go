// 字典序算法
// 给定一个正整数, 求出离该整数最近的大于自身的“换位数(将整数中的任意位数作任意次交换)”
// 输入12345 返回12354, 12354 => 12435, 12435 => 12453, 123520 => 125023
package main

import (
	"errors"
	"fmt"
)

// 返回大于原数的最小换位数
//
// 当一个整数从高位到低位是降序排列，则该给定整数为这些数字组合的最大值 => 不存在对应结果
// 当一个整数从高位到低位是升序排列，则该给定整数为这些数字组合的最小值
//
// 相关思路:
// 从低位往高位寻找近的一组升序(升序是指从左往右升序)排列，将较小的一位替换为其右侧大于其的数中最小的一位, 剩余的数字升序排列即可
func MinBiggerChangedNumber(num int) (int, error) {
	var result int
	var rightNums []int = make([]int, 0)
	var beforeNum, currentNum int // 当前位的前一位，当前位
	var hasUpSort bool
	var err error = errors.New("ResultNotExist")
	if num < 10 {
		return result, err
	}
	i := 10 // 从倒数第二位开始计算
	for ; num >= i; i *= 10 {
		beforeNum = (num / (i / 10)) % 10 // 当前位前一位
		currentNum = (num / i) % 10       // 当前位
		rightNums = append(rightNums, beforeNum)
		if currentNum < beforeNum {
			hasUpSort = true
			break
		}
	}
	if !hasUpSort {
		return result, err
	}
	var changedNum int
	for index, val := range rightNums {
		if val > currentNum {
			changedNum = val
			rightNums[index] = currentNum
			break
		}
	}
	result = (num/(i*10))*i*10 + changedNum*i
	for j := 0; j < len(rightNums); j++ {
		i /= 10
		result += rightNums[j] * i
	}
	return result, nil
}

func main() {
	a := 1234954321
	b, err := MinBiggerChangedNumber(a)
	fmt.Printf("原数为%d, 大于原数的最小换位数为 %d\n", a, b)
	if err != nil {
		fmt.Println(err)
	}
}
