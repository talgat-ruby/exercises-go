package problem3

import (
	"slices"
)

type Set struct {
	items []any
}

func NewSet() *Set {
	return &Set{
		items: make([]any, 0),
	}
}

func (s *Set) Add(value any) {
	if !s.Has(value) {
		s.items = append(s.items, value)
	}
}

func (s *Set) Remove(value any) {
	for i, v := range s.items {
		if v == value {
			s.items = append(s.items[:i], s.items[i+1:]...)
		}
	}
}

func (s *Set) Size() int {
	return len(s.items)
}

func (s *Set) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Set) Has(value any) bool {
	return slices.Contains(s.items, value)
}

func (s *Set) List() []any {
	return s.items
}

func (s *Set) Copy() *Set {
	nSet := s

	return nSet
}

func (s *Set) Difference(B *Set) *Set {
	diff := NewSet()

	for _, value := range s.items {
		if !B.Has(value) {
			diff.Add(value)
		}
	}

	return diff
}

func (s *Set) IsSubset(B *Set) bool {
	for _, value := range s.items {
		if !B.Has(value) {
			return false
		}
	}

	return true
}

func Union(sets ...*Set) *Set {
	s := NewSet()

	for _, set := range sets {
		for _, value := range set.List() {
			if !s.Has(value) {
				s.Add(value)
			}
		}
	}

	return s
}

func Intersect(sets ...*Set) *Set {
	inter := NewSet()

	for _, set := range sets {
		if inter.IsEmpty() {
			inter = set.Copy()

			continue
		}

		diff := inter.Difference(set)

		for _, value := range diff.List() {
			inter.Remove(value)
		}
	}

	return inter
}
