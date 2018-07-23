// https://www.cnblogs.com/chengxiao/p/6129630.html
package main

import "fmt"

func HeapSort(values []int) {
	// 构建初始堆
	//
	buildHeap(values)
	// 将i-1节点放在最后，对[: i - 1]重新维护堆
	for i := len(values); i > 1; i-- {
		values[0], values[i-1] = values[i-1], values[0]
		HeapPercolateUp(values[:i-1], 0)
	}
}

// 仅将一个最大值放在堆顶
func buildHeap(values []int) {
	for i := len(values) - 1; i >= 0; i-- { //////一定得从后往前调整，
		HeapPercolateUp(values, i)
	}
}

// 将最大值放在index的位置
func HeapPercolateUp(values []int, index int) {
	iMax := index
	length := len(values)
	iLeft := 2*index + 1
	iRight := 2*index + 2
	if iLeft < length && values[iLeft] > values[index] {
		iMax = iLeft
	}
	if iRight < length && values[iRight] > values[iMax] {
		iMax = iRight
	}
	if values[index] < values[iMax] {
		values[iMax], values[index] = values[index], values[iMax]
	}
}

func main() {
	data := []int{1, 2, 3, 4, 5, 8, 7, 6, 6}
	HeapSort(data)
	fmt.Println(data)
}
