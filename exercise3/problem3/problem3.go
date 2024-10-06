package problem3

import (
	"errors"
)

type Set struct {
	elements map[any]struct{}
}

func NewSet() *Set {
	return &Set{
		elements: make(map[any]struct{}),
	}
}

func (s *Set) Add(value any) {
	s.elements[value] = struct{}{}
}

func (s *Set) Remove(value any) {
	delete(s.elements, value)
}

func (s *Set) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Set) Size() int {
	return len(s.elements)
}

func (s *Set) List() []any {
	list := make([]any, 0, len(s.elements))
	for key := range s.elements {
		list = append(list, key)
	}
	return list
}

func (s *Set) Has(value any) bool {
	_, exists := s.elements[value]
	return exists
}

func (s *Set) Copy() *Set {
	newSet := NewSet()
	for key := range s.elements {
		newSet.Add(key)
	}
	return newSet
}

func (s *Set) Difference(other *Set) *Set {
	diffSet := NewSet()
	for key := range s.elements {
		if !other.Has(key) {
			diffSet.Add(key)
		}
	}
	return diffSet
}

func (s *Set) IsSubset(other *Set) bool {
	for key := range s.elements {
		if !other.Has(key) {
			return false
		}
	}
	return true
}

func Union(sets ...*Set) *Set {
	unionSet := NewSet()
	for _, set := range sets {
		for key := range set.elements {
			unionSet.Add(key)
		}
	}
	return unionSet
}

func Intersect(sets ...*Set) *Set {
	if len(sets) == 0 {
		return NewSet()
	}

	intersectSet := sets[0].Copy()
	for _, set := range sets[1:] {
		for key := range intersectSet.elements {
			if !set.Has(key) {
				intersectSet.Remove(key)
			}
		}
	}
	return intersectSet
}

func (s *Set) Pop() (any, error) {
	if s.IsEmpty() {
		return nil, errors.New("Set is empty")
	}
	var key any
	for k := range s.elements {
		key = k
		break
	}
	s.Remove(key)
	return key, nil
}
