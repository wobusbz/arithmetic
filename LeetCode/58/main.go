package main

import (
	"bytes"
	"fmt"
)

func main() {
	var s = "lrloseumgh"

	fmt.Println(reverseLeftWords1(s, 6))
	fmt.Println(reverseLeftWords2(s, 6))
}

func reverseLeftWords1(s string, n int) string {
	if len(s) < n {
		return ""
	}
	return s[n:] + s[:n]
}

// 看别人思路
func reverseLeftWords2(s string, n int) string {
	if len(s) < n {
		return ""
	}
	var result = bytes.NewBuffer(make([]byte, 0, len(s)))
	for i := n; i < n + len(s); i++ {
		result.WriteByte(s[i%len(s)])
	}
	return result.String()
}
