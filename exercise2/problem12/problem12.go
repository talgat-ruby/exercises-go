package problem11

func keysAndValues[T string | int, S any](m map[T]S) ([]T, []S) {
	left := make([]T, 0, len(m))
	right := make([]S, 0, len(m))
	for k, v := range m {
		left, right = appendWithSorting(left, right, k, v)
	}
	return left, right
}

func appendWithSorting[T string | int, S any](left []T, right []S, key T, val S) ([]T, []S) {
	for i, v := range left {
		if key <= v {
			left = append(left[:i], append([]T{key}, left[i:]...)...)
			right = append(right[:i], append([]S{val}, right[i:]...)...)
			return left, right
		}
	}
	return append(left, key), append(right, val)
}
