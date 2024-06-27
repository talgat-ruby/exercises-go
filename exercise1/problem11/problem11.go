package problem11

import "slices"

func removeDups[T comparable](someSlice []T) []T {
	var newSlice []T
	for _, item := range someSlice {
		if slices.Contains(newSlice, item) {
			continue
		}
		newSlice = append(newSlice, item)
	}
	return newSlice
}
