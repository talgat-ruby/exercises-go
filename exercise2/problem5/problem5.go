package problem5

import "sort"

func products(m map[string]int, price int) []string {
	result := []string{}
	for key, val := range m {
		if val > price {
			result = append(result, key)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		if m[result[i]] == m[result[j]] {
			return result[i] < result[j]
		}
		return m[result[i]] > m[result[j]]
	})

	return result
}
