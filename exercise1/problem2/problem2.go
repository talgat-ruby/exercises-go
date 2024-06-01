package problem2

import (
	"strings"
)

func capitalize(words []string) (capitalizedWords []string) {

	for _, word := range words {
		word = strings.Title(strings.ToLower(word))

		capitalizedWords = append(capitalizedWords, word)
	}

	return
}
