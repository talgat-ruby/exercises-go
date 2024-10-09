package problem8

func simplify(list []string) map[string]int {
	result := make(map[string]int, len(list))
	for i, name := range list {
		result[name] = i
	}
	return result
}

func load(m map[string]int, students []string) {
	for i, name := range students {
		m[name] = i
	}
}
