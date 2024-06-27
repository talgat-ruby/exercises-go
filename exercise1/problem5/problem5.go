package problem5

type kv struct {
	Key   string
	Value int
}

func products(catalog map[string]int, price int) []string {
	var kvs []kv
	var products []string

	for k, v1 := range catalog {
		if v1 < price {
			delete(catalog, k)
		} else {
			kvs = append(kvs, kv{k, v1})
		}
	}

	n := len(kvs)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if kvs[j].Value < kvs[j+1].Value {
				kvs[j], kvs[j+1] = kvs[j+1], kvs[j]
			}
			if kvs[j].Value == kvs[j+1].Value && kvs[j].Key > kvs[j+1].Key {
				kvs[j], kvs[j+1] = kvs[j+1], kvs[j]
			}
		}
	}

	for _, v1 := range kvs {
		products = append(products, v1.Key)
	}

	return products
}
