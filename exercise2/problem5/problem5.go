package problem5

import "sort"

func products(prices map[string]int, minimum int) []string {
	var result []string

	for product, price := range prices {
		if price >= minimum {
			result = append(result, product)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return prices[result[i]] < prices[result[j]] ||
			(prices[result[i]] == prices[result[j]] && result[i] < result[j])
	})

	return result
}
