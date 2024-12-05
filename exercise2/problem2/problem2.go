package problem2

import (
	"strings"
	"unicode"
)

func capitalize(names []string) []string {
	for i, name := range names {
		names[i] = strings.ToLower(name)
		r := []rune(names[i])
		if len(r) == 0 {
			continue
		}
		r[0] = unicode.ToUpper(r[0])
		names[i] = string(r)
	}
	return names
}
