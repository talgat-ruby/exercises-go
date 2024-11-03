package problem5

import (
	"sort"
)

func products(products map[string]int, price int) []string {

	pricesDict := make(map[int][]string)
	prices := make([]int, 0, len(products))

	for k, v := range products {
		if v < price {
			delete(products, k)
			continue
		}
		if _, ok := pricesDict[v]; !ok {
			prices = append(prices, v)
		}
		pricesDict[v] = append(pricesDict[v], k)
	}
	sort.Slice(prices, func(i, j int) bool {
		return prices[i] > prices[j]
	})

	out := make([]string, 0, len(products))

	for _, price := range prices {
		keys := pricesDict[price]

		if len(keys) > 1 {
			sort.Strings(keys)
		}
		out = append(out, keys...)
	}

	return out
}
