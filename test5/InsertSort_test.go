package test5

import "testing"

func TestInsertSort(t *testing.T) {
	intArr := []int{10, 2, 5, 7, 6, 100, 99}
	InsertSort(intArr)
	intArr1 := []int{1, 2, 3, 4, 5, 6}
	InsertSort1(intArr1)
}
