package problem2

import (
	"strings"
)

func capitalize(array []string) []string {
	var newArray []string
	for _, word := range array {
		loweredWord := strings.ToLower(word)
		if loweredWord != "" {
			firstLetterCapitalized := strings.ToUpper(loweredWord[0:1]) + loweredWord[1:]
			newArray = append(newArray, firstLetterCapitalized)
		} else {
			newArray = append(newArray, loweredWord)
		}
	}
	return newArray
}
