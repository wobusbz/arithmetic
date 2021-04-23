/*
 * @Author: your name
 * @Date: 2021-04-23 11:14:30
 * @LastEditTime: 2021-04-23 11:30:14
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \arithmetic\LeetCode\58.lengthOfLastWord\main.go
 */
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLastWord("a "))
	fmt.Println(' ')
	fmt.Println('\n')
}

func lengthOfLastWord(s string) int {
	var count int
	s = strings.TrimRight(s, " ")
	for i := len(s) - 1; i > -1; i-- {
		if s[i] == 32 {
			break
		}
		count++
	}
	return count
}
