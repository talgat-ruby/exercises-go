package problem2

import "strings"

func capitalize(s []string) []string {
	capitalized := make([]string, len(s))
	for i, word := range s {
		if len(word) > 0 {
			capitalized[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		} else {
			capitalized[i] = word
		}
	}
	return capitalized

}
