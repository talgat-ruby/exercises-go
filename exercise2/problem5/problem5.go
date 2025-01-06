package problem5

import "sort"

func products(prices map[string]int, minPrice int) []string {

	var result []string

	for product, price := range prices {
		if price >= minPrice {
			result = append(result, product)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		if prices[result[i]] == prices[result[j]] {
			return result[i] < result[j]
		}
		return prices[result[i]] < prices[result[j]]
	})

	return result
}
