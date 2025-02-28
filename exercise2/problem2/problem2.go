package problem2

import "strings"

func capitalize(names []string) []string {

	capitalizedNames := []string{}
	for _, name := range names{
		capitalizedNames = append(capitalizedNames, strings.Title(strings.ToLower(name)))
	}

	return capitalizedNames
}
