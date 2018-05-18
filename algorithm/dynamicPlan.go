package main

// 动态规划
//
// - 最优子结构
// - 获取边界情况
// - 状态转移方程

import "fmt"

// 有一座高度是10级台阶的楼梯，从下往上走，每跨一步只能向上1级或者2级台阶。要求用程序来求出一共有多少种走法。
//
func getClimbingWay(totalStep int) (resultWay int) {
	// 当只有或只剩一格时走法
	if totalStep == 1 {
		return 1
	}
	// 当只有或只剩2格时走法
	if totalStep == 2 {
		return 2
	}

	// 假设到倒数第二格的走法有x种(最后一步走两步, ps: 连着走两步的情况记入后一统计), 假设到倒数第1格的走法有y种(最后一步走一步) 总方法有x+y
	// W(n) = W(n - 1) + W(n - 2)
	// W(1) = 1
	// W(2) = 2

	last1StepWay := 2 // W(2)
	last2StepWay := 1 // W(1)
	for i := 3; i <= totalStep; i++ {
		resultWay = last1StepWay + last2StepWay
		last2StepWay = last1StepWay
		last1StepWay = resultWay
	}
	return
}

func getMaxGold(workerNum int) (maxGold int) {
	type Gold struct {
		reserve    int
		needWorker int
	}
	var GOLDS = []Gold{
		{500, 5},
		{200, 3},
		{350, 3},
		{300, 4},
		{400, 5},
	}
	// F(1,10) = GOLDS[1].reserve ( GOLDS[1].needWorker <= 10 ) || 0
	// F(2,10) = MAX( F(1,10), F(1,10-GOLDS[2].needWorker) + GOLDS[2].reserve )

	Max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	var maxGoldCal func(golds []Gold, workerNum int) int
	maxGoldCal = func(golds []Gold, workerNum int) int {
		length := len(golds)
		if length == 0 {
			return 0
		} else if length == 1 {
			maxGold := golds[0].reserve
			if workerNum < GOLDS[0].needWorker {
				maxGold = 0
			}
			return maxGold
		} else {
			if workerNum >= golds[length-1].needWorker {
				return Max(maxGoldCal(golds[:length-1], workerNum), maxGoldCal(golds[:length-1], workerNum-golds[length-1].needWorker)+golds[length-1].reserve)
			} else {
				return maxGoldCal(golds[:length-1], workerNum)
			}
		}
	}
	maxGold = maxGoldCal(GOLDS, workerNum)

	return
}

func main() {
	fmt.Println("总10格楼梯，最多一次跨2格，几种走法 ", getClimbingWay(10))
	fmt.Println("国王金矿问题, 10名工人最多挖多少金矿: ", getMaxGold(10))
}
