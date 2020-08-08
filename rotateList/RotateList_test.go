package rotateList

import "testing"

func TestRotateList(t *testing.T) {
	intArr := []int32{1, 2, 3, 4, 5, 6, 7}
	t.Log(RotateList(intArr, 3))
}
