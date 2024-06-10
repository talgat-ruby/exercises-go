package problem3

type Set struct {
	m    map[any]bool
	size int
}

func NewSet() *Set {
	return &Set{make(map[any]bool), 0}
}

func (s *Set) Add(value any) {
	if !s.Has(value) {
		s.m[value] = true
		s.size++
	}
}

func (s *Set) Remove(value any) {
	if s.Has(value) {
		delete(s.m, value)
		s.size--
	}
}

func (s *Set) Size() int {
	return s.size
}

func (s *Set) IsEmpty() bool {
	return s.size == 0
}

func (s *Set) Copy() *Set {
	set := &Set{make(map[any]bool), s.size}
	for k := range s.m {
		set.m[k] = true
	}
	return set
}

func (s *Set) Has(value any) bool {
	_, ok := s.m[value]
	return ok
}

func (s *Set) List() []any {
	out := make([]any, 0, s.size)
	for k := range s.m {
		out = append(out, k)
	}
	return out
}

func (s *Set) IsSubset(other *Set) bool {
	for k := range s.m {
		if _, ok := other.m[k]; !ok {
			return false
		}
	}
	return true
}

func (s *Set) Difference(other *Set) *Set {
	out := NewSet()
	for k := range s.m {
		if !other.Has(k) {
			out.m[k] = true
			out.size++
		}
	}
	return out
}

func Union(sets ...*Set) *Set {
	out := NewSet()
	for _, set := range sets {
		for k := range set.m {
			out.Add(k)
		}
	}
	return out
}

func Intersect(sets ...*Set) *Set {
	out := Union(sets...)
	for k := range out.m {
		for _, set := range sets {
			if !set.Has(k) {
				out.Remove(k)
				break
			}
		}
	}
	return out
}
