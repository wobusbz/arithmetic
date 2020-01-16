package _7

import "testing"

// oldArr = []int{1, 1, 2, 3, 5, 7}
// newArr = []int{2, 3, 5, 7, 1, 1}
// 不重复的长度为4
func TestRemoveArr(t *testing.T) {
	var intArr = []int{1, 1, 2, 3, 5, 7}
	if index := RemoveArr(intArr, 1); index == 4 {
		t.Logf("返回数组的新长度:%d\n", index)
		t.Log(index)
	}
}
