package uniteList

import (
	"sort"
)

// 输入: nums1 = [1,2,2,1], nums2 = [2,2]

// 输出: [2,2]

// BF 破解

func UniteList(arrs0, arrs1 []int) []int {
	var intMap = map[int]int{}
	for _, v := range arrs0 {
		intMap[v] += 1
	}
	result := make([]int, 0, len(intMap))
	for _, v := range arrs1 {
		if val, ok := intMap[v]; ok {
			result = append(result, val)
		}
	}
	return result
}

func UniteList1(arrs0, arrs1 []int) []int {
	var i, j, k = 0, 0, 0
	sort.Ints(arrs0)
	sort.Ints(arrs1)
	for i < len(arrs0) && j < len(arrs1) {
		if arrs0[i] > arrs1[j] {
			j++
		} else if arrs0[i] < arrs1[j] {
			i++
		} else {
			arrs0[k] = arrs0[i]
			i++
			j++
			k++
		}
	}
	return arrs0[:k]
}
