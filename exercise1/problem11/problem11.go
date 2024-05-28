package problem11

func removeDups[T any](s []T) []T {
	m := make(map[any]struct{})
	var arr []T
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; !ok {
			m[s[i]] = struct{}{}
			arr = append(arr, s[i])
		}
	}
	return arr

}
