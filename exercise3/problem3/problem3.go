package problem3

type Set struct {
	elements []interface{}
}

func NewSet() *Set {
	return &Set{make([]interface{}, 0)}
}

func (s *Set) Add(element interface{}) {
	if !s.Has(element) {
		s.elements = append(s.elements, element)
	}
}

func (s *Set) Remove(element interface{}) {
	for i, e := range s.elements {
		if e == element {
			s.elements = append(s.elements[:i], s.elements[i+1:]...)
			break
		}
	}
}

func (s *Set) Has(element interface{}) bool {
	for _, e := range s.elements {
		if e == element {
			return true
		}
	}
	return false
}

func (s *Set) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Set) Size() int {
	return len(s.elements)
}

func (s *Set) List() []interface{} {
	list := make([]interface{}, len(s.elements))
	for i, e := range s.elements {
		list[i] = e
	}
	return list
}

func (s *Set) Copy() *Set {
	setCopy := NewSet()
	for _, e := range s.elements {
		setCopy.Add(e)
	}
	return setCopy
}

func (s *Set) Difference(other *Set) *Set {
	diffSet := NewSet()
	for _, e := range s.elements {
		if !other.Has(e) {
			diffSet.Add(e)
		}
	}
	return diffSet
}

func Union(sets ...*Set) *Set {
	unionSet := NewSet()
	for _, set := range sets {
		for _, e := range set.elements {
			unionSet.Add(e)
		}
	}
	return unionSet
}

func Intersect(sets ...*Set) *Set {
	if len(sets) == 0 {
		return NewSet()
	}

	intersectSet := NewSet()
	for _, e := range sets[0].elements {
		inAllSets := true
		for _, set := range sets[1:] {
			if !set.Has(e) {
				inAllSets = false
				break
			}
		}
		if inAllSets {
			intersectSet.Add(e)
		}
	}

	return intersectSet
}

func (s *Set) IsSubset(other *Set) bool {
	for _, e := range s.elements {
		if !other.Has(e) {
			return false
		}
	}
	return true
}
