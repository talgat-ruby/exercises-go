package problem3

import (
	"slices"
)

type Set struct {
	vals []any
	size int
}

func NewSet() *Set {
	return &Set{[]any{}, 0}
}

func (s *Set) Add(val any) {
	if !s.Has(val) {
		s.vals = append(s.vals, val)
		s.size += 1
	}
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

func (s Set) Copy() *Set {
	newVals := make([]any, len(s.vals))
	copy(newVals, s.vals)
	return &Set{vals: newVals, size: s.size}
}

func (s *Set) Difference(data *Set) Set {
	result := s.Copy()
	for _, v := range s.vals {
		if data.Has(v) {
			result.Remove(v)
		}
	}
	return *result
}

func (s *Set) IsSubset(data *Set) bool {
	for _, v := range s.vals {
		if !data.Has(v) {
			return false
		}
	}
	return true
}

func Union(data ...*Set) Set {
	slice := NewSet()
	for _, s := range data {
		for _, v := range s.vals {
			if !slice.Has(v) {
				slice.Add(v)
			}
		}
	}
	return *slice
}

func Intersect(data ...*Set) Set {
	if len(data) == 0 {
		return Set{}
	}

	slice := NewSet()
	for _, v := range data[0].vals {
		foundInAll := true
		for _, s := range data[1:] {
			if !s.Has(v) {
				foundInAll = false
				break
			}
		}
		if foundInAll {
			slice.Add(v)
		}
	}
	return *slice
}
