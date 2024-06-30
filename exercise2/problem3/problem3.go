package problem3

import "sync"

type Set struct {
	mu   sync.RWMutex
	data map[interface{}]struct{}
}

func NewSet() *Set {
	return &Set{
		data: make(map[interface{}]struct{}),
	}
}

func (s *Set) Add(elem interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[elem] = struct{}{}
}

func (s *Set) Remove(elem interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, elem)
}

func (s *Set) IsEmpty() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.data) == 0
}

func (s *Set) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.data)
}

func (s *Set) List() []interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()
	l := make([]interface{}, 0, len(s.data))
	for k := range s.data {
		l = append(l, k)
	}
	return l
}

func (s *Set) Has(elem interface{}) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, exists := s.data[elem]
	return exists
}

func (s *Set) Copy() *Set {
	s.mu.RLock()
	defer s.mu.RUnlock()
	copySet := NewSet()
	for k := range s.data {
		copySet.data[k] = struct{}{}
	}
	return copySet
}

func (s *Set) Difference(other *Set) *Set {
	s.mu.RLock()
	defer s.mu.RUnlock()
	other.mu.RLock()
	defer other.mu.RUnlock()

	diffSet := NewSet()
	for k := range s.data {
		if _, exists := other.data[k]; !exists {
			diffSet.data[k] = struct{}{}
		}
	}
	return diffSet
}

func (s *Set) IsSubset(other *Set) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	other.mu.RLock()
	defer other.mu.RUnlock()

	for k := range s.data {
		if _, exists := other.data[k]; !exists {
			return false
		}
	}
	return true
}

func Union(sets ...*Set) *Set {
	unionSet := NewSet()
	for _, set := range sets {
		set.mu.RLock()
		for k := range set.data {
			unionSet.data[k] = struct{}{}
		}
		set.mu.RUnlock()
	}
	return unionSet
}

func Intersect(sets ...*Set) *Set {
	if len(sets) == 0 {
		return NewSet()
	}
	intersectSet := sets[0].Copy()
	for _, set := range sets[1:] {
		set.mu.RLock()
		intersectSet.mu.Lock()
		for k := range intersectSet.data {
			if _, exists := set.data[k]; !exists {
				delete(intersectSet.data, k)
			}
		}
		set.mu.RUnlock()
		intersectSet.mu.Unlock()
	}
	return intersectSet
}
