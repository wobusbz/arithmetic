package test9

import "fmt"

type Statck struct {
	stack []int
}

func NewStatck() *Statck {
	return &Statck{}
}

func (s *Statck) push(n int) {
	s.stack = append(s.stack, n)
}

func (s *Statck) pop() int {
	if s.len() == 0 {
		return -1
	}
	value := s.stack[s.len()-1:][0]
	s.stack = s.stack[:s.len()-1]
	return value
}

// 返回栈得长度
func (s *Statck) len() int {
	return len(s.stack)
}

func (s *Statck) print() {
	fmt.Println(s.stack)
}
