package problem9

type resultType func(...int) []int

func factory(multiple int) resultType {
	return func(list ...int) []int {
		result := make([]int, len(list))
		for i, value := range list {
			result[i] = value * multiple
		}
		return result
	}
}
