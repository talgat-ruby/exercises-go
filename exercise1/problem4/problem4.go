package problem4

import "strings"

func mapping(inp []string) map[string]string {
	exp := map[string]string{}
	for i := 0; i < len(inp); i++ {
		exp[inp[i]] = strings.ToUpper(inp[i])
	}
	return exp
}
