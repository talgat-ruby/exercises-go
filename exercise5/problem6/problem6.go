package problem6

type pipe func(in <-chan int) <-chan int

var multiplyBy2 pipe = func() {}

var add5 pipe = func() {}

func piper(in <-chan int, pipes []pipe) <-chan int {}
