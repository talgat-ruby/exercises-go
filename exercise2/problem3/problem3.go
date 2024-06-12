package problem3

type Set struct {
	elements map[any]struct{}
}

func NewSet() *Set {
	return &Set{elements: make(map[any]struct{})}
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
	for elem := range s.elements {
		list = append(list, elem)
	}
	return list
}
func (s *Set) Has(value any) bool {
	_, ok := s.elements[value]
	return ok
}
func (s *Set) Copy() *Set {
	newSet := NewSet()
	for elem := range s.elements {
		newSet.Add(elem)
	}
	return newSet
}
func (s *Set) Difference(other *Set) *Set {
	newSet := NewSet()
	for elem := range s.elements {
		if !other.Has(elem) {
			newSet.Add(elem)
		}
	}
	return newSet
}
func (s *Set) IsSubset(other *Set) bool {
	for elem := range s.elements {
		if !other.Has(elem) {
			return false
		}
	}
	return true
}
func Union(sets ...*Set) *Set {
	newSet := NewSet()
	for _, set := range sets {
		for elem := range set.elements {
			newSet.Add(elem)
		}
	}
	return newSet
}
func Intersect(sets ...*Set) *Set {
	if len(sets) == 0 {
		return NewSet()
	}
	newSet := NewSet()
	for elem := range sets[0].elements {
		isCommon := true
		for _, set := range sets[1:] {
			if !set.Has(elem) {
				isCommon = false
				break
			}
		}
		if isCommon {
			newSet.Add(elem)
		}
	}
	return newSet
}
