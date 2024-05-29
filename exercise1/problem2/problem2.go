package problem2

import "strings"

func capitalize(names []string) []string {
	for i, name := range names {
		if len(name) == 1 {
			names[i] = strings.ToTitle(name[:1])
		} else if len(name) > 1 {
			names[i] = strings.ToTitle(name[:1]) + strings.ToLower(name[1:])
		}
	}
	return names
}
