package problem3

type Set struct {
	array []any
}

func NewSet() *Set {
	return &Set{}
}

func (set *Set) Add(input any) {
	isDuplicate := false
	for _, val := range set.array {
		if val == input {
			isDuplicate = true
		}
	}

	if isDuplicate == false {
		set.array = append(set.array, input)
	}
}

func (set *Set) Remove(input any) {
	var newArray []any
	for _, val := range set.array {
		if val != input {
			newArray = append(newArray, val)
		}
	}
	set.array = newArray
}

func (set *Set) IsEmpty() bool {
	return len(set.array) == 0
}

func (set *Set) Size() int {
	return len(set.array)
}

func (set *Set) List() []any {
	return set.array
}

func (set *Set) Has(input any) bool {
	for _, val := range set.array {
		if val == input {
			return true
		}
	}
	return false
}

func (set *Set) Copy() *Set {
	newSet := NewSet()
	for _, val := range set.array {
		newSet.Add(val)
	}
	return newSet
}

func (set *Set) Difference(s2 *Set) *Set {
	resultSet := NewSet()
	for _, val := range set.array {
		var flag bool
		for _, v2 := range s2.array {
			if val == v2 {
				flag = true
				break
			}
		}

		if !flag {
			resultSet.Add(val)
		}
	}
	return resultSet
}

func (set *Set) IsSubset(s2 *Set) bool {
	for _, val := range set.array {
		var flag bool
		for _, v2 := range s2.array {
			if val == v2 {
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
		for _, val := range set.array {
			resultSet.Add(val)
		}
	}
	return resultSet
}

func Intersect(sets ...*Set) *Set {
	result := NewSet()
	for _, val := range sets[0].array {
		result.Add(val)
	}

	for _, set := range sets[1:] {
		temp := NewSet()
		for _, val := range result.array {
			if set.Has(val) {
				temp.Add(val)
			}
		}
		result = temp
	}
	return result
}
