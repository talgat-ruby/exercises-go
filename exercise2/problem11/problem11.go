package problem11

func removeDups[T int | bool | string](sl []T) []T {
	var result []T
	var l bool = false
	for i, a := range sl {
		l = false
		for _, b := range result {
			if a == b {
				l = true
				break
			}
		}
		if l == false {
			result = append(result, sl[i])
		}
	}
	return result
}
