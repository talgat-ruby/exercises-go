package problem2

import (
	"strings"
)

func capitalize(names []string) []string {
	for i, name := range names {
		if len(name) > 0 {
			names[i] = strings.ToUpper(string(name[0])) + strings.ToLower(name[1:])
		}
	}
	return names
}
