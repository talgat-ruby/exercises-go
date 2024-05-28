package problem4

import "strings"

func mapping(inp []string) map[string]string {
	res := make(map[string]string)
	for i := 0; i < len(inp); i++ {
		res[inp[i]] = strings.ToUpper(inp[i])
	}
	return res
}
