package problem4

import "strings"

func mapping(letters []string) map[string]string {
	customMap := make(map[string]string)
	for _, letter := range letters {
		customMap[letter] = strings.ToUpper(string(letter))
	}
	return customMap
}
