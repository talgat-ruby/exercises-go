# Problem 12

Create a function that takes a map and returns the keys and values as separate slices.
Return the keys sorted alphabetically, and their corresponding values in the same order.

```go
keysAndValues(map[string]int{"a": 1, "b": 2, "c": 3}) // []string{"a", "b", "c"}, []int{1, 2, 3}

keysAndValues(map[string]string{"a": "Apple", "b": "Microsoft", "c": "Google"}) // []string{"a", "b", "c"}, []string{"Apple", "Microsoft", "Google"}

keysAndValues(map[int]bool{1: true, 2: false, 3: false}, ) // []int{1, 2, 3}, []bool{true, false, false},
```
