package problem5

import "sort"

type KVPair struct {
	key string
	val int
}

func products(items map[string]int, price int) []string {
	// smth
	output := []KVPair{}
	for k, v := range items {
		if v >= price {
			output = append(output, KVPair{k, v})
		}
	}

	sort.Slice(output, func(i int, j int) bool {
		if output[i].val == output[j].val {
			return output[i].key < output[j].key
		}
		return output[i].val > output[j].val
	})

	newOutput := []string{}
	for _, item := range output {
		newOutput = append(newOutput, item.key)
	}
	return newOutput
}
