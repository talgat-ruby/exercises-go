package problem6

type pipe func(in <-chan int) <-chan int

var multiplyBy2 pipe = func(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			out <- num * 2
		}
	}()
	return out
}

var add5 pipe = func(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			out <- num + 5
		}
	}()
	return out
}

func piper(in <-chan int, pipes []pipe) <-chan int {
	out := in
	for _, p := range pipes {
		out = p(out)
	}
	return out
}
