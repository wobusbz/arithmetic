package maxPrefix

import (
	"strings"
)

func MaxPrefix(prefixs []string) string {
	if len(prefixs) < 1 {
		return ""
	}

	prefix := prefixs[0]

	for k, _ := range prefixs {
		for strings.Index(prefixs[k], prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}
