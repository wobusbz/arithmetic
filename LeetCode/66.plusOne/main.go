/*
 * @Author: your name
 * @Date: 2021-04-23 11:36:06
 * @LastEditTime: 2021-04-23 13:04:19
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \arithmetic\LeetCode\66.plusOne\main.go
 */
package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{1, 9, 9}))
	fmt.Println(2 % 10)
}

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return nil
	}
	var newDigits = make([]int, 0, len(digits)+1)
	for i := len(digits) - 1; i > -1; i-- {
		digits[i]++
		if digits[i]%10 == 0 {
			digits[i] = 0
		}
		if digits[i]%10 > 0 {
			break
		}
		if digits[0]%10 == 0 {
			newDigits = append(newDigits, 1)
		}
	}
	return append(newDigits, digits...)
}
