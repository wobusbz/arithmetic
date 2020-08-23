package queue

type IQueue interface {
	Size() int
	Front() interface{}
	End() interface{}
	IsEmpty() bool
	EnQueue(data interface{})
	DeQueue() interface{}
	Clear()
}

type Queue struct {
	dataStore []interface{}
	theSize   int
}

func NewQueue() *Queue {
	var queue = new(Queue)
	queue.dataStore = make([]interface{}, 0)
	queue.theSize = 0
	return queue
}

func (q *Queue) Size() int {
	return q.theSize
}

func (q *Queue) Front() interface{} {
	if q.Size() == 0 {
		return nil
	}
	return q.dataStore[0]
}

func (q *Queue) End() interface{} {
	if q.Size() == 0 {
		return nil
	}
	return q.dataStore[q.Size()-1]
}

func (q *Queue) IsEmpty() bool {
	return q.theSize == 0
}

func (q *Queue) EnQueue(data interface{}) {
	q.dataStore = append(q.dataStore, data)
	q.theSize++
}

func (q *Queue) DeQueue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	data := q.dataStore[0]
	if q.Size() > 1 {
		q.dataStore = q.dataStore[1:q.Size()]
	} else {
		q.dataStore = make([]interface{}, 0)
	}
	q.theSize--
	return data
}

func (q *Queue) Clear() {
	q.dataStore = make([]interface{}, 0)
	q.theSize = 0
}
