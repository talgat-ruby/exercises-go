package problem4

import (
	"strings"
)

func mapping(inp []string) map[string]string {
	m := make(map[string]string)

	for _, letter := range inp {
		m[letter] = strings.ToUpper(letter)
	}
	return m
}
