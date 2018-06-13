package main

import "fmt"

type HeapNode struct {
	Val   int
	Left  *HeapNode
	Right *HeapNode
}

// 先序遍历二叉树
//
func PrintHeap(node *HeapNode) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.Val)
	if node.Left != nil {
		PrintHeap(node.Left)
	}
	if node.Right != nil {
		PrintHeap(node.Right)
	}
}

func BuildCompleteBinaryTree(rawData []int) *HeapNode {
	tmpNodes := make([]*HeapNode, len(rawData))
	for index, data := range rawData {
		tmpNodes[index] = &HeapNode{
			Val: data,
		}
		if index != 0 {
			fatherIndex := index - index/2 - 1
			if tmpNodes[fatherIndex].Left == nil {
				tmpNodes[fatherIndex].Left = tmpNodes[index]
			} else if tmpNodes[fatherIndex].Right == nil {
				tmpNodes[fatherIndex].Right = tmpNodes[index]
			}
		}
	}
	return tmpNodes[0]
}

func main() {
	rawData := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	node := BuildCompleteBinaryTree(rawData)
	PrintHeap(node)
}
