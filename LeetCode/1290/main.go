package main

import (
	"fmt"
)

func main() {
	var node1 = &ListNode{Val: 1}
	var node2 = &ListNode{Val: 0}
	var node3 = &ListNode{Val: 1}
	var node4 = &ListNode{Val: 1}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	fmt.Println(getDecimalValue(node1))
	// 1,0,0,1,0,0,1,1,1,0,0,0,0,0,0
	fmt.Println(1<<6 + 1<<7 + 1<<8 + 1<<11 + 1<<14)
}

//  * Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	var result int
	for ; head != nil; head = head.Next {
		result = result<<1 + head.Val
	}
	return result
}

