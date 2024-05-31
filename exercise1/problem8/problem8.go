package problem8

func simplify(keys []string) map[string]int {
	m := make(map[string]int)
	load(&m, &keys)
	return m
}

func load(m *map[string]int, keys *[]string) {
	for i, key := range *keys {
		(*m)[key] = i
	}
}
