package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxSubArray([]int{-1, 1, -3, 4, -1, 2, 1, -5, -4}))
	fmt.Println(-2 + 1)
}

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var max = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
