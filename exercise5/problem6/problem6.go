package problem6

type pipe func(in <-chan int) <-chan int

var multiplyBy2 pipe = func(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range in {
			out <- num * 2
		}
		close(out)
	}()
	return out
}

var add5 pipe = func(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range in {
			out <- num + 5
		}
		close(out)
	}()
	return out
}

func piper(in <-chan int, pipes []pipe) <-chan int {
	for _, p := range pipes {
		in = p(in) // Применяем текущую pipe функцию
	}
	return in // Возвращаем канал с окончательным результатом
}
