package problem3

type Set struct {
	items map[any]struct{}
}

func NewSet() *Set {
	return &Set{items: make(map[any]struct{})}
}

func (s *Set) Add(value any) {
	s.items[value] = struct{}{}
}

func (s *Set) Remove(value any) {
	delete(s.items, value)
}

func (s *Set) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Set) Size() int {
	return len(s.items)
}

func (s *Set) List() []any {
	list := make([]any, 0, len(s.items))
	for key := range s.items {
		list = append(list, key)
	}
	return list
}

func (s *Set) Has(value any) bool {
	_, exists := s.items[value]
	return exists
}

func (s *Set) Copy() *Set {
	newSet := NewSet()
	for key := range s.items {
		newSet.Add(key)
	}
	return newSet
}

// Difference returns a new Set containing elements in Set A but not in Set B
func (s *Set) Difference(other *Set) *Set {
	diffSet := NewSet()
	for key := range s.items {
		if !other.Has(key) {
			diffSet.Add(key)
		}
	}
	return diffSet
}

// IsSubset checks if Set B is a subset of Set A
func (s *Set) IsSubset(other *Set) bool {
	for key := range s.items {
		if !other.Has(key) {
			return false
		}
	}
	return true
}

// Union returns a new Set that consists of elements in either Set
func Union(sets ...*Set) *Set {
	result := NewSet()
	for _, set := range sets {
		for key := range set.items {
			result.Add(key)
		}
	}
	return result
}

// Intersect returns a new Set that consists of elements in all Sets
func Intersect(sets ...*Set) *Set {
	if len(sets) == 0 {
		return NewSet()
	}

	result := NewSet()
	for key := range sets[0].items {
		isCommon := true
		for _, set := range sets[1:] {
			if !set.Has(key) {
				isCommon = false
				break
			}
		}
		if isCommon {
			result.Add(key)
		}
	}
	return result
}
