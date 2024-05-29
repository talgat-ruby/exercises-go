package problem8

func simplify(list []string) map[string]int {

	indMap := make(map[string]int, len(list))
	load(indMap, list)

	return indMap
}

func load(m map[string]int, students []string) {
	for i, name := range students {
		m[name] = i
	}
}
