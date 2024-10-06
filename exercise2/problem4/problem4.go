package problem4

import (
	"strings"
)

func mapping(arr []string) map[string]string {
	result := map[string]string{}
	for _, value := range arr {
		result[value] = strings.ToUpper(value)
	}
	return result
}
