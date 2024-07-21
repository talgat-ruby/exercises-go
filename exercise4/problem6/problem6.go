package problem6

type pipe func(in <-chan int) <-chan int

// multiplyBy2 умножает каждое число на 2
func multiplyBy2(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * 2
		}
	}()
	return out
}

// add5 добавляет 5 к каждому числу
func add5(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n + 5
		}
	}()
	return out
}

// piper комбинирует несколько операций (pipes)
func piper(in <-chan int, pipes []pipe) <-chan int {
	for _, p := range pipes {
		in = p(in)
	}
	return in
}
