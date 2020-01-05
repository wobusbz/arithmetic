package test1

import "testing"

func TestBinarySearch(t *testing.T) {
	intArr := []int{10, 2, 5, 7, 6, 100, 99}
	BinarySearch(intArr, 50)
}
