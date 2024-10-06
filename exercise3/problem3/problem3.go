package problem3

import "fmt"

type Set struct {
	values map[any]bool
}

func NewSet() *Set {
	return &Set{
		values: make(map[any]bool),
	}
}

func (s *Set) Add(key any) {
	s.values[key] = true
}
func (s *Set) Remove(key any) {
	delete(s.values, key)
}

func (s *Set) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Set) Size() int {
	return len(s.values)
}

func (s *Set) List() []any {
	keys := make([]any, 0, s.Size())
	for k := range s.values {
		keys = append(keys, k)
	}
	return keys
}

func (s *Set) Has(key any) bool {
	_, ok := s.values[key]
	return ok
}
func (s *Set) Copy() *Set {
	copies := make(map[any]bool)
	for k, v := range s.values {
		copies[k] = v
	}
	return &Set{
		values: copies,
	}
}

func (s *Set) Difference(b *Set) *Set {
	differences := make(map[any]bool)
	for k := range s.values {
		ok := b.Has(k)
		if !ok {
			differences[k] = true
		}
	}
	return &Set{
		values: differences,
	}
}

func (a *Set) IsSubset(b *Set) bool {
	for k := range a.values {
		if !b.Has(k) {
			fmt.Println(k, a.Has(k), a.List(), b.List())
			return false
		}
	}
	return true
}

func Union(sets ...*Set) *Set {
	unions := make(map[any]bool)
	for _, s := range sets {
		for k := range s.values {
			unions[k] = true
		}
	}
	return &Set{
		values: unions,
	}
}

func Intersect(sets ...*Set) *Set {
	elements := make(map[any]int)
	intersections := make(map[any]bool)
	for _, s := range sets {
		for k := range s.values {
			elements[k] += 1
		}
	}

	for k := range elements {
		if elements[k] == len(sets) {
			intersections[k] = true
		}
	}

	return &Set{
		values: intersections,
	}
}
