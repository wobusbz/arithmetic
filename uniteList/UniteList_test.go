package uniteList

import "testing"

func TestUniteList(t *testing.T) {
	var nums1 = []int{1, 2, 2, 1}
	var nums2 = []int{2, 2}
	t.Log(UniteList(nums1, nums2))
	t.Log(UniteList1(nums1, nums2))
	t.Log(nums1)
}
