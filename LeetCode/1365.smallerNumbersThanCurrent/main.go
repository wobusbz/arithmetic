package main

import (
	"fmt"
)

func main() {
	fmt.Println(smallerNumbersThanCurrent([]int{8, 1, 2, 2, 3}))
}

func smallerNumbersThanCurrent(nums []int) []int {
	var s = make([]int, len(nums))
	for k := range nums {
		var n int
		for i := 0; i < len(nums); i++ {
			if nums[k] > nums[i] {
				n++
			}
		}
		s[k] = n
	}
	return s
}
