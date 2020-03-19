package test14

import "testing"

func TestNewSingleLinkList(t *testing.T) {
	list := NewSingleLinkList()
	node1 := NewSingleLinkedListNode(1)
	node2 := NewSingleLinkedListNode(2)
	node3 := NewSingleLinkedListNode(3)
	node4 := NewSingleLinkedListNode(5)
	list.InsertNodeBack(node1)
	list.InsertNodeBack(node2)
	list.InsertNodeBack(node3)
	list.InsertNodeBack(node4)
	node5 := NewSingleLinkedListNode(4)
	list.InsertNodeValueBack(3, node5)
	t.Log(list.String())
	t.Log(list.GetNodeAtIndex(2))
	t.Log(list.String())
	list.DeleteNode(node5)
	list.DeleteAtiIndex(1)
	t.Log(list.String())
}
