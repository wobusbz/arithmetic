/*
 * @Author: your name
 * @Date: 2021-04-25 16:35:55
 * @LastEditTime: 2021-04-25 18:15:49
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \arithmetic\LeetCode\88.merge\main.go
 */
package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums1 = []int{1, 2, 3, 0, 0, 0, 0}
	var nums2 = []int{2, 5, 6, 7}
	merge(nums1, 3, nums2, 4)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) != m+n {
		return
	}
	copy(nums1[m:], nums2[:n])
	sort.Ints(nums1)
}
