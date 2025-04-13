package problem2

import (
	"strings"
)

func capitalize(names []string) []string {
	capitalizedNames := make([]string, len(names))
	for i, name := range names {

		capitalizedNames[i] = strings.Title(strings.ToLower(name))
	}
	return capitalizedNames
}
