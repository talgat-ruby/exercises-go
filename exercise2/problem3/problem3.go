package problem3

import "slices"

type Set struct {
	values []any
}

func NewSet() *Set {
	return &Set{}
}

func (s *Set) Add(e any) {
	if !slices.Contains(s.values, e) {
		s.values = append(s.values, e)
	}
}

func (s *Set) Size() int {
	return len(s.values)
}

func (s *Set) IsEmpty() any {
	return len(s.values) == 0
}

func (s *Set) List() []any {
	return s.values
}

func (s *Set) Remove(element any) {
	var result []any
	for _, value := range s.values {
		if value != element {
			result = append(result, value)
		}
	}
	s.values = result
}

func (s *Set) Has(element any) bool {
	return slices.Contains(s.values, element)
}

func (s *Set) Copy() *Set {
	return &Set{s.values}
}

func (s *Set) Difference(B *Set) Set {
	var result Set
	for i, value := range s.values {
		if !B.Has(s.values[i]) {
			result.Add(value)
		}
	}
	return result
}

func (s *Set) IsSubset(S *Set) bool {
	for _, value := range s.values {
		return S.Has(value)
	}
	return true
}

func Union(S ...*Set) Set {
	var result Set
	for _, set := range S {
		for _, value := range set.values {
			result.Add(value)
		}
	}
	return result
}

func Intersect(S ...*Set) Set {
	var result Set
	union := Union(S...)

	for i := 0; i < len(union.values); i++ {
		isExistInEverySet := true
		for j := 0; j < len(S); j++ {
			if !S[j].Has(union.values[i]) {
				isExistInEverySet = false
			}
		}
		if isExistInEverySet {
			result.Add(union.values[i])
		}
	}
	return result
}
