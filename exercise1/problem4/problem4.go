package problem4

import (
	"strings"
)

func mapping(letters []string) map[string]string {

	charsMap := make(map[string]string)

	for _, char := range letters {
		charsMap[char] = strings.ToUpper(char)
	}

	return charsMap
}
