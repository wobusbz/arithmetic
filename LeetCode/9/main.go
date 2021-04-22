/*
 * @Author: your name
 * @Date: 2021-04-22 09:33:06
 * @LastEditTime: 2021-04-22 09:59:19
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \LeetCode\9\main.go
 */
package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(1221))
}

func isPalindrome(x int) bool {
	if (x < 0 || x%10 == 0) && x > 0 {
		return false
	}
	var (
		num  int
		temp = x
	)
	for x > 0 {
		num = num*10 + x%10
		x /= 10
	}
	return num == temp
}
