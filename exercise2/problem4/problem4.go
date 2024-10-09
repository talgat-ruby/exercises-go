package problem4

import "strings"

func mapping(letters []string) map[string]string {
	result := make(map[string]string)
	for _, letter := range letters {
		result[letter] = strings.ToUpper(letter)
	}
	return result
}
