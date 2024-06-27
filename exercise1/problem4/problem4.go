package problem4

import "strings"

func mapping(values []string) map[string]string {
	m := map[string]string{}
	for _, s := range values {
		m[s] = strings.ToUpper(s)
	}

	return m
}
