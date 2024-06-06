package problem3

type Set struct {
	Map map[any]int
}

func NewSet() Set {
	m := make(map[any]int)
	return Set{m}
}
func (s *Set) Add(elem any) {
	s.Map[elem]++
}
func (s *Set) Size() int {
	return len(s.Map)
}
func (s *Set) Remove(elem any) {
	delete(s.Map, elem)
}
func (s *Set) IsEmpty() bool {
	if len(s.Map) == 0 {
		return true
	}
	return false
}
func (s *Set) List() []any {
	var res []any
	for k := range s.Map {
		res = append(res, k)
	}
	return res
}
func (s *Set) Has(elem any) bool {
	if _, ok := s.Map[elem]; ok {
		return true
	}
	return false
}
func (s *Set) Copy() Set {
	newM := make(map[any]int)
	for k, v := range s.Map {
		newM[k] = v
	}
	return Set{newM}
}
func (s *Set) Difference(set2 Set) Set {
	res := make(map[any]int)
	for k, v := range s.Map {
		var flag bool
		for k2 := range set2.Map {
			if k == k2 {
				flag = true
				break
			}
		}
		if !flag {
			res[k] = v
		}
	}
	return Set{res}
}
func (s *Set) IsSubset(set Set) bool {
	for k := range s.Map {
		if _, ok := set.Map[k]; !ok {
			return false
		}
	}
	return true
}
func Union(arr ...Set) Set {
	m := make(map[any]int)
	for i := 0; i < len(arr); i++ {
		for k, v := range arr[i].Map {
			m[k] = v
		}
	}
	return Set{m}
}
func Intersect(arr ...Set) Set {
	m := make(map[any]int)
	for i := 0; i < len(arr); i++ {
		for k := range arr[i].Map {
			m[k]++
		}
	}
	for k, v := range m {
		if v != len(arr) {
			delete(m, k)
		}
	}
	return Set{m}
}
