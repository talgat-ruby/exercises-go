package problem5

import (
	"fmt"
	"sort"
)

func products(catalog map[string]int, minPrice int) []string {
	var filtered [][]string

	for name, price := range catalog {
		if price >= minPrice {
			filtered = append(filtered, []string{name, fmt.Sprintf("%d", price)})
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		if filtered[i][1] == filtered[j][1] {
			return filtered[i][0] < filtered[j][0]
		}
		return filtered[i][1] > filtered[j][1]
	})

	var result []string
	for _, product := range filtered {
		result = append(result, product[0])
	}
	return result
}
