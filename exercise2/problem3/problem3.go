package problem3

type Set struct {
	m map[any]struct{}
}

func NewSet() *Set {
	return &Set{m: make(map[any]struct{})}
}

func (s *Set) Add(v any) {
	if _, ok := s.m[v]; ok {
		return
	}

	s.m[v] = struct{}{}
}

func (s *Set) Remove(v any) {
	delete(s.m, v)
}

func (s *Set) IsEmpty() bool {
	return len(s.m) == 0
}

func (s *Set) Size() int {
	return len(s.m)
}

func (s *Set) List() []any {
	list := make([]any, len(s.m))
	index := 0
	for k := range s.m {
		list[index] = k
		index++
	}
	return list
}

func (s *Set) Has(v any) bool {
	_, ok := s.m[v]
	return ok
}

func (s *Set) Copy() *Set {
	set2 := NewSet()
	for k := range s.m {
		set2.Add(k)
	}
	return set2
}

func (s *Set) Difference(s2 *Set) Set {
	diff := NewSet()
	list := s.List()
	for _, v := range list {
		if !s2.Has(v) {
			diff.Add(v)
		}
	}
	return *diff
}

func (s *Set) IsSubset(s2 *Set) bool {
	sList := s.List()
	for _, v := range sList {
		if !s2.Has(v) {
			return false
		}
	}
	return true
}

func Union(sets ...*Set) Set {
	unionSet := NewSet()
	for _, v := range sets {
		list := v.List()
		for _, elem := range list {
			unionSet.Add(elem)
		}
	}
	return *unionSet
}

func Intersect(sets ...*Set) Set {
	if len(sets) == 0 {
		return Set{}
	}

	intersectSet := NewSet()
	list := sets[0].List()
	for _, v := range list {
		contains := true
		for i := 1; i < len(sets); i++ {
			if sets[i].Has(v) {
				continue
			}
			contains = false
		}

		if contains {
			intersectSet.Add(v)
		}
	}

	return *intersectSet
}
