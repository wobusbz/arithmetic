package test15

import "fmt"

type DoubleLinkedList struct {
	head   *DoubleLinkedNode
	length int
}

func NewDoubleLinkedList() *DoubleLinkedList {
	head := NewDoubleLinkedNode(nil)
	return &DoubleLinkedList{head: head, length: 0}
}

func (dList *DoubleLinkedList) GetLength() int {
	return dList.length
}

func (dList *DoubleLinkedList) GetFirstNode() *DoubleLinkedNode {
	return dList.head.next
}

func (dList *DoubleLinkedList) InsertHead(node *DoubleLinkedNode) {
	pHead := dList.head
	if pHead.next == nil {
		node.next = nil
		pHead.next = node
		node.prev = pHead
		dList.length++
	} else {
		pHead.next.prev = node
		node.next = pHead.next
		pHead.next = node
		node.prev = pHead
		dList.length++
	}
}

func (dList *DoubleLinkedList) InsertBack(node *DoubleLinkedNode) {
	pHead := dList.head
	if pHead.next == nil {
		node.next = nil
		pHead.next = node
		node.prev = pHead
		dList.length++
	} else {
		for pHead.next != nil {
			pHead = pHead.next
		}
		pHead.next = node
		node.prev = pHead
		dList.length++
	}
}

func (dList *DoubleLinkedList) String() string {
	var listString1 string
	var listString2 string
	p := dList.head
	for p.next != nil {
		listString1 += fmt.Sprintf("%v --> ", p.next.value)
		p = p.next
	}
	listString1 += fmt.Sprintf("nil")
	listString1 += "\n"

	for p != dList.head {
		listString2 += fmt.Sprintf("%v --> ", p.value)
		p = p.prev
	}
	listString2 += fmt.Sprintf("nil")

	return listString1 + listString2
}
