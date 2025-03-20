package problem3

type Set struct {
	data []any
}

func NewSet() *Set {
	return &Set{}
}

func (s *Set) Add(new any) {

	if !s.Has(new) {
		s.data = append(s.data, new)
	}

}
func (s *Set) Remove(deleted any) {
	if s.data == nil {
		return
	}

	if len(s.data) == 1 {
		s.data = nil
		return
	}
	if s.data[len(s.data)-1] == deleted {
		s.data = s.data[:len(s.data)-1]
		return
	}
	for i, v := range s.data {

		if v == deleted {
			s.data = append(s.data[:i], s.data[i+1:]...)
		}
	}
}
func (s *Set) IsEmpty() bool {
	return s.data == nil
}
func (s *Set) Size() int {
	return len(s.data)
}
func (s *Set) Has(correspond any) bool {
	for _, v := range s.data {
		if v == correspond {
			return true
		}
	}
	return false
}
func (s *Set) Copy() *Set {
	newset := s
	return newset

}
func (s *Set) Difference(diff *Set) Set {
	res := Set{}
	for _, v := range s.data {
		if !diff.Has(v) {
			res.Add(v)
		}
	}
	return res
}
func (s *Set) IsSubset(subs *Set) bool {
	if subs.Size() == 0 {
		return true
	}

	for _, v := range s.data {
		if !subs.Has(v) {
			return false
		}
	}
	return true
}

func Union(uni ...*Set) Set {
	var newany Set

	return newany
}
func Intersect(inter ...*Set) Set {
	if len(inter) == 0 {
		return Set{}
	}

	var lastany Set
	// newany := []any{}
	// 	for _, v := range inter {
	// }

	// }
	return lastany
}
func (s *Set) List() []any {
	var newany []any
	var lastany any
	for _, v := range s.data {
		if v == lastany {
			lastany = nil
			continue
		} else {
			newany = append(newany, v)
			lastany = v

		}

	}
	return newany
}
