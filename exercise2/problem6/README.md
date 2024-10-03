# Problem 6

Given two unique integer slices `a` and `b`, and an integer target value `v`, 
create a function to determine whether there is a pair of numbers that add up to the target value `v`, 
where one number comes from one array `a` and the other comes from the second array `b`. 
Return `true` if there is a pair that adds up to the target value and `false` otherwise.

```go
sumOfTwo([1, 2], [4, 5, 6], 5) // true

sumOfTwo([1, 2], [4, 5, 6], 8) // true

sumOfTwo([1, 2], [4, 5, 6], 3) // false

sumOfTwo([1, 2], [4, 5, 6], 9) // false
```
