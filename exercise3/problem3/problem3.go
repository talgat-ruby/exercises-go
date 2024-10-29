package problem3

type Set struct{
	Elements []any
}

func (s *Set)Add(element any) {
	if !s.Has(element) {
	s.Elements = append(s.Elements, element)
	}
}

func (s *Set)Remove(element any) {
	for k, i := range s.Elements {
		if i == element {
			s.Elements =append(s.Elements[:k], s.Elements[k+1:]... )
		}
	}
}

func (s *Set)IsEmpty() bool {
	return len(s.Elements) == 0
}

func (s *Set)Size() int {
	return len(s.Elements)
}

func (s *Set)List() []any {
	elements := []any{}
	elements = append(elements, s.Elements...)
	return elements
}

func (s *Set)Has(element any) bool {
	for _, i := range s.Elements {
		if i == element {
			return true
		}
	}
	return false
}

func (s *Set)Copy() *Set {
	return s
}

func (s *Set)Difference(elements *Set) *Set {
	element := NewSet()
	
	for _, i := range s.Elements {
		count := 0
		for _, j := range elements.Elements {
			if i == j {
				count++
			}
		}
		if count == 0 {
			element.Elements = append(element.Elements, i)
		}
	}

	return element
}

func (s *Set)IsSubset(elements *Set) bool {

	for _, i := range s.Elements {
		count := 0
		for _, j := range elements.Elements {
			if i == j {
				count++
			}
		}
		if count == 0 {
			return false
		}
	}

	return true
}

func Union(s ...*Set) *Set {
	newElements := &Set{}

	for _, i:= range s {
		for _, j := range i.Elements {
			if !newElements.Has(j) {
				newElements.Elements = append(newElements.Elements, j) 
			}
		}
	}
	return newElements
}

func Intersect(s ...*Set) *Set {
	newElements := &Set{}

	s1 := s[0]
	s2 := s[1:]

	for _, i:= range s1.Elements {
		count1 := 0
		for _, j := range s2 {
			count2 := 0
			for _,  k := range j.Elements {
				if i == k {
					count2++
				}
			}
			if count2 != 0 {
				count1++
			}
		}
		if count1 == len(s2) {
			newElements.Elements = append(newElements.Elements, i)
		}
	}
	return newElements
}

func NewSet() *Set {
    return &Set{
        Elements: []any{}, 
    }
}


