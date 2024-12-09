package problem5

import "sort"

func products(goods map[string]int, num int) []string {
	filtered := []string{}

	for product, price := range goods {
		if price >= num {
			filtered = append(filtered, product)
		}
	}
	sort.Slice(filtered, func(i, j int) bool {
		if goods[filtered[i]] != goods[filtered[j]] {
			return goods[filtered[i]] > goods[filtered[j]]
		}
		return filtered[i] < filtered[j]
	})
	return filtered
}
