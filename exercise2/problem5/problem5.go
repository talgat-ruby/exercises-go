package problem5

import "sort"

func products(goods map[string]int, num int) []string {
	res := []string{}
	prices := []int{}
	newPrices := []int{}
	// var check bool
	for _, v := range goods {
		prices = append(prices, v)
	}
	sort.Ints(prices)
	for i := len(prices) - 1; i >= 0; i-- {
		newPrices = append(newPrices, prices[i])
	}
	// prevname := ""
	for _, v1 := range newPrices {
		if v1 > num {
			for key, value := range goods {
				if value == v1 {
					res = append(res, key)
					break

				}
			}
		}

	}
	return res
}
