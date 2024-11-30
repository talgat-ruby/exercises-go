package problem6

type pipe func(in <-chan int) <-chan int

//var multiplyBy2 pipe = func(in <-chan int) <-chan int {
//	output := make(chan int)
//	defer close(output)
//
//	for n := range in {
//		output <- n * 2
//	}
//
//	return output
//}
//
//var add5 pipe = func(in <-chan int) <-chan int {
//	output := make(chan int)
//	defer close(output)
//
//	for n := range in {
//		output <- n + 5
//	}
//
//	return output
//}

var multiplyBy2 pipe = func(in <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for n := range in {
			output <- n * 2
		}
	}()
	return output
}

var add5 pipe = func(in <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for n := range in {
			output <- n + 5
		}
	}()
	return output
}

func piper(in <-chan int, pipes []pipe) <-chan int {
	output := in
	for _, pipe := range pipes {
		output = pipe(output)
	}

	return output

}
