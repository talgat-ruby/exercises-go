package problem6

type pipe func(in <-chan int) <-chan int

func multiplyBy2(in <-chan int) <-chan int {
	return in
}

func add5(in <-chan int) <-chan int {
	return in
}

func piper(in <-chan int, pipes []pipe) <-chan int {
	return in
}
