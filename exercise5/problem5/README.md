# Problem 5

`send` combines words into a single message. It accepts slice of strings which needs to combine, producer and consumer.

Your job is to implement correct `producer` and `consumer`.

```go
list := []string{
"Hello",
"dear",
"friend!",
"Learn",
"from",
"yesterday.",
"Save",
"our",
"soles.",
}

send(list, producer, consumer) // "Hello dear friend! Learn from yesterday. Save our soles."
```
