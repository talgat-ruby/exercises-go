package problem3

type Set struct {
	items map[any]struct{}
}

func NewSet() *Set {
	return &Set{items: make(map[any]struct{})}
}

func (s *Set) Add(val any) {
	s.items[val] = struct{}{}
}

func (s *Set) Remove(val any) {
	delete(s.items, val)
}

func (s *Set) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Set) Size() int {
	return len(s.items)
}

func (s *Set) List() []any {
	list := make([]any, 0, len(s.items))
	for k := range s.items {
		list = append(list, k)
	}
	return list
}

func (s *Set) Has(val any) bool {
	_, exists := s.items[val]
	return exists
}

func (s *Set) Copy() *Set {
	newSet := NewSet()
	for k := range s.items {
		newSet.Add(k)
	}
	return newSet
}

func (s *Set) Difference(other *Set) *Set {
	diff := NewSet()
	for k := range s.items {
		if !other.Has(k) {
			diff.Add(k)
		}
	}
	return diff
}

func (s *Set) IsSubset(other *Set) bool {
	for k := range s.items {
		if !other.Has(k) {
			return false
		}
	}
	return true
}

func Union(sets ...*Set) *Set {
	unionSet := NewSet()
	for _, set := range sets {
		for k := range set.items {
			unionSet.Add(k)
		}
	}
	return unionSet
}

func Intersect(sets ...*Set) *Set {
	if len(sets) == 0 {
		return NewSet()
	}

	intersectSet := sets[0].Copy()
	for _, set := range sets[1:] {
		for k := range intersectSet.items {
			if !set.Has(k) {
				intersectSet.Remove(k)
			}
		}
	}
	return intersectSet
}
