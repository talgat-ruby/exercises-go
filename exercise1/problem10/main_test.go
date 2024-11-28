package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	table := []struct {
		a, b string
		exp  string
		err  error
	}{
		{a: "1", b: "2", exp: "3", err: nil},
		{a: "10", b: "20", exp: "30", err: nil},
		{a: "a", b: "2", exp: "", err: fmt.Errorf("string: a cannot be converted")},
		{a: "1", b: "2c", exp: "", err: fmt.Errorf("string: 2c cannot be converted")},
	}

	for _, r := range table {
		out, err := sum(r.a, r.b)
		if out != r.exp || (err != nil && errors.Is(err, r.err)) {
			t.Errorf("sum(%s, %s) was incorrect, expected: %s, %v, got: %s, %v.", r.a, r.b, r.exp, r.err, out, err)
		}
	}
}
