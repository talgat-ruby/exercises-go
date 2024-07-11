package problem6

type pipe func(in <-chan int) <-chan int

func multiplyBy2(in <-chan int) <-chan int {
	out := make(chan int, len(in))
	go func() {
		for n := range in {
			out <- n * 2
		}
		close(out)
	}()
	return out
}

func add5(in <-chan int) <-chan int {
	out := make(chan int, len(in))
	go func() {
		for n := range in {
			out <- n + 5
		}
		close(out)
	}()
	return out
}

func piper(in <-chan int, pipes []pipe) <-chan int {
	for _, p := range pipes {
		in = p(in)
	}
	return in
}
