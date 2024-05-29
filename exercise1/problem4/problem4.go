package problem4

import "strings"

func mapping(letters []string) map[string]string {
	lettersMap := map[string]string{}
	for i := range len(letters) {
		lettersMap[letters[i]] = strings.ToUpper(letters[i])
	}
	return lettersMap
}
