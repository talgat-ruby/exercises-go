# Problem 5

Create a method in the `Person` struct which returns how another person's age compares. Format is **{other person name}
is {older than / younger than / the same age as} me.**

```go
p1 := &Person{"Samuel", 24}
p2 := &Person{"Joel", 36}
p3 := &Person{"Lily", 24}

p1.compareAge(p2) // "Joel is older than me."
p2.compareAge(p1) // "Samuel is younger than me."
p1.compareAge(p3) // "Lily is the same age as me."
```
