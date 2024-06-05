package problem2

import "strings"

func capitalize(names []string) []string {
	var result []string
	for _, name := range names {
		if len(name) > 0 {
			capitalizedName := strings.ToUpper(string(name[0])) + strings.ToLower(name[1:])
			result = append(result, capitalizedName)
		} else {
			result = append(result, name)
		}
	}
	return result
}
