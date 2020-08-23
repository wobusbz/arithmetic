package tailStrings

import (
	"strings"
)

func TailStrings(s string) int {
	s = strings.Trim(s, " ")
	var count int
	for i := 0; i < len(s); i++ {
		if string(s[i]) != " " {
			count++
		} else if string(s[i]) == " " {
			count = 0
		}
	}
	return count
}
