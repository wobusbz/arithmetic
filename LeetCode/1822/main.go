package main

import "fmt"

func main() {
	var nums = []int{9,72,34,29,-49,-22,-77,-17,-66,-75,-44,-30,-24}
	var result int = 1
	for i := 0; i < len(nums); i ++ {
		result *= nums[i]
	}
	fmt.Println(arraySign(nums))
}

func arraySign(nums []int) int {
	var result int = 1
	for i := 0; i < len(nums); i ++ {
		switch true {
		case nums[i] < 0:
			result *= -1
		case nums[i] > 0 :
			result *= 1
		default:
			return 0
		}
	}
	return result
}
