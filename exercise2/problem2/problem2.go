package problem2

import "strings"

func capitalize(arr []string) []string {

	var capitalizedNames []string

	for _, name := range arr {
		if len(name) > 0 {

			capitalized := strings.ToUpper(string(name[0])) + strings.ToLower(name[1:])
			capitalizedNames = append(capitalizedNames, capitalized)
		} else {
			// Если имя пустое, добавляем его как есть
			capitalizedNames = append(capitalizedNames, "")
		}
	}

	return capitalizedNames
}
