package problem4

import "strings"

func mapping(letters []string) map[string]string {
	letterMapping := make(map[string]string)

	for _, letter := range letters {
		letterMapping[letter] = strings.ToUpper(letter)
	}

	return letterMapping
}
