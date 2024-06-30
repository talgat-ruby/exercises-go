package problem2

import (
	"strings"
	"unicode"
)

func capitalize(names []string) []string {
	var capitalizedNames []string
	for _, name := range names {
		if len(name) > 0 {
			first := string(unicode.ToUpper(rune(name[0])))
			next := strings.ToLower(name[1:])
			capitalizedNames = append(capitalizedNames, first+next)
		} else {
			capitalizedNames = append(capitalizedNames, "")
		}
	}
	return capitalizedNames
}
