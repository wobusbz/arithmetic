package test14

type SingleLinkedListNode struct {
	value interface{}
	pNext *SingleLinkedListNode
}

func NewSingleLinkedListNode(data interface{}) *SingleLinkedListNode {
	return &SingleLinkedListNode{value: data}
}

func (node *SingleLinkedListNode) Value() interface{} {
	return node.value
}

func (node *SingleLinkedListNode) PNext() *SingleLinkedListNode {
	return node.pNext
}
