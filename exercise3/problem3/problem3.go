package problem3

type Set struct {
	items []any
}

func NewSet() *Set {
	return &Set{items: make([]any, 0)}
}

func (s *Set) Add(a any) {
	if !s.Has(a) {
		s.items = append(s.items, a)
	}
}

func (s *Set) Remove(a any) {
	if s.IsEmpty() {
		return
	}

	var num int
	for i, ch := range s.items {
		if ch == a {
			num = i
		}
	}

	s.items[num] = s.items[s.Size()-1]
	s.items = s.items[:s.Size()-1]
}

func (s *Set) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Set) Size() int {
	return len(s.items)
}

func (s *Set) List() []any {
	return s.items
}

func (s *Set) Has(item any) bool {
	for _, v := range s.items {
		if v == item {
			return true
		}
	}

	return false
}

func (s *Set) Copy() *Set {
	cp := make([]any, s.Size())
	copy(s.items, cp)
	return &Set{items: cp}
}

func (s1 *Set) Difference(s2 *Set) *Set {
	diff := make([]any, 0)
	for _, v1 := range s1.items {
		if !s2.Has(v1) {
			diff = append(diff, v1)
		}
	}

	return &Set{items: diff}
}

func (sub *Set) IsSubset(s *Set) bool {
	for _, v := range sub.items {
		if !s.Has(v) {
			return false
		}
	}

	return true
}

func Union(sets ...*Set) *Set {
	un := NewSet()
	for _, v := range sets {
		for _, v1 := range v.items {
			un.Add(v1)
		}
	}

	return un
}

func Intersect(sets ...*Set) *Set {
	in := NewSet()
	if len(sets) == 0 {
		return in
	}

	if len(sets) == 1 {
		return sets[0]
	}

	for _, v := range sets[0].items {
		in.Add(v)
		for _, v1 := range sets[1:] {
			if !v1.Has(v) {
				in.Remove(v)
			}
		}
	}

	return in
}
