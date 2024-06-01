package problem4

import "strings"

func mapping(letters []string) map[string]string {
	m := make(map[string]string)

	for _, letter := range letters {
		m[letter] = strings.ToUpper(letter)
	}

	return m
}
