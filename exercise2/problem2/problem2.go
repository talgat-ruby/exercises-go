package problem2

import (
	"strings"
)

func capitalize(names []string) []string {
	Upper := make([]string, len(names))
	for i, name := range names {
		if len(name) > 0 {
			Upper[i] = strings.ToUpper(string(name[0])) + strings.ToLower(name[1:])
		} else {
			Upper[i] = name
		}
	}
	return Upper
}
