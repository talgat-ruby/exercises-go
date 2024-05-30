package problem2

import "strings"

func capitalize(names []string) []string {
	for i := range names {
		names[i] = strings.Title(strings.ToLower(names[i]))
	}
	return names
}
