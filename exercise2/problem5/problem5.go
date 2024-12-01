package problem5

import "sort"

func products(items map[string]int, minPrice int) []string {
	var result []string

	for name, price := range items {
		if price >= minPrice {
			result = append(result, name)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		if items[result[i]] != items[result[j]] {
			return items[result[i]] > items[result[j]]
		}
		return result[i] < result[j]
	})

	return result
}
