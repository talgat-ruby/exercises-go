package problem3

type Set struct {
	elements map[interface{}]struct{}
}

func NewSet() *Set {
	return &Set{elements: make(map[interface{}]struct{})}
}

func (s *Set) Add(element interface{}) {
	s.elements[element] = struct{}{}

}

func (s *Set) Remove(element interface{}) {
	delete(s.elements, element)

}

func (s *Set) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Set) Size() int {
	return len(s.elements)
}

func (s *Set) List() []interface{} {
	list := make([]interface{}, 0, s.Size())
	for element := range s.elements {
		list = append(list, element)
	}
	return list
}

func (s *Set) Has(element interface{}) bool {
	_, ok := s.elements[element]
	return ok
}
func (s *Set) Copy() *Set {
	newSet := NewSet()
	for element := range s.elements {
		newSet.Add(element)
	}
	return newSet
}

func (s *Set) Difference(other *Set) *Set {
	diffSet := NewSet()
	for element := range s.elements {
		if !other.Has(element) {
			diffSet.Add(element)
		}
	}
	return diffSet
}

//func (s *Set) Union(other *Set) *Set {
//	unionSet := NewSet()
//	for element := range s.elements {
//		unionSet.Add(element)
//	}
//	for element := range other.elements {
//		unionSet.Add(element)
//	}
//}

func (s *Set) IsSubset(other *Set) bool {
	for element := range s.elements {
		if !other.Has(element) {
			return false
		}
	}
	return true
}

func (s *Set) Union(other ...*Set) *Set {
	if s.IsEmpty() {
		return NewSet()
	}
	unionSet := NewSet()
	for element := range s.elements {
		isInAll := true
		for _, otherElement := range other {
			if otherElement.Has(element) {
				isInAll = false
				break
			}
		}
		if isInAll {
			unionSet.Add(element)
		}
	}
	return unionSet
}
