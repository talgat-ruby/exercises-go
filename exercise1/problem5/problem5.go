package problem5

import (
	"sort"
	"strings"
)

func products(products map[string]int, threshold int) []string {
	keys := make([]string, 0, len(products))
	for key := range products {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return products[keys[i]] > products[keys[j]] ||
			products[keys[i]] == products[keys[j]] && strings.Compare(keys[i], keys[j]) < 0
	})

	res := make([]string, 0)
	for _, k := range keys {
		if products[k] > threshold {
			res = append(res, k)
		}
	}

	return res
}
