package problem2

import "strings"

func capitalize(names []string) []string {
	for i, s := range names {
		if len(s) != 0 {
			names[i] = strings.ToUpper(s[0:1]) + strings.ToLower(s[1:])
		}
	}

	return names
}
