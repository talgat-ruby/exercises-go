# Problem 6

Create structs `Animal` and `Insect`. That will be accepted by `sumOfAllLegsNum` and calculate correct amount of legs

```go
horse := &Animal{
    name:    "horse",
    legsNum: 4,
}
kangaroo := &Animal{
    name:    "kangaroo",
    legsNum: 2,
}
ant := &Insect{
    name:    "ant",
    legsNum: 6,
}
spider := &Insect{
    name:    "spider",
    legsNum: 8,
}

sumOfAllLegsNum(horse, kangaroo, ant, spider) // 20
```
