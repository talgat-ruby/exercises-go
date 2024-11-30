package problem3

import "errors"

type Set struct {
	set []any
}

func NewSet() *Set {
	return &Set{
		set: make([]any, 0), // Инициализируем пустой срез с 0 элементов
	}
}

func (s *Set) Has(val any) bool {
	for i := 0; i < len(s.set); i++ {
		if s.set[i] == val {
			return true
		}
	}
	return false
}

func (s *Set) Add(val any) {
	if !s.Has(val) {
		s.set = append(s.set, val)
	}
}

func (s *Set) Remove(val any) error {
	for i := 0; i < len(s.set); i++ {
		if s.set[i] == val {
			s.set = append(s.set[:i], s.set[i+1:]...)
			return nil
		}
	}
	return errors.New("element not found")
}

func (s *Set) IsEmpty() bool {
	return len(s.set) == 0
}

func (s *Set) Size() int {
	return len(s.set)
}

func (s *Set) List() []any {
	return s.set
}

func (s *Set) Copy() *Set {
	cSet := NewSet()
	cSet.set = append(cSet.set, s.set...)
	return cSet
}

func (a *Set) Difference(b *Set) *Set {
	diffSet := NewSet()
	for _, val := range a.set {
		if !b.Has(val) {
			diffSet.Add(val)
		}
	}
	return diffSet
}

func (b *Set) IsSubset(a *Set) bool {
	for _, val := range b.set {
		if !a.Has(val) {
			return false
		}
	}
	return true
}

func Union(sets ...*Set) *Set {
	unionSet := NewSet()
	for _, set := range sets {
		for _, val := range set.set {
			unionSet.Add(val)
		}
	}
	return unionSet
}

func Intersect(sets ...*Set) *Set {
	if len(sets) == 0 {
		return NewSet()
	}
	if len(sets) == 1 {
		return sets[0]
	}

	mini := sets[0].Size()
	index := 0

	for i := 1; i < len(sets); i++ {
		if sets[i].Size() < mini {
			mini = sets[i].Size()
			index = i
		}
	}

	output := NewSet()
	minSetList := sets[index].List()

	for i := 0; i < len(minSetList); i++ {
		isInside := true
		for j := 0; j < len(sets); j++ {
			if j == index {
				continue
			}
			if !sets[j].Has(minSetList[i]) {
				isInside = false
				break
			}
		}
		if isInside {
			output.Add(minSetList[i])
		}
	}
	return output
}
