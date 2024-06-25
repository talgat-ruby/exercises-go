package problem4

import (
	"strings"
)

func mapping(someStrings []string) map[string]string {
	result := make(map[string]string)
	for _, char := range someStrings {
		result[char] = strings.ToUpper(char)
	}
	return result
}
