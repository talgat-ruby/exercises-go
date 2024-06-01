package problem4

import "strings"

func mapping(symbols []string) map[string]string {
	result := make(map[string]string)
	for _, symbol := range symbols {
		result[strings.ToLower(symbol)] = strings.ToUpper(symbol)
	}

	return result
}
