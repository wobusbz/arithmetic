/*
 * @Author: your name
 * @Date: 2021-04-22 10:01:51
 * @LastEditTime: 2021-04-22 11:04:24
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \LeetCode\14\main.go
 */
package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flower", "flower", "flower"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	for j := 0; j < len(strs[0]); j++ {
		for i := 1; i < len(strs); i++ {
			if j == len(strs[i]) || strs[i][j] != strs[0][j] {
				return strs[0][:j]
			}
		}
	}
	return strs[0]
}
