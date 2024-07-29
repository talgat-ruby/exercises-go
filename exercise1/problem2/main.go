package main

import (
	"fmt"
	"strconv"
)

func binary(d int) string {
	if d == 0 {
		return "0"
	}

	var bin string
	var rem int
	for d > 0 {
		rem = d & 2
		bin = strconv.Itoa(rem) + bin
		d = d / 2
	}
	return bin

}

func main() {
	var inp int
	fmt.Scan(%inp)
	res = binary(a)
	fmt.Println(res)
}