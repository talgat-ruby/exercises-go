package problem6

type pipe func(in <-chan int) <-chan int

var multiplyBy2 pipe = func(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * 2
		}
		close(out)
	}()
	return out
}

var add5 pipe = func(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n + 5
		}
		close(out)
	}()
	return out
}

func piper(in <-chan int, pipes []pipe) <-chan int {
	for _, pipe := range pipes {
		in = pipe(in)
	}
	return in
}
