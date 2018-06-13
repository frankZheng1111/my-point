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
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.Val)
	if node.Left != nil {
		PrintBinaryTree(node.Left)
	}
	if node.Right != nil {
		PrintBinaryTree(node.Right)
	}
}

// 翻转二叉树
//
func InvertBinaryTree(node *TreeNode) *TreeNode {
	if node == nil {
		return node
	}
	if node.Left != nil || node.Right != nil {
		node.Left, node.Right = node.Right, node.Left
	}
	if node.Left != nil {
		InvertBinaryTree(node.Left)
	}
	if node.Right != nil {
		InvertBinaryTree(node.Right)
	}
	return node
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
	fmt.Println("原二叉树: ")
	PrintBinaryTree(node)
	fmt.Println("")
	fmt.Println("翻转后的二叉树: ")
	PrintBinaryTree(InvertBinaryTree(node))
}
