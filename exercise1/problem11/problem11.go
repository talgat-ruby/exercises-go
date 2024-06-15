package problem11

import "slices"

func removeDups[T comparable](s []T) []T {
	distinct := make([]T, 0)
	for _, v := range s {
		if !slices.Contains(distinct, v) {
			distinct = append(distinct, v)
		}
	}
	return distinct
}
