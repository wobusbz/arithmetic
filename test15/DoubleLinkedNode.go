package test15

type DoubleLinkedNode struct {
	value interface{}
	prev  *DoubleLinkedNode
	next  *DoubleLinkedNode
}

func NewDoubleLinkedNode(value interface{}) *DoubleLinkedNode {
	return &DoubleLinkedNode{value: value}
}

func (node *DoubleLinkedNode) Value() interface{} {
	return node.value
}

func (node *DoubleLinkedNode) Prev() *DoubleLinkedNode {
	return node.prev
}

func (node *DoubleLinkedNode) Next() *DoubleLinkedNode {
	return node.next
}
