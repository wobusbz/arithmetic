/*
 * @Author: your name
 * @Date: 2021-04-22 13:27:24
 * @LastEditTime: 2021-04-22 13:44:22
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \arithmetic\LeetCode\38\main.go
 */
package main

import "fmt"

func main() {
	fmt.Println(countAndSay(13))
}

func countAndSay(n int) string {
	if n < 10 {
		return fmt.Sprintf("%d%d", 1, n)
	}
	for n > 0 {
		fmt.Println(n % 10)
		n /= 10
	}
	return ""
}

func Format(num, n int) {

}
