package problem2

import "strings"

func capitalize(names []string) []string {
	for i, name := range names {
		if len(name) < 1 {
			continue
		}
		names[i] = strings.ToUpper(name[:1]) + strings.ToLower(name[1:])
	}
	return names
}
