package problem2

import "strings"

func capitalize(names []string) []string {
	result := make([]string, len(names))
	for i, name := range names {
		if len(name) >= 2 {
			result[i] = strings.ToUpper(string(name[0])) + strings.ToLower(string(name[1:]))
		} else {
			result[i] = strings.ToUpper(name)
		}
	}
	return result
}
