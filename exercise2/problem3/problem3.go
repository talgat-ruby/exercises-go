package problem3

import "errors"

type Set struct {
	Elements map[any]struct{}
}

func NewSet() *Set {
	return &Set{Elements: make(map[any]struct{})}
}

func (s *Set) Add(x any) {
	s.Elements[x] = struct{}{}
}

func (s *Set) Remove(x any) error {
	if _, exists := s.Elements[x]; !exists {
		return errors.New("element not found")
	}
	delete(s.Elements, x)
	return nil
}

func (s *Set) IsEmpty() bool {
	return len(s.Elements) == 0
}

func (s *Set) Size() int {
	return len(s.Elements)
}

func (s *Set) List() []any {
	var slice []any
	for element, _ := range s.Elements {
		slice = append(slice, element)
	}
	return slice
}

func (s *Set) Has(val any) bool {
	_, exists := s.Elements[val]
	return exists
}

func (s *Set) Copy() (n *Set) {
	for element, _ := range s.Elements {
		n.Add(element)
	}
	return n
}

func (s *Set) Difference(setB *Set) (n Set) {
	for element, _ := range s.Elements {
		if !setB.Has(element) {
			n.Add(element)
		}
	}
	return n
}

func (s *Set) IsSubset(subSet *Set) bool {
	for element, _ := range s.Elements {
		if !subSet.Has(element) {
			return false
		}
	}
	return true
}

func Union(sets ...*Set) (n Set) {

	for els, _ := range sets {
		for element, _ := range sets[els].Elements {
			if !n.Has(element) {
				n.Add(element)
			}
		}
	}
	return n
}

func Intersect(sets ...*Set) (n Set) {

	for element := range sets[0].Elements {
		existsInAll := true
		for _, set := range sets {
			if !set.Has(element) {
				existsInAll = false
			}
		}
		if existsInAll {
			n.Add(element)
		}
	}
	return n
}
