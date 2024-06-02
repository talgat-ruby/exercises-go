package problem4

import "unicode"

func mapping(letters []string) map[string]string {
	result := make(map[string]string)
	for _, letter := range letters {
		if len(letter) > 0 {
			result[letter] = string(unicode.ToUpper(rune(letter[0])))
		}
	}
	return result
}
