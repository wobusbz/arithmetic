package test9

import "testing"

func TestNewStatck(t *testing.T) {
	s := NewStatck()
	s.push(1)
	s.push(2)
	s.push(3)
	s.print()
	t.Log(s.pop())
	s.print()
}
