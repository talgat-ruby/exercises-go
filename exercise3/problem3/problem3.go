package problem3

type Set struct {
	slc []any
}

func NewSet() *Set {
	return &Set{}
}

func (s *Set) Add(x any) {
	var flag bool
	for _, v := range s.slc {
		if v == x {
			flag = true
		}
	}

	if !flag {
		s.slc = append(s.slc, x)
	}
}

func (s *Set) Remove(x any) {
	var newSlc []any
	for _, v := range s.slc {
		if v != x {
			newSlc = append(newSlc, v)
		}
	}
	s.slc = newSlc
}

func (s *Set) IsEmpty() bool {
	return len(s.slc) == 0
}

func (s *Set) Size() int {
	return len(s.slc)
}

func (s *Set) List() []any {
	return s.slc
}

func (s *Set) Has(x any) bool {
	for _, v := range s.slc {
		if v == x {
			return true
		}
	}
	return false
}

func (s *Set) Copy() *Set {
	newSet := NewSet()
	for _, v := range s.slc {
		newSet.Add(v)
	}
	return newSet
}

func (s *Set) Difference(s2 *Set) *Set {
	resultSet := NewSet()
	for _, v := range s.slc {
		var flag bool
		for _, v2 := range s2.slc {
			if v == v2 {
				flag = true
				break
			}
		}

		if !flag {
			resultSet.Add(v)
		}
	}
	return resultSet
}

func (s *Set) IsSubset(s2 *Set) bool {
	for _, v := range s.slc {
		var flag bool
		for _, v2 := range s2.slc {
			if v == v2 {
				flag = true
				break
			}
		}

		if !flag {
			return false
		}
	}
	return true
}

func Union(sets ...*Set) *Set {
	resultSet := NewSet()
	for _, set := range sets {
		for _, v := range set.slc {
			resultSet.Add(v)
		}
	}
	return resultSet
}

func Intersect(sets ...*Set) *Set {
	result := NewSet()
	for _, v := range sets[0].slc {
		result.Add(v)
	}

	for _, set := range sets[1:] {
		temp := NewSet()
		for _, v := range result.slc {
			if set.Has(v) {
				temp.Add(v)
			}
		}
		result = temp
	}
	return result
}
