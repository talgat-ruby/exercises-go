package problem3

type Set struct {
	data map[any]struct{}
}

func NewSet() *Set {
	return &Set{
		data: make(map[any]struct{}),
	}
}

func (s *Set) Add(v any) {
	s.data[v] = struct{}{}
}

func (s *Set) Remove(v any) {
	delete(s.data, v)
}

func (s *Set) Size() int {
	return len(s.data)
}

func (s *Set) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Set) List() []any {
	var list []any
	for k := range s.data {
		list = append(list, k)
	}
	return list
}

func (s *Set) Has(v any) bool {
	_, ok := s.data[v]
	return ok
}

func (s *Set) Copy() *Set {
	newSet := NewSet()
	for k := range s.data {
		newSet.Add(k)
	}
	return newSet
}

func (s *Set) Difference(other *Set) *Set {
	newSet := NewSet()
	for k := range s.data {
		if !other.Has(k) {
			newSet.Add(k)
		}
	}
	return newSet
}

func (s *Set) IsSubset(other *Set) bool {
	for k := range s.data {
		if !other.Has(k) {
			return false
		}
	}
	return true
}

func Union(sets ...*Set) *Set {
	newSet := NewSet()
	for _, s := range sets {
		for k := range s.data {
			newSet.Add(k)
		}
	}
	return newSet
}

func Intersect(sets ...*Set) *Set {
	newSet := NewSet()
	if len(sets) == 0 {
		return newSet
	}
	for k := range sets[0].data {
		has := true
		for _, s := range sets[1:] {
			if !s.Has(k) {
				has = false
				break
			}
		}
		if has {
			newSet.Add(k)
		}
	}
	return newSet
}
