package problem4

import "strings"

func mapping(inp []string) map[string]string {
	intM := make(map[string]string)
	for _, v := range inp {
		intM[v] = strings.ToUpper(v)
	}
	return intM
}
