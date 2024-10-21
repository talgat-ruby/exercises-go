package problem2

import (
	"strings"
)

func capitalize(names []string) []string {
	result := make([]string, len(names))

	for i, name := range names {
		result[i] = strings.Title(strings.ToLower(name))
	}
	return result
}
