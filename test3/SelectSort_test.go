package test3

import "testing"

func TestSelectSort(t *testing.T) {
	intArr := []int{10, 2, 5, 7, 6, 100, 99, 33}
	SelectSort(intArr)
	t.Log(intArr)
}
