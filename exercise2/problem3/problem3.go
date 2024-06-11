package problem3

type Set struct {
	items []any
	size  int
}

func NewSet() *Set {
	return &Set{}
}

func (s *Set) Add(element any) {
	for _, item := range s.items {
		if element == item {
			return
		}
	}
	s.items = append(s.items, element)
	s.size++
}

func (s *Set) Remove(element any) {
	for i, v := range s.items {
		if v == element {
			s.items = append(s.items[:i], s.items[i+1:]...)
			s.size--
		}
	}
}

func (s *Set) IsEmpty() bool {
	return s.size == 0
}

func (s *Set) Size() int {
	return s.size
}

func (s *Set) List() []any {
	return s.items
}

func (s *Set) Has(element any) bool {
	for _, v := range s.items {
		if v == element {
			return true
		}
	}
	return false
}

func (s *Set) Copy() *Set {
	return &Set{
		items: s.items,
		size:  s.size,
	}
}

// optional

func (s *Set) Difference(set2 *Set) *Set {
	return nil
}

func (s *Set) IsSubset(set2 *Set) bool {
	return false
}

func Union(...*Set) *Set {
	return &Set{}
}

func Intersect(...*Set) *Set {
	return &Set{}
}
