package problem3

type Set struct {
	elements []any
}

func NewSet() Set {
	return Set{}
}

func (s *Set) Add(x any) {
	if !s.Has(x) {
		s.elements = append(s.elements, x)
	}
}

func (s *Set) Size() int {
	return len(s.elements)
}

func (s *Set) Remove(x any) {
	if s.Size() != 0 {
		for i, el := range s.elements {
			if el == x {
				s.elements = append(s.elements[:i], s.elements[i+1:]...)
			}
		}
	}
}

func (s *Set) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Set) List() []any {
	return s.elements
}

func (s *Set) Has(x any) bool {
	for _, el := range s.elements {
		if el == x {
			return true
		}
	}

	return false
}

func (s *Set) Copy() Set {
	elements := make([]any, len(s.elements))
	copy(elements, s.elements)

	return Set{elements: elements}
}

func (s *Set) Difference(other Set) Set {
	diff := NewSet()

	for _, el := range s.elements {
		if !other.Has(el) {
			diff.Add(el)
		}
	}

	return diff
}

func (s *Set) IsSubset(other Set) bool {
	for _, el := range s.elements {
		if !other.Has(el) {
			return false
		}
	}

	return true
}

func Union(sets ...Set) Set {
	union := Set{}

	for _, elements := range sets {
		for _, el := range elements.List() {
			if !union.Has(el) {
				union.Add(el)
			}
		}
	}

	return union
}

func Intersect(sets ...Set) Set {
	intersection := Set{}

	for _, el := range sets[0].List() {
		count := 1

		for i := 1; i < len(sets); i++ {
			set := sets[i]

			if set.Has(el) {
				count++
			}
		}

		if count == len(sets) {
			intersection.Add(el)
		}
	}

	return intersection
}
