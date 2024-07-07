package problem3

type Set struct {
	values map[any]int
}

func NewSet() *Set {
	return &Set{values: map[any]int{}}
}

func (s *Set) Add(e any) {
	s.values[e] = 0
}

func (s *Set) Remove(e any) {
	delete(s.values, e)
}

func (s *Set) Size() int {
	return len(s.values)
}

func (s *Set) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Set) List() []any {
	l := make([]any, 0, s.Size())

	for k := range s.values {
		l = append(l, k)
	}

	return l
}

func (s *Set) Has(e any) bool {
	if _, ok := s.values[e]; ok {
		return true
	}
	return false
}

func (s *Set) Copy() *Set {
	cS := Set{values: map[any]int{}}

	for k, v := range s.values {
		cS.values[k] = v
	}

	return &cS
}

func (s *Set) Difference(otherSet *Set) Set {
	cS := Set{values: map[any]int{}}

	for k, v := range s.values {
		if !otherSet.Has(k) {
			cS.values[k] = v
		}
	}

	return cS
}

func (s *Set) IsSubset(otherSet *Set) bool {
	for k := range s.values {
		if !otherSet.Has(k) {
			return false
		}
	}

	return true
}

func Union(ss ...*Set) Set {
	cS := Set{values: map[any]int{}}

	for _, oS := range ss {
		for k, v := range oS.values {
			if !cS.Has(k) {
				cS.values[k] = v
			}
		}
	}

	return cS
}

func Intersect(ss ...*Set) *Set {
	firstSet := ss[0]
	ss = ss[1:]

	for _, oS := range ss {
		for k := range firstSet.values {
			if !oS.Has(k) {
				firstSet.Remove(k)
			}
		}
	}

	return firstSet
}
