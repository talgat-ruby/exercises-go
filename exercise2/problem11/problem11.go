package problem11

func removeDups[T comparable](slc []T) []T {
	new_slc := make([]T, 0, len(slc))
	proverka_slc := make(map[any]struct{})
	for _, v := range slc {
		if _, ok := proverka_slc[v]; !ok {
			proverka_slc[v] = struct{}{}
			new_slc = append(new_slc, v)
		}
	}
	return new_slc
}
