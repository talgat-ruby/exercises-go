package problem6

type pipe func(in <-chan int) <-chan int

var multiplyBy2 pipe = func(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * 2
		}
	}()
	return out
}

var add5 pipe = func(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n + 5
		}
	}()
	return out
}

func piper(in <-chan int, pipes []pipe) <-chan int {
	output := in
	for _, p := range pipes {
		output = p(output)
	}
	return output
}
