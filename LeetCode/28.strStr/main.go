/*
 * @Author: your name
 * @Date: 2021-04-25 10:44:22
 * @LastEditTime: 2021-04-25 16:29:54
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \arithmetic\LeetCode\28.strStr\main.go
 */
package main

import (
	"fmt"
)

func main() {
	// "mississippi"
	// "issip"
	fmt.Println(strStr("hello", "ll"))
}

func strStr(haystack string, needle string) int {
	if haystack == "" && needle == "" || needle == "" || haystack == needle {
		return 0
	}
	if len(haystack) < len(needle) || len(haystack) > 5*10<<4 && len(needle) > 5*10<<4 {
		return -1
	}
	var j int
	for i := 0; i < len(haystack); i++ {
		for j = 0; j < len(needle); j++ {
			if i < len(haystack) && haystack[i] == needle[j] {
				fmt.Println("----------->", i, j, string(needle[j]))
				i++
			} else {
				break
			}
		}
		if j == len(needle) {
			return i - j
		}
		i -= j
	}
	return -1
}
