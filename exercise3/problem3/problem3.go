package problem3

type Set struct {
	items map[any]struct{}
}

func NewSet() *Set {
	return &Set{
		items: make(map[any]struct{}),
	}
}

func (s *Set) Add(item any) {
	s.items[item] = struct{}{}
}

func (s *Set) Remove(item any) {
	delete(s.items, item)
}

func (s *Set) Has(item any) bool {
	_, exists := s.items[item]
	return exists
}

func (s *Set) Size() int {
	return len(s.items)
}

func (s *Set) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Set) List() []any {
	list := make([]any, 0, len(s.items))
	for item := range s.items {
		list = append(list, item)
	}
	return list
}

func (s *Set) Copy() *Set {
	copySet := NewSet()
	for item := range s.items {
		copySet.Add(item)
	}
	return copySet
}

func (s *Set) Difference(other *Set) *Set {
	result := NewSet()
	for item := range s.items {
		if !other.Has(item) {
			result.Add(item)
		}
	}
	return result
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
	result := NewSet()
	for _, set := range sets {
		for item := range set.items {
			result.Add(item)
		}
	}
	return result
}

func Intersect(sets ...*Set) *Set {
	if len(sets) == 0 {
		return NewSet()
	}

	result := NewSet()
	for item := range sets[0].items {
		intersects := true
		for _, set := range sets[1:] {
			if !set.Has(item) {
				intersects = false
				break
			}
		}
		if intersects {
			result.Add(item)
		}
	}

	return result
}
