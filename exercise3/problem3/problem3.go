package problem3

type Set struct {
	elements []any
}

func (s *Set) Has(el any) bool {
	var isHasElement bool
	for _, i := range s.elements {
		if i == el {
			isHasElement = true
		}
	}
	return isHasElement
}

func (s *Set) Add(el any) {
	if !s.Has(el) {
		s.elements = append(s.elements, el)
	}
}

func (s *Set) Remove(el any) {
	var isHas bool
	var num int
	for i, v := range s.elements {
		if v == el {
			num = i
			isHas = true
			break
		}
	}
	if isHas {
		s.elements = append(s.elements[:num], s.elements[num+1:]...)
	}
}

func (s *Set) Size() int {
	return len(s.elements)
}

func (s *Set) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Set) List() []any {
	return s.elements
}

func (s *Set) Copy() *Set {
	newSet := &Set{
		elements: make([]any, len(s.elements)),
	}
	copy(newSet.elements, s.elements)
	return newSet
}

func NewSet() *Set {
	return &Set{
		elements: []any{},
	}
}

func (s1 *Set) Difference(s2 *Set) *Set {
	newSet := NewSet()
	for _, el := range s1.elements {
		if !s2.Has(el) {
			newSet.Add(el)
		}
	}
	return newSet
}

func (s1 *Set) IsSubset(s2 *Set) bool {
	for _, el := range s1.elements {
		if !s2.Has(el) {
			return false
		}
	}
	return true
}

func Union(s ...*Set) *Set {
	newSet := NewSet()
	for _, i := range s {
		for _, j := range i.elements {
			if !newSet.Has(j) {
				newSet.Add(j)
			}
		}
	}
	return newSet
}

func Intersect(sets ...*Set) *Set {
	if len(sets) == 0 {
		return NewSet()
	}
	newSet := NewSet()
	for _, el := range sets[0].List() {
		hasSet := NewSet()
		hasSet.Add(el)
		isHas := true
		for _, set := range sets[1:] {
			if hasSet.Difference(set).Size() != 0 {
				isHas = false
				break
			}
		}
		if isHas {
			newSet.Add(el)
		}
	}

	return newSet
}
