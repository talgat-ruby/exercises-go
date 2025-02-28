package problem3

type Set struct {
	slc []interface{}
}

func NewSet() *Set {
	return &Set{}
}

func (s *Set) Add(value interface{}) {
	flag := false
	for _, v := range s.slc {
		if v == value {
			flag = true
		}
	}

	if !flag {
		s.slc = append(s.slc, value)
	}
}

func (s *Set) Remove(value interface{}) {
	var newSlc []interface{}
	for _, v := range s.slc {
		if v != value {
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

func (s *Set) List() []interface{} {
	return s.slc
}

func (s *Set) Has(value interface{}) bool {
	for _, v := range s.slc {
		if v == value {
			return true
		}
	}
	return false
}

func (s *Set) Copy() *Set {
	newSet := NewSet()
	for _, value := range s.slc {
		newSet.Add(value)
	}
	return newSet
}

func (s *Set) Difference(ss *Set) *Set {
	newSet := NewSet()
	for _, value := range s.slc {
		flag := false
		for _, value2 := range ss.slc {
			if value == value2 {
				flag = true
				break
			}
		}
		if !flag {
			newSet.Add(value)
		}
	}
	return newSet
}

func (s *Set) IsSubset(ss *Set) bool {
	for _, value := range s.slc {
		flag := false
		for _, value2 := range ss.slc {
			if value == value2 {
				flag = true
				break
			} else {
				flag = false
			}
		}
		if !flag {
			return false
		}
	}
	return true
}

func Union(sets ...*Set) *Set {
	newSet := NewSet()
	for _, set := range sets {
		for _, value := range set.slc {
			newSet.Add(value)
		}
	}
	return newSet
}

func Intersect(sets ...*Set) *Set {
	newSet := NewSet()
	for _, value := range sets[0].slc {
		newSet.Add(value)
	}

	for _, set := range sets[1:] {
		temp := NewSet()
		for _, value := range newSet.slc {
			if set.Has(value) {
				temp.Add(value)
			}
		}
		newSet = temp
	}
	return newSet
}
