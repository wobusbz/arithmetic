package test10

import "testing"

//队列
func TestQueue_Init(t *testing.T) {
	q := new(Queue)
	q.Init(5)
	for i := 2; i < 5; i++ {
		q.Push(i)
	}
	t.Log(q.Pop())
	t.Log(q.Pop())
	t.Log(q.Pop())
}
