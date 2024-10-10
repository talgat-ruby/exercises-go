package problem4

import "strings"

func mapping(lowerCase []string) map[string]string {
	result := make(map[string]string)

	for _, value := range lowerCase {
		result[value] = strings.ToUpper(value)
	}
	return result
}
