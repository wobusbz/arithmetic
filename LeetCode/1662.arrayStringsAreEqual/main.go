package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(arrayStringsAreEqual([]string{"abc", "d", "defg"}, []string{"abcddefg"}))
}

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	return strings.EqualFold(strings.Join(word1, ""), strings.Join(word2, ""))
}
