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

func main() {
	node := &HeapNode{
		Val: 1,
		Left: &HeapNode{
			Val: 2,
			Left: &HeapNode{
				Val: 3,
			},
			Right: &HeapNode{
				Val: 4,
			},
		},
		Right: &HeapNode{
			Val: 5,
			Left: &HeapNode{
				Val: 6,
			},
			Right: &HeapNode{
				Val: 7,
			},
		},
	}
	fmt.Println("原堆: ")
	PrintHeap(node)
	fmt.Println("")
}
