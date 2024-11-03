package problem2

import (
	"unicode"
)

func capitalize(names []string) []string {
	for i, v := range names {
		// names[i] = strings.Title(strings.ToLower(v))

		capitalized := []rune(v)
		for pos, char := range capitalized {
			if pos == 0 {
				capitalized[pos] = unicode.ToUpper(char)
				continue
			}
			capitalized[pos] = unicode.ToLower(char)
		}
		names[i] = string(capitalized)
	}

	return names
}
