package problem11

import (
	"reflect"
	"sort"
)

func keysAndValues[T ~int | ~string | ~bool, P ~int | ~string | ~bool](itemsMap map[T]P) ([]T, []P) {
	keys := []T{}
	vals := []P{}
	for k, _ := range itemsMap {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		a := reflect.ValueOf(keys[i])
		b := reflect.ValueOf(keys[j])
		if a.Kind() == reflect.String {
			return a.String() < b.String()
		} else {
			return a.Int() < b.Int()
		}
	})

	for _, key := range keys {
		vals = append(vals, itemsMap[key])
	}
	return keys, vals
}
