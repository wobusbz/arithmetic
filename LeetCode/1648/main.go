package main

import (
	"fmt"
)

func main() {
	fmt.Println(countConsistentStrings("cad", []string{"cc","acd","b","ba","bac","bad","ac","d"}))
}

func countConsistentStrings(allowed string, words []string) int {
	var count int
	var allowedMap = make(map[uint8]bool, len(allowed))
	for k := range allowed {
		allowedMap[allowed[k]] = true
	}
	fmt.Println(allowedMap['b'])
	for i := 0; i < len(words); i++ {
		var flag bool
		for j := 0; j < len(words[i]); j++ {
			if !allowedMap[words[i][j]] {
				break
			}
			if allowedMap[words[i][j]] && len(words[i])-1 == j {
				fmt.Println(words[i])
				flag = true
			}
		}
		if flag {
			count++
		}
	}
	return count
}
