package problem2

import (
	"strings"
)

func capitalize(names []string) []string {
	for i, name := range names {
		names[i] = strings.Title(strings.ToLower(name))
	}
	return names
}
