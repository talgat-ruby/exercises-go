package problem2

import "unicode"

func capitalize(names []string) []string {
	for i, name := range names {
		if len(name) > 0 {
			runes := []rune(name)
			for j := range runes {
				if j == 0 {
					runes[j] = unicode.ToUpper(runes[j])
				}
			}
			names[i] = string(runes)
		}
	}
	return names
}
