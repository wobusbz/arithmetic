package main

import "fmt"

func main() {
	var nums = []int{1, 1, 1, 1}
	fmt.Println(runningSum(nums))
}

func runningSum(nums []int) []int {
	var result = make([]int, 0, len(nums))
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		result =append(result, sum)
	}
	return result
}