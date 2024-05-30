package problem5

import "sort"

func products(products map[string]int, min int) []string {
	var filtered []string

	for name, price := range products {
		if price >= min {
			filtered = append(filtered, name)
		}
	}

	// проверка на пустую мапу
	if len(filtered) == 0 {
		return []string{}
	}

	sort.Slice(filtered, func(i, j int) bool {
		priceI, priceJ := products[filtered[i]], products[filtered[j]]
		if priceI == priceJ {
			return filtered[i] < filtered[j]
		}
		return priceI > priceJ
	})

	return filtered
}
