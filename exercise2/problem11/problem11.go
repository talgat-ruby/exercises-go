package problem11

func removeDups[T comparable](slice []T) []T {
	res := []T{}

	check := map[T]bool{}
	for _, n := range slice {
		if !check[n] {
			res = append(res, n)
			check[n] = true
		}

	}
	return res
}

// func removeDups[T comparable](slice []T) []T {
// 	res := []T{}
// 	var prevn T

// 	for i, n := range slice {
// 		if i == 0 || !reflect.DeepEqual(n, prevn) { // Проверяем на первом элементе и на дубликатах
// 			res = append(res, n)
// 			prevn = n
// 		}
// 	}

// 	return res
// }
