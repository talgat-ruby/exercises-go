package problem2

import (
	"strings"
	"unicode"
)

func capitalize(names []string) []string {
	capitalizedNames := make([]string, len(names))
	for i, name := range names {
		if len(name) > 0 {
			firstLetter := unicode.ToUpper(rune(name[0]))
			restOfName := strings.ToLower(name[1:])
			capitalizedNames[i] = string(firstLetter) + restOfName
		} else {
			capitalizedNames[i] = name
		}
	}
	return capitalizedNames
}
