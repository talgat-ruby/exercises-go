package problem3

type Set struct {
	items map[interface{}]struct{}
}

func NewSet() *Set {
	return &Set{
		items: make(map[interface{}]struct{}),
	}
}

func (s *Set) Add(item interface{}) {
	s.items[item] = struct{}{}
}

func (s *Set) Remove(item interface{}) {
	delete(s.items, item)
}

func (s *Set) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Set) Size() int {
	return len(s.items)
}

func (s *Set) Has(item interface{}) bool {
	_, exists := s.items[item]
	return exists
}

func (s *Set) Copy() *Set {
	resultSet := NewSet()
	for key := range s.items {
		resultSet.Add(key)
	}
	return resultSet
}

func (s *Set) Difference(r *Set) *Set {
	resultSet := NewSet()
	for key := range s.items {
		if !r.Has(key) {
			resultSet.Add(key)
		}
	}
	return resultSet
}

func (s *Set) IsSubset(r *Set) bool {
	for key := range s.items {
		if !r.Has(key) {
			return false
		}
	}
	return true
}

func (s *Set) List() []interface{} {
	result := []interface{}{}
	for key := range s.items {
		result = append(result, key)
	}
	return result
}

func Union(sets ...*Set) *Set {
	resultSet := NewSet()
	for _, set := range sets {
		for key := range set.items {
			resultSet.Add(key)
		}
	}
	return resultSet
}

func Intersect(sets ...*Set) *Set {
	if len(sets) == 0 {
		return NewSet()
	}

	minSet := sets[0]
	for _, set := range sets[1:] {
		if set.Size() < minSet.Size() {
			minSet = set
		}
	}

	resultSet := NewSet()
	for key := range minSet.items {
		isCommon := true
		for _, set := range sets {
			if !set.Has(key) {
				isCommon = false
				break
			}
		}
		if isCommon {
			resultSet.Add(key)
		}
	}

	return resultSet
}
