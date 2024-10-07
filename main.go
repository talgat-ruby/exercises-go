package main

import (
	"fmt"
	"slices"
)

func main() {
	set1 := Set{vals: []any{1, 2, 3, 4}, size: 4}
	set2 := Set{vals: []any{1, 6, 7}, size: 2}
	diff := set1.Union(set2)
	fmt.Println(diff)
}

type Set struct {
	vals []any
	size int
}

func (s *Set) Add(val any) {
	s.vals = append(s.vals, val)
	s.size += 1
}

func (s *Set) Remove(val any) {
	key := slices.Index(s.vals, val)
	if key >= 0 {
		s.vals = slices.Delete(s.vals, key, key+1)
		s.size -= 1
	}
}

func (s *Set) IsEmpty() bool {
	return s.size == 0
}

func (s *Set) Size() int {
	return s.size
}

func (s *Set) List() []any {
	return s.vals
}

func (s *Set) Has(val any) bool {
	if slices.Contains(s.vals, val) {
		return true
	}
	return false
}

func (s *Set) Copy() []any {
	copySlice := make([]any, 0)
	copy(copySlice, s.vals)
	return copySlice
}

func (s *Set) Difference(data Set) Set {
	for _, v := range s.vals {
		if data.Has(v) {
			s.Remove(v)
		}
	}
	return *s
}

func (s *Set) IsSubset(data Set) bool {
	for _, v := range data.vals {
		if !s.Has(v) {
			return false
		}
	}
	return true
}

func (s *Set) Union(data Set) Set {
	for _, v := range data.vals {
		if !s.Has(v) {
			s.Add(v)
		}
	}
	return *s
}
