package problem11

func removeDups[T comparable](input []T) []T {
	var res []T
    for _, i := range input {
		count := 0
		for _, j := range res {
			if j == i {
				count++
			}
		}
		if count == 0 {
			res = append(res, i)
		}
	}
	return res
}

