package problem11

type Comparable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~string | ~float32 | ~float64
}

func keysAndValues[K Comparable, V any](m map[K]V) ([]K, []V) {
	keys := make([]K, 0, len(m))
	values := make([]V, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	bubleSort(keys)
	for _, v := range keys {
		values = append(values, m[v])
	}
	return keys, values
}

func bubleSort[T Comparable](slice []T) {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}
