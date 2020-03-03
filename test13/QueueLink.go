package test13

type Node struct {
	data  interface{}
	pNext *Node
}

type QueueStack struct {
	Font *Node
	Tail *Node
}

type QueueLink interface {
	Length() int
	Push(data interface{})
	Pop() interface{}
}

func (q *QueueStack) Length() int {
	var index int
	var font = q.Font
	for font != nil {
		font = font.pNext
		index++
	}
	return index
}

func (q *QueueStack) Push(data interface{}) {
	newNode := &Node{data, nil}
	if q.Font == nil {
		q.Font = newNode
		q.Tail = newNode
	} else {
		q.Tail.pNext = newNode
		q.Tail = q.Tail.pNext
	}
}

func (q *QueueStack) Pop() interface{} {
	if q.Font == nil {
		return nil
	}
	newNode := q.Font
	if q.Font == q.Tail {
		q.Font = nil
		q.Tail = nil
	} else {
		q.Font = q.Font.pNext
	}
	return newNode.data
}
