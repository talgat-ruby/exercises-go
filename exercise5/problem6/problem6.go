package problem6

type pipe func(in <-chan int) <-chan int

var multiplyBy2 pipe = func(in <-chan int) <-chan int {
	ot := make(chan int)
	go func() {
		for i := range in {
			ot <- i * 2
		}
		defer close(ot)
	}()
	return ot
}

var add5 pipe = func(in <-chan int) <-chan int {
	ot := make(chan int)
	go func() {
		for i := range in {
			ot <- i + 5
		}
		defer close(ot)
	}()
	return ot
}

func piper(in <-chan int, pipes []pipe) <-chan int {
	var ot <-chan int
	for _, pipe := range pipes {
		ot = pipe(in)
		in = ot
	}
	return ot
}
