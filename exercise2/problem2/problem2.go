package problem2

import (
	"strings"
)

func capitalize(names []string) []string {
	var capitalized []string
	for _, name := range names {
		if len(name) > 0 {

			capitalizedName := strings.ToUpper(string(name[0])) + strings.ToLower(name[1:])
			capitalized = append(capitalized, capitalizedName)
		} else {
			capitalized = append(capitalized, name)
		}
	}
	return capitalized
}
