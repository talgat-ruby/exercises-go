package main

import "fmt"

func binary(a int) string {
	bi := ""
	if a == 0 {
		return "0"
	}
	for a > 0 {
		bit := a % 2
		bi += fmt.Sprintf("%d", bit)
		a /= 2
	}
	runedbi := []rune(bi)
	for i := 0; i < len(runedbi)/2; i++ {
		runedbi[i], runedbi[len(runedbi)-1-i] = runedbi[len(runedbi)-1-i], runedbi[i]
	}
	return string(runedbi)
}
