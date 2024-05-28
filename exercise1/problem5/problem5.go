package problem5

import "sort"

type kv struct {
	Key   string
	Value int
}

func products(catalog map[string]int, minPrice int) []string {
	var kvs []kv
	for key, value := range catalog {
		if value >= minPrice {
			kvs = append(kvs, kv{key, value})

		}
	}

	sort.Slice(kvs, func(i, j int) bool {
		if kvs[i].Value == kvs[j].Value {
			return kvs[i].Key < kvs[j].Key
		}
		return kvs[i].Value > kvs[j].Value
	})
	res := make([]string, len(kvs))
	for i := 0; i < len(res); i++ {
		res[i] = kvs[i].Key
	}
	return res
}
