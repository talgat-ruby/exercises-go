package problem5

import "sort"

func products(catalog map[string]int, minimalPrice int) []string {
	var result []string

	for product, price := range catalog {
		if price >= minimalPrice {
			result = append(result, product)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		if catalog[result[i]] == catalog[result[j]] {
			return result[i] < result[j]
		}
		return catalog[result[i]] > catalog[result[j]]
	})

	return result
}
