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
	for item := range s.items {
		list = append(list, item)
	}
	return list
}

func (s *Set) Has(input any) bool {
	_, found := s.items[input]
	return found
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
	for item := range s.items {
		if !other.Has(item) {
			diff.Add(item)
		}
	}
	return diff
}

func (s *Set) IsSubset(other *Set) bool {
	for item := range s.items {
		if !other.Has(item) {
			return false
		}
	}
	return true
}

func Union(sets ...*Set) *Set {
	union := NewSet()
	for _, set := range sets {
		for item := range set.items {
			union.Add(item)
		}
	}
	return union
}

func Intersect(sets ...*Set) *Set {
	intersect := NewSet()
	if len(sets) == 0 {
		return intersect
	}
	first := sets[0]
	for item := range first.items {
		foundInAll := true
		for _, set := range sets[1:] {
			if !set.Has(item) {
				foundInAll = false
				break
			}
		}
		if foundInAll {
			intersect.Add(item)
		}
	}
	return intersect
}
