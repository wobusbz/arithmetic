package test11

type Node struct {
	data  interface{}
	pNext *Node
}

type CodeLink struct {
	Head *Node
	Tail *Node
}

type CodeStack interface {
	IsEmpty() bool
	Push(data interface{})
	Pop() interface{}
	Length() int
}

func (n *Node) IsEmpty() bool {
	if n.pNext == nil {
		return true
	} else {
		return false
	}
}

func (n *Node) Push(data interface{}) {
	newNode := &Node{data, nil}
	newNode.pNext = n.pNext
	n.pNext = newNode
}

func (n *Node) Pop() interface{} {
	if n.IsEmpty() {
		return nil
	}
	value := n.pNext.data
	n.pNext = n.pNext.pNext
	return value
}

func (n *Node) Length() int {
	var index int
	var pNext = n
	for pNext != nil {
		pNext = pNext.pNext
		index++
	}
	return index
}
