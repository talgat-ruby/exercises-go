package problem3

type Set struct {
	m    map[any]struct{}
	size int
}

func NewSet() *Set {
	return &Set{make(map[any]struct{}), 0}
}

func (set *Set) Add(x any) {
	if !set.Has(x) {
		set.m[x] = struct{}{}
		set.size++
	}
}

func (set *Set) Size() int {
	return set.size
}

func (set *Set) Has(x any) bool {
	if _, ok := set.m[x]; ok {
		return true
	}
	return false
}

func (set *Set) IsEmpty() bool {
	return set.size == 0
}

func (set *Set) Remove(x any) {
	if set.Has(x) {
		delete(set.m, x)
		set.size--
	}
}

func (set *Set) List() []any {
	out := make([]any, set.size)
	var j int
	for i := range set.m {
		out[j] = i
		j++
	}
	return out
}

func (set *Set) Copy() *Set {
	out := make(map[any]struct{})
	for x := range set.m {
		out[x] = struct{}{}
	}
	return &Set{m: out, size: set.size}
}

func Union(sets ...*Set) *Set {
	m := map[any]struct{}{}
	for _, set := range sets {
		for x := range set.m {
			m[x] = struct{}{}
		}
	}
	return &Set{m: m, size: len(m)}
}

func Intersect(sets ...*Set) *Set {
	out := Union(sets...)
	for x := range out.m {
		for _, y := range sets {
			if !y.Has(x) {
				out.Remove(x)
			}
		}
	}
	return out
}

func (set *Set) Difference(other *Set) *Set {
	m := map[any]struct{}{}
	for x := range set.m {
		if !other.Has(x) {
			m[x] = struct{}{}
		}
	}
	for x := range other.m {
		if !set.Has(x) {
			m[x] = struct{}{}
		}
	}
	return &Set{m: m, size: len(m)}
}

func (set *Set) IsSubset(other *Set) bool {
	for x := range set.m {
		if !other.Has(x) {
			return false
		}
	}
	return true
}
