package main

// 基础的冒泡排序可优化的点
// 2018程序员小灰: 什么是冒泡排序

import (
	"fmt"
)

func BubbleSort(nums []int) []int {
	var isSorted bool          // 优化2: 若再一次大循环中已经没有元素交换, 则数列已经有序，则可以跳出循环
	var lastExchangedIndex int // 最后一个交换的索引, 该索引后的元素皆有序
	len := len(nums)
	sortBorder := len - 1
	for i := 0; i < len-1; i++ {
		isSorted = true
		// 优化1: 内循环可以随外循环的次数减少(已经作为最大数排最后的值不用再比), 即从末尾开始计，有序区为i
		// for j := 0; j < len-1-i; j++ {
		// 优化1-1: 记录最后一次交换的边界,该边界右边的集合有序
		// (这里的边界记录交换后元素靠前的那一个位置索引: 因为下一轮循环的条件时j < sortborder, 故下一次循环最后j=sortborder-1, 是sortborder-1对比sortborder)
		for j := 0; j < sortBorder; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				isSorted = false
				lastExchangedIndex = j
			}
		}
		sortBorder = lastExchangedIndex
		if isSorted {
			break
		}
	}
	return nums
}

func main() {
	nums := []int{3, 9, 4, 2, 1, 5, 6, 7, 8}
	fmt.Println(BubbleSort(nums))
}
