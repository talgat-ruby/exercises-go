package problem2

import (
	"strings"
	"unicode"
)

func capitalize(names []string) []string {
	capitalizedNames := []string{}
	for i := range len(names) {
		name := names[i]
		if len(name) != 0 {
			name = strings.ToLower(name)
			chars := []rune(name)
			chars[0] = unicode.ToUpper(chars[0])
			name = strings.TrimSpace(string(chars))
		}
		capitalizedNames = append(capitalizedNames, name)
	}
	return capitalizedNames
}
