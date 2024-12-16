package problem4

import "strings"

func mapping(s []string) map[string]string {
	m := map[string]string{}
	for _, v := range s {
		m[strings.ToLower(v)] = strings.ToUpper(v)
	}
	return m
}
