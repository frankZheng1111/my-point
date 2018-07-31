package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 对折链表
//1 2 3 4 5 => 1 5 2 4 3
//1 2 3 4 => 4 1 3 2
func FoldList(head *ListNode) *ListNode {
	// 找到中间节点 偶数节点向下取整
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	midNode := head
	lQuick := head
	var odd bool
	for lQuick.Next != nil && lQuick.Next.Next != nil {
		lQuick = lQuick.Next.Next
		midNode = midNode.Next
		if lQuick.Next == nil {
			odd = true
		}
	}
	h1 := head
	h2 := midNode.Next
	midNode.Next = nil
	if odd {
		return MergeTwoList(h1, reverseListIteratively(h2))
	} else {
		return MergeTwoList(reverseListIteratively(h2), h1)
	}
}

func MergeTwoList(l1, l2 *ListNode) *ListNode {
	head := l1
	for l2 != nil {
		l1Next := l1.Next
		l1.Next = l2
		l2Next := l2.Next
		l2.Next = l1Next
		l2 = l2Next
		if l1.Next.Next == nil {
			l1.Next.Next = l2
			break
		}
		l1 = l1.Next.Next
	}
	return head
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
		fmt.Printf("%d ", node.Val)
		node = node.Next
	}
	fmt.Println("")
}

func QuickSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var tmpFunc func(head, sn, en *ListNode)
	tmpFunc = func(head, sn, en *ListNode) {
		if sn == nil || sn == en || sn.Next == nil || sn.Next == en {
			return
		}
		pn := sn
		pVal := pn.Val
		maxPtr := sn.Next
		for maxPtr != en {
			if maxPtr.Val <= pVal {
				pn = pn.Next
				pn.Val, maxPtr.Val = maxPtr.Val, pn.Val
			}
			maxPtr = maxPtr.Next
		}
		pn.Val, sn.Val = sn.Val, pn.Val
		tmpFunc(head, sn, pn)
		tmpFunc(head, pn.Next, en)
		return
	}
	tmpFunc(head, head, nil)
	return head
}

func main() {
	node0 := ListNode{5, nil}
	node1 := ListNode{4, &node0}
	node2 := ListNode{3, &node1}
	node3 := ListNode{2, &node2}
	head := ListNode{1, &node3}
	fmt.Println("翻转前(原链表):")
	printList(&head)
	fmt.Println("翻转后(迭代翻转):")
	reverseHead := *reverseListIteratively(&head)
	printList(&reverseHead)
	fmt.Println("翻转后(递归翻转):")
	// reReverseHead := reverseListRecursively(&reverseHead)
	// printList(reReverseHead)
	fmt.Println("链表快排:")
	printList(QuickSortList(&reverseHead))
	fmt.Println("对折链表:")
	printList(FoldList(&head))
}
