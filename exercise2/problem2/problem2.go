package problem2

import "strings"

func capitalize(names []string) []string {
	for i, v := range names {
		if v == "" {
			continue
		}
		firstLetter := strings.ToUpper(string(v[0]))
		lowerS := strings.ToLower(v[1:])
		names[i] = firstLetter + lowerS
	}
	return names
}
