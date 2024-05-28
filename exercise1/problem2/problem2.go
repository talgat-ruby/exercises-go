package problem2

import "strings"

func capitalize(names []string) []string {
	for i := 0; i < len(names); i++ {
		names[i] = strings.ToLower(names[i])
		names[i] = strings.Title(names[i])
	}
	return names
}
