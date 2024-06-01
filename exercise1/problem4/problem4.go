package problem4

import "strings"

func mapping(letters []string) map[string]string {
	newMap := make(map[string]string, len(letters))

	for _, letter := range letters {
		newMap[letter] = strings.ToUpper(letter)
	}

	return newMap
}
