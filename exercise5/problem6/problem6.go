package problem6

type pipe func(in <-chan int) <-chan int

var multiplyBy2 pipe = func(in <-chan int) <-chan int {
	return channelIterator(in, func(n int) int { return n * 2 })
}

var add5 pipe = func(in <-chan int) <-chan int {
	return channelIterator(in, func(n int) int { return n + 5 })
}

func channelIterator(in <-chan int, f func(n int) int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- f(n)
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
