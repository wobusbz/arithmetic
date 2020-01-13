package test10

// 队列
type Queue struct {
	queue []int
	rear  int // 尾部
	front int // 首部
	size  int
}

const (
	ERROREMPTY = 0
)

func (this *Queue) Init(size int) {
	this.queue = make([]int, size)
	this.rear = 0
	this.front = 0
	this.size = size
}

func (this *Queue) Push(element int) {
	if !this.is_filled() {
		this.rear = (this.rear + 1) % this.size
		this.queue[this.rear] = element
	}
}

func (this *Queue) Pop() int {
	if !this.is_Empty() {
		this.front = (this.front + 1) % this.size
		return this.queue[this.front]
	} else {
		return ERROREMPTY
	}

}

func (this *Queue) is_Empty() bool {
	return this.rear == this.front
}

func (this *Queue) is_filled() bool {
	return (this.rear+1)%this.size == this.front
}
