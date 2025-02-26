package problem6

type pipe func(in <-chan int) <-chan int

var multiplyBy2 pipe = func(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v * 2
		}
		close(out)
	}()
	return out
}

var add5 pipe = func(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v + 5
		}
		close(out)
	}()
	return out
}

func piper(in <-chan int, pipes []pipe) <-chan int {
	var out <-chan int
	for _, pipe := range pipes {
		out = pipe(in)
		in = out
	}
	return out
}
