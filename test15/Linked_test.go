package test15

import "testing"

func TestNewDoubleLinkedList(t *testing.T) {
	dList := NewDoubleLinkedList()

	node1 := NewDoubleLinkedNode(1)
	node2 := NewDoubleLinkedNode(2)
	node3 := NewDoubleLinkedNode(3)
	node4 := NewDoubleLinkedNode(4)
	dList.InsertHead(node1)
	dList.InsertHead(node2)
	dList.InsertHead(node3)
	dList.InsertHead(node4)
	node5 := NewDoubleLinkedNode(5)
	node6 := NewDoubleLinkedNode(6)
	node7 := NewDoubleLinkedNode(7)
	node8 := NewDoubleLinkedNode(8)
	dList.InsertBack(node5)
	dList.InsertBack(node6)
	dList.InsertBack(node7)
	dList.InsertBack(node8)
	t.Log(dList.String())
}
