package problem5

func products(m map[string]int, min int) []string {
	out := make([]string, 0, len(m))
	intAr := make([]int, 0, len(m))
	for s, i := range m {
		if i > min {
			intAr, out = appendWithSorting(intAr, out, i, s)
		}
	}
	return out
}

func appendWithSorting(left []int, right []string, key int, val string) ([]int, []string) {
	for i, v := range left {
		if key > v || (key == v && compareStrings(val, right[i]) < 0) {
			left = append(left[:i], append([]int{key}, left[i:]...)...)
			right = append(right[:i], append([]string{val}, right[i:]...)...)
			return left, right
		}
	}
	return append(left, key), append(right, val)
}

func compareStrings(a, b string) int {
	var minLen int
	if len(a) < len(b) {
		minLen = len(a)
	} else {
		minLen = len(b)
	}
	for i := 0; i < minLen; i++ {
		if a[i] > b[i] {
			return 1
		} else if a[i] < b[i] {
			return -1
		}
	}
	if len(a) < len(b) {
		return -1
	} else if len(a) > len(b) {
		return 1
	} else {
		return 0
	}
}
