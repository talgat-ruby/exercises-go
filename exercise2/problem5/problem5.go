package problem5

import "sort"

func products(list map[string]int, price int) []string {
	tempMap := make(map[int][]string)
	output := make([]string, 0)
	costs := make([]int, 0)

	for key, value := range list {
		if value >= price {
			tempMap[value] = append(tempMap[value], key)
		}
	}
	for key, value := range tempMap {
		sort.Strings(value)
		costs = append(costs, key)
	}

	sort.Ints(costs)

	for i := len(costs) - 1; i >= 0; i-- {
		output = append(output, tempMap[costs[i]]...)
	}

	return output
}
