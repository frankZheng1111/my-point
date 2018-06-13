package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 先序遍历二叉树
//
func PrintBinaryTree(node *TreeNode) {
	fmt.Printf("%d ", node.Val)
	if node.Left != nil {
		PrintBinaryTree(node.Left)
	}
	if node.Right != nil {
		PrintBinaryTree(node.Right)
	}
}

func main() {
	node := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	PrintBinaryTree(node)
}
