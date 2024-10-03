package problem4

import "strings"

func mapping(letters []string) map[string]string {
	output := make(map[string]string, len(letters))
	for _, letter := range letters {
		output[letter] = strings.ToUpper(letter)
	}
	return output
}
