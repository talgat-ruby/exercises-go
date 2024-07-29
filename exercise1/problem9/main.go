package main

import (
	"strconv"
)

func bitwiseAND(a int, b int) int {
	joinab := ""
	ares,bres := backJob(a,b)

	for i := range bres {
		if bres[i] == '1' && ares[i] == '1' {
			joinab += "1"
		} else {
			joinab += "0"
		}
	}

	num, err := strconv.ParseInt(joinab, 2, 64)
	if err != nil {
		return 0
	}
	res := int(num)
	return res

}

func bitwiseOR(a int, b int) int {
	joinab := ""
	ares,bres := backJob(a,b)

	for i := range bres {
		if bres[i] == '1' || ares[i] == '1' {
			joinab += "1"
		} else {
			joinab += "0"
		}
	}
	num, err := strconv.ParseInt(joinab, 2, 64)
	if err != nil {
		return 0
	}
	res := int(num)
	return res
}

func bitwiseXOR(a int, b int) int {
	joinab := ""
	ares,bres := backJob(a,b)

	for i := range bres {
		if bres[i] != ares[i] {
			joinab += "1"
		} else {
			joinab += "0"
		}
	}
	num, err := strconv.ParseInt(joinab, 2, 64)
	if err != nil {
		return 0
	}
	res := int(num)
	return res

}

func backJob(a int, b int) (string, string) {
	ares := ""
	count := 0
	for a != 0 {
		k := a % 2
		c := strconv.Itoa(k)
		ares = c + ares
		a = a / 2
	}
	bres := ""
	for b != 0 {
		k := b % 2
		c := strconv.Itoa(k)
		bres = c + bres
		b = b / 2
	}

	if len(bres) > len(ares) {
		count = len(bres) - len(ares)
		for i := 0; i < count; i++ {
			ares = "0" + ares
		}
	} else {
		count = len(ares) - len(bres)
		for i := 0; i < count; i++ {
			bres = "0" + bres
		}
	}
	return ares, bres

}