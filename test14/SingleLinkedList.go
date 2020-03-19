package test14

import (
	"fmt"
)

type SingleLinked interface {
	GetFirstNode() *SingleLinkedListNode
	InsertNodeFront(node *SingleLinkedListNode) // 头部插入
	InsertNodeBack(node *SingleLinkedListNode)  // 尾部插入

	InsertNodeValueFront(dest interface{}, node *SingleLinkedListNode) bool
	InsertNodeValueBack(dest interface{}, node *SingleLinkedListNode) bool

	GetNodeAtIndex(index int) *SingleLinkedListNode // 根据索引抓取指定位置的节点
	DeleteNode(dest *SingleLinkedListNode) bool     // 删除节点
	DeleteAtiIndex(index int) bool                  // 删除指定位置的节点
	String() string                                 // 返回链表字符串
}

type SingleLinkedList struct {
	head   *SingleLinkedListNode
	length int
}

func NewSingleLinkList() *SingleLinkedList {
	head := NewSingleLinkedListNode(nil)
	return &SingleLinkedList{head: head, length: 0}
}

func (list *SingleLinkedList) GetFirstNode() *SingleLinkedListNode {
	return list.head.pNext
}

func (list *SingleLinkedList) InsertNodeFront(node *SingleLinkedListNode) {
	if list.head == nil {
		list.head.pNext = node
		node.pNext = nil
		list.length++
	} else {
		head := list.head
		node.pNext = head.pNext
		head.pNext = node
	}
}

func (list *SingleLinkedList) InsertNodeBack(node *SingleLinkedListNode) {
	if list.head == nil {
		list.head.pNext = node
		node.pNext = nil
		list.length++
	} else {
		head := list.head
		for head.pNext != nil {
			head = head.pNext
		}
		head.pNext = node
		list.length++
	}
}

// InsertNodeValueBack
func (list *SingleLinkedList) InsertNodeValueFront(dest interface{}, node *SingleLinkedListNode) bool {
	pHead := list.head
	isFind := false
	for pHead.pNext != nil {
		if pHead.pNext.value == dest {
			isFind = true
			break
		}
		pHead = pHead.pNext
	}

	if isFind {
		node.pNext = pHead.pNext
		pHead.pNext = node
		list.length++
		return true
	} else {
		return false
	}
	return false
}

func (list *SingleLinkedList) InsertNodeValueBack(dest interface{}, node *SingleLinkedListNode) bool {
	pHead := list.head
	isFind := false
	for pHead.pNext != nil {
		if pHead.value == dest {
			isFind = true
			break
		}
		pHead = pHead.pNext
	}

	if isFind {
		node.pNext = pHead.pNext
		pHead.pNext = node
		list.length++
		return true
	} else {
		return false
	}
}

func (list *SingleLinkedList) String() string {
	var listString string
	p := list.head
	for p.pNext != nil {
		listString += fmt.Sprintf("%v --> ", p.pNext.value)
		p = p.pNext
	}
	listString += fmt.Sprintf("nil")
	return listString
}

func (list *SingleLinkedList) GetNodeAtIndex(index int) *SingleLinkedListNode {
	if index > list.length-1 || index < 0 {
		return nil
	} else {
		pHead := list.head
		for index > -1 {
			pHead = pHead.pNext
			index--
		}
		return pHead
	}
}

func (list *SingleLinkedList) DeleteNode(dest *SingleLinkedListNode) bool {
	if dest == nil {
		return false
	}

	pHead := list.head
	for pHead.pNext != nil && pHead.pNext != dest {
		pHead = pHead.pNext
	}

	if pHead.pNext == dest {
		pHead.pNext = pHead.pNext.pNext
		list.length--
		return true
	} else {
		return false
	}
}

func (list *SingleLinkedList) DeleteAtiIndex(index int) bool {
	if index > list.length-1 || index < 0 {
		return false
	} else {
		pHead := list.head
		for index > 0 {
			pHead = pHead.pNext
			index--
		}

		pHead.pNext = pHead.pNext.pNext
		list.length--
		return true
	}
}
