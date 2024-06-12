package problem3

type Set struct {
	elements []any
}

func NewSet() *Set {
	return &Set{}
}

func (s *Set) Add(newElement any) {
	isExist := false
	for i := 0; i < len(s.elements); i++ {
		if s.elements[i] == newElement {
			isExist = true
			break
		}
	}
	if !isExist {
		s.elements = append(s.elements, newElement)
	}
}

func (s *Set) Remove(element any) {
	var newSetElements []any
	for i := 0; i < len(s.elements); i++ {
		if s.elements[i] != element {
			newSetElements = append(newSetElements, s.elements[i])
		}
	}
	s.elements = newSetElements
}

func (s *Set) IsEmpty() any {
	return len(s.elements) == 0
}

func (s *Set) Size() any {
	return len(s.elements)
}

func (s *Set) List() []any {
	return s.elements
}

func (s *Set) Has(element any) bool {
	for i := 0; i < len(s.elements); i++ {
		if s.elements[i] == element {
			return true
		}
	}
	return false
}

func (s *Set) Copy() *Set {
	return &Set{s.elements}
}

func (s *Set) Difference(B *Set) Set {
	var differenceSet Set
	for i := 0; i < len(s.elements); i++ {
		if !B.Has(s.elements[i]) {
			differenceSet.elements = append(differenceSet.elements, s.elements[i])
		}
	}
	return differenceSet
}

func (s *Set) IsSubset(B *Set) bool {
	for i := 0; i < len(s.elements); i++ {
		if !B.Has(s.elements[i]) {
			return false
		}
	}
	return true
}

func Union(B ...*Set) Set {
	var unionSet Set
	for i := 0; i < len(B); i++ {
		for j := 0; j < len(B[i].elements); j++ {
			unionSet.Add(B[i].elements[j])
		}
	}
	return unionSet
}

func Intersect(B ...*Set) Set {
	var intersectSet Set
	unionSet := Union(B...)
	for i := 0; i < len(unionSet.elements); i++ {
		isExistInEverySet := true
		for j := 0; j < len(B); j++ {
			if !B[j].Has(unionSet.elements[i]) {
				isExistInEverySet = false
			}
		}
		if isExistInEverySet {
			intersectSet.Add(unionSet.elements[i])
		}
	}
	return intersectSet
}
