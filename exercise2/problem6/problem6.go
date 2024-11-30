package problem6

import "slices"

func sumOfTwo(a, b []int, target int) bool {
	if len(a) == 0 || len(b) == 0 {
		return false
	}
	slices.Sort(a)
	slices.Sort(b)

	if !(target >= a[0]+b[0] && target <= a[len(a)-1]+b[len(b)-1]) {
		return false
	}

	for i := 0; i < len(a); i++ {
		for j := len(b) - 1; j >= 0; j-- {
			if a[i] <= target && b[j] <= target {
				if a[i]+b[j] == target {
					return true
				}
			}
		}
	}
	return false
}
