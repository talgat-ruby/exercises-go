package problem6

type pipe func(in <-chan int) <-chan int

var multiplyBy2 pipe = func(in <-chan int) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)
		for val := range in {
			result <- val * 2
		}
	}()
	return result
}

var add5 pipe = func(in <-chan int) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)
		for val := range in {
			result <- val + 5
		}
	}()
	return result
}

func piper(in <-chan int, pipes []pipe) <-chan int {
	for _, p := range pipes {
		in = p(in)
	}
	return in
}
