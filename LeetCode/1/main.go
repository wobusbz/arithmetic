/*
 * @Author: your name
 * @Date: 2021-04-22 13:10:57
 * @LastEditTime: 2021-04-22 13:21:48
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \arithmetic\LeetCode\1\main.go
 */
package main

import "fmt"

func main() {
	var nums = []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
}

func twoSum(nums []int, target int) []int {
	var resultMap = make(map[int]int, len(nums))
	for k := range nums {
		if i, ok := resultMap[target-nums[k]]; ok {
			return []int{i, k}
		}
		resultMap[nums[k]] = k
	}
	return nil
}
