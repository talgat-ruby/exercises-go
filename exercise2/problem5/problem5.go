package problem5

import "sort"

func products(catalog map[string]int, minPrice int) []string {
	var filtered []string
	for product, price := range catalog {
		if price >= minPrice {
			filtered = append(filtered, product)
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		if catalog[filtered[i]] == catalog[filtered[j]] {
			return filtered[i] < filtered[j]
		}
		return catalog[filtered[i]] > catalog[filtered[j]]
	})

	return filtered
}
