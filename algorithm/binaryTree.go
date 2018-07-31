package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BinaryTreePreOrder2(node *TreeNode) []*TreeNode {
	var nodes []*TreeNode = []*TreeNode{}
	var stack []*TreeNode = []*TreeNode{}
	for len(stack) > 0 || node != nil {
		if node == nil {
			return nodes
		}
		nodes = append(nodes, node)
		if node.Left != nil {
			stack = append(stack, node)
			node = node.Left
		} else if node.Right != nil {
			node = node.Right
			continue
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return nodes
}

// 先序遍历二叉树
//
func BinaryTreePreOrder(node *TreeNode) []*TreeNode {
	var nodes []*TreeNode = []*TreeNode{}
	var recursivefunc func(node *TreeNode)
	recursivefunc = func(node *TreeNode) {
		if node == nil {
			return
		}
		nodes = append(nodes, node)
		fmt.Printf("%d ", node.Val)
		if node.Left != nil {
			recursivefunc(node.Left)
		}
		if node.Right != nil {
			recursivefunc(node.Right)
		}
	}
	recursivefunc(node)
	return nodes
}

// 中序遍历二叉树
//
func BinaryTreeLDR(node *TreeNode) []*TreeNode {
	var nodes []*TreeNode = []*TreeNode{}
	var recursivefunc func(node *TreeNode)
	recursivefunc = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			recursivefunc(node.Left)
		}
		nodes = append(nodes, node)
		fmt.Printf("%d ", node.Val)
		if node.Right != nil {
			recursivefunc(node.Right)
		}
	}
	recursivefunc(node)
	return nodes
}

func RebuildBinaryTree(preOrderNodes, LDRNodes []*TreeNode) *TreeNode {
	if len(preOrderNodes) == 0 {
		return nil
	}
	node := preOrderNodes[0]
	node.Left = nil
	node.Right = nil
	var nodeIndexInLDR int
	for index, value := range LDRNodes {
		if value == node {
			nodeIndexInLDR = index
		}
	}
	node.Left = RebuildBinaryTree(preOrderNodes[1:1+nodeIndexInLDR], LDRNodes[:nodeIndexInLDR])
	node.Right = RebuildBinaryTree(preOrderNodes[1+nodeIndexInLDR:], LDRNodes[nodeIndexInLDR+1:])
	return node
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
	fmt.Println("原二叉树(先序遍历): ")
	preOrderNodes := BinaryTreePreOrder2(node)
	fmt.Println("")
	fmt.Println("原二叉树(中序遍历): ")
	LDRNodes := BinaryTreeLDR(node)
	fmt.Println("")
	fmt.Println("原二叉树和通过先序遍历与中序遍历重构出的二叉树先序遍历")
	rebuildNode := RebuildBinaryTree(preOrderNodes, LDRNodes)
	BinaryTreePreOrder(rebuildNode)
	fmt.Println("")
	fmt.Println("翻转后的二叉树: ")
	BinaryTreePreOrder(InvertBinaryTree(node))
}
