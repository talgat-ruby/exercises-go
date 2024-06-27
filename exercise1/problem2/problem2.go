package problem2

import (
	"strings"
)

func capitalize(names []string) []string {
	var result []string
	for _, name := range names {
		result = append(result, strings.Title(strings.ToLower(name)))
	}
	return result
}
