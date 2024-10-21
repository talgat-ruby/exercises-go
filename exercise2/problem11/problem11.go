package problem11

import "slices"

func removeDups[T comparable](nums []T) []T {
	result := make([]T, 0)
	for _, v := range nums {
		if !slices.Contains(result, v) {
			result = append(result, v)
		}
	}
	return result
}
