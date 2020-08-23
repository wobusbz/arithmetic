package queue

import "testing"

func TestQueue(t *testing.T) {
	var queue = NewQueue()

	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)
	queue.EnQueue(5)
	queue.EnQueue(6)
	t.Log(queue.DeQueue())
	t.Log(queue.DeQueue())
	t.Log(queue.DeQueue())
}
