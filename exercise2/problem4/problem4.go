package problem4

import "strings"

func mapping(inp []string) map[string]string {
	exp := make(map[string]string)
	for _, value := range inp {
		exp[value] = strings.ToUpper(value)
	}
	return exp
}
