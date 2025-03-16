package problem3

import (
	"errors"
)

type Set struct {
	Elements map[any]struct{}
}

func NewSet() *Set {
	var set Set
	set.Elements = make(map[any]struct{})
	return &set
}

func (s *Set) Add(elem any) {
	s.Elements[elem] = struct{}{}
}

func (s *Set) Remove(elem any) error {
	if _, exists := s.Elements[elem]; !exists {
		return errors.New("Stack is empty")
	}
	delete(s.Elements, elem)
	return nil
}

func (s *Set) IsEmpty() bool {
	return len(s.Elements) == 0
}
func (s *Set) Size() int {
	return len(s.Elements)
}
func (s *Set) List() []any {
	list := make([]any, 0, len(s.Elements))
	for elem := range s.Elements {
		list = append(list, elem)
	}
	return list
}
func (s *Set) Has(elem any) bool {
	_, exists := s.Elements[elem]
	return exists
}

func (s *Set) Copy() *Set {
	newSet := NewSet()
	for elem := range s.Elements {
		newSet.Elements[elem] = struct{}{}
	}
	return newSet
}
func (s *Set) Difference(other *Set) *Set {
	difference := NewSet()
	for elem := range s.Elements {
		if !other.Has(elem) {
			difference.Add(elem)
		}
	}
	return difference
}

func (s *Set) IsSubset(other *Set) bool {
	if s.Size() > other.Size() {
		return false
	}
	return s.Difference(other).Size() == 0

}

func Union(sets ...*Set) *Set {
	unionSet := NewSet()

	for _, set := range sets {
		for elem := range set.Elements {
			unionSet.Add(elem)
		}
	}
	return unionSet
}

func Intersect(sets ...*Set) *Set {
	if len(sets) == 0 {
		return NewSet()
	}
	intersection := sets[0].Copy()
	for _, set := range sets[1:] {
		for elem := range intersection.Elements {
			if !set.Has(elem) {
				intersection.Remove(elem)
			}
		}
	}
	return intersection
}
