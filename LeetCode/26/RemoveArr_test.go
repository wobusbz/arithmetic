package _6

import (
	"testing"
)

var arrRemoveTest = []struct {
	fname string
	f     func([]int) int
	out   int
}{
	{"RemoveArr", RemoveArr, 2},
}

func TestRemoveArr(t *testing.T) {
	nums := []int{1, 1, 2, 2}
	for _, v := range arrRemoveTest {
		if v.f(nums) != v.out {
			t.Errorf("RemoveArr:%d actual:%d", v.f(nums), v.out)
		}
	}
}
