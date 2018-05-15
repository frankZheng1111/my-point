package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseListIteratively(head *ListNode) *ListNode {
	var newListNode *ListNode
	var node *ListNode
	newListNode = nil
	node = head
	for {
		if node == nil {
			break
		}
		next := node.Next
		node.Next = newListNode
		newListNode = node
		node = next
	}
	return newListNode
}

func reverseListRecursively(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	newHead := reverseListRecursively(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func printList(head *ListNode) {
	node := head
	for {
		if node == nil {
			break
		}
		fmt.Println(node.Val)
		node = node.Next
	}
}

func main() {
	node1 := ListNode{1, nil}
	node2 := ListNode{2, &node1}
	node3 := ListNode{3, &node2}
	head := ListNode{4, &node3}
	fmt.Println("翻转前(原链表):")
	printList(&head)
	fmt.Println("翻转后(迭代翻转):")
	reverseHead := *reverseListIteratively(&head)
	printList(&reverseHead)
	fmt.Println("翻转后(递归翻转):")
	reReverseHead := reverseListRecursively(&reverseHead)
	printList(reReverseHead)
}
