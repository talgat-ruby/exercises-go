package problem4

import "strings"

func mapping(letters []string) map[string]string {
	result := make(map[string]string)
	for i, letter := range letters {
		letters[i] = strings.ToLower(letter)
		result[letters[i]] = strings.ToUpper(letter)
	}
	return result
}
