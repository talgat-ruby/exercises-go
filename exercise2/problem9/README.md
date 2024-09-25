# Problem 9

Create a function that takes an `int` as its parameter and returns another function.
The returned function must take a as many as needed `int` as its parameter, and return a `slice` of the `int`
multiplied by the `int` that was passed into the first function.

```go
first := factory(15)
first(2, 3, 4) // [30, 45, 60]

second := factory(2)
second(1, 2, 3, 4) // [2, 4, 6, 8]
```
