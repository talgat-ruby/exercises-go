package problem5

import (
	"sort"
)

func products(productMap map[string]int, minPrice int) []string {
	var filteredProducts []string

	for product, price := range productMap {
		if price >= minPrice {
			filteredProducts = append(filteredProducts, product)
		}
	}

	sort.Slice(filteredProducts, func(i, j int) bool {
		if productMap[filteredProducts[i]] == productMap[filteredProducts[j]] {
			return filteredProducts[i] < filteredProducts[j]
		}
		return productMap[filteredProducts[i]] > productMap[filteredProducts[j]]
	})

	return filteredProducts
}
