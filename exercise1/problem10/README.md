# Problem 10

Sum two strings. If any of the strings cannot converted return an error as a second parameter:
`string: %value% cannot be converted`

```go
sum("1", "2") // "3", nil

sum("10", "20") // "30", nil

sum("a", "2") // "", "string: a cannot be converted"
```
