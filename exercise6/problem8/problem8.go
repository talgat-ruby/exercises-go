package problem8

import (
	"reflect"
)

func multiplex(channels []<-chan string) []string {
	var result []string
	done := make(chan struct{})

	go func() {
		cases := make([]reflect.SelectCase, len(channels))
		for i, ch := range channels {
			cases[i] = reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(ch),
			}
		}

		for len(cases) > 0 {
			chosen, value, ok := reflect.Select(cases)
			if ok {
				result = append(result, value.String())
			} else {
				cases = append(cases[:chosen], cases[chosen+1:]...)
			}
		}

		close(done)
	}()

	<-done
	return result
}
