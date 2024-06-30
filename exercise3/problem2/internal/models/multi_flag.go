package models

import "fmt"

type MultiFlag []string

// Set implements flag.Value.
func (m *MultiFlag) Set(value string) error {
	*m = append(*m, value)
	return nil
}

// String implements flag.Value.
func (m *MultiFlag) String() string {
	return fmt.Sprintf("%v", *m)
}
