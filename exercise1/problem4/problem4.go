package problem4

import "github.com/talgat-ruby/exercises-go/pkg/util"

func mapping(keys []string) map[string]string {
	m := make(map[string]string, len(keys))
	for _, v := range keys {
		m[v] = util.ToUpper(v)
	}

	return m
}
