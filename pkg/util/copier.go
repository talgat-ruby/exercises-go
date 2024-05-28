package util

func copySlice[V any](slice []V) []V {
	cp := make([]V, len(slice))
	for i := 0; i < len(slice); i++ {
		cp[i] = slice[i]
	}

	return cp
}
