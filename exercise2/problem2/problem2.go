package problem2

import (
	"strings"
)

func capitalize(names []string) []string {

	result := make([]string, len(names))

	for i, name := range names {
		if len(name) > 0 {
			result[i] = strings.ToUpper(name[:1]) + strings.ToLower(name[1:])
		} else {
			result[i] = name
		}
	}
	return result
}
