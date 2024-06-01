package problem2

import "unicode"

func capitalize(names []string) []string {
	for i, name := range names {
		names[i] = capitalizeName(name)
	}
	return names
}

func capitalizeName(name string) string {
	runes := []rune(name)
	for i := 0; i < len(runes); i++ {
		if i == 0 {
			runes[i] = unicode.ToUpper(runes[i])
			continue
		}
		runes[i] = unicode.ToLower(runes[i])
	}
	return string(runes)
}
