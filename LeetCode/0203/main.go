package main

import "fmt"

func main() {

	var node1 = &ListNode{val: 4}
	var node2 = &ListNode{val: 5}
	var node3 = &ListNode{val: 1}
	var node4 = &ListNode{val: 9}
	node1.next = node2
	node2.next = node3
	node3.next = node4
	node2.val = node2.next.val
	node2.next = node2.next.next
	fmt.Println(node2)
	for ; node1 != nil; node1 = node1.next {
		fmt.Println(node1.val)
	}
}

type ListNode struct {
	val  int
	next *ListNode
}

func deleteNode(node *ListNode) {
	if node != nil {
		node.val = node.next.val
		node.next = node.next.next
	}
}