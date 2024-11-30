package problem2

import "strings"

func capitalize(names []string) []string {
	capitalN := make([]string, len(names))

	for i, name := range names {
		if len(name) > 0 {
			capitalN[i] = strings.ToUpper(name[:1]) + strings.ToLower(name[1:])
		} else {
			capitalN[i] = name
		}
	}

	return capitalN
}
