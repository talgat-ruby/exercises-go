# Problem 11

Create a function that takes a slice of items, removes all duplicate items and returns a new slice in the 
same sequential order as the old slice (minus duplicates).

```go
removeDups([]int{1, 0, 1, 0}) // [1, 0]

removeDups([]bool{true, false, false, true}) // ["The", "big", "cat"]

removeDups([]string{"John", "Taylor", "John"}) // ["John", "Taylor"]
```
