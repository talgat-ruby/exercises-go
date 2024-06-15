package problem5

import (
	"slices"
	"sort"
)

func products(m map[string]int, min int) []string {
	for k, v := range m {
		if v < min {
			delete(m, k)
		}
	}

	if len(m) < 1 {
		return make([]string, 0)
	}

	values := make([]int, 0)

	for _, v := range m {
		if !slices.Contains(values, v) {
			values = append(values, v)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	res := make([]string, 0)

	for _, outerV := range values {
		similar := make([]string, 0)
		for k, innerV := range m {
			if outerV == innerV {
				similar = append(similar, k)
			}
		}
		sort.Strings(similar)
		for _, v := range similar {
			res = append(res, v)
		}
	}
	return res
}
