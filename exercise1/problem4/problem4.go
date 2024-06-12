package problem4

import "strings"

func mapping(letters []string) map[string]string {
	res := make(map[string]string)
	for _, letter := range letters {
		res[letter] = strings.ToUpper(letter)
	}
	return res
}
