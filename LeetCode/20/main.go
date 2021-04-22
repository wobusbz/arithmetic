/*
 * @Author: your name
 * @Date: 2021-04-22 11:09:58
 * @LastEditTime: 2021-04-22 13:07:54
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \LeetCode\20\main.go
 */
package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid("()"))
}

// "(){}}{"
func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	var stack = New(len(s))
	var smap = map[uint8]uint8{
		')': '(',
		']': '[',
		'}': '{',
	}
	for i := 0; i < len(s); i++ {
		if _, ok := smap[s[i]]; ok {
			if stack.Pop() != smap[s[i]] {
				return false
			}
		} else {
			stack.Push(s[i])
		}
	}
	fmt.Println(stack.StackList, stack.ll)
	return stack.IsEmpty()
}

type StackQueue struct {
	StackList []uint8
	ll        int
}

func New(n int) *StackQueue {
	return &StackQueue{StackList: make([]uint8, 0, n), ll: -1}
}

func (s *StackQueue) Push(v uint8) {
	s.StackList = append(s.StackList, v)
	s.ll++
}

func (s *StackQueue) Pop() uint8 {
	if s.ll == -1 {
		return 0
	}
	var data = s.StackList[s.ll]
	if s.ll >= 0 {
		s.StackList = s.StackList[:s.ll]
	}
	s.ll--
	return data
}

func (s *StackQueue) IsEmpty() bool {
	return s.ll == -1
}
