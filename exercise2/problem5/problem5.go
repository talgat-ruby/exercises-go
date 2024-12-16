package problem5

import (
	"sort"
)

func products(m map[string]int, price int) []string {
	s := make([]string, len(m), len(m))
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		if m[keys[i]] == m[keys[j]] {
			return keys[i] < keys[j]
		}
		return m[keys[i]] > m[keys[j]]
	})
	for _, k := range keys {
		if m[k] >= price {
			s = append(s, k)
		}
	}
	return s[3:]
}
