package problem3

type Set struct {
	elements []any
}

func NewSet() *Set {
	return &Set{elements: make([]any, 0)}
}

func (q *Set) Add(element any) {
	if !q.Has(element) {
		q.elements = append(q.elements, element)
	}
}

func (q *Set) Remove(e any) {
	i := q.indexOf(e)
	if i >= 0 {
		q.RemoveByIndex(i)
	}
}

func (q *Set) RemoveByIndex(i int) {
	q.elements = append(q.elements[:i], q.elements[i+1:]...)
}

func (q *Set) Has(e any) bool {
	return q.indexOf(e) >= 0
}

func (q *Set) indexOf(e any) int {
	for i, x := range q.elements {
		if e == x {
			return i
		}
	}

	return -1
}

func (q *Set) Size() int {
	return len(q.elements)
}

func (q *Set) IsEmpty() bool {
	return q.elements == nil || len(q.elements) == 0
}

func (q *Set) List() []any {
	return q.elements
}

func (q *Set) Copy() *Set {
	c := NewSet()
	c.elements = append(c.elements, q.elements...)

	return c
}

func (q *Set) Difference(o *Set) *Set {
	diff := q.Copy()
	for i := len(diff.elements) - 1; i >= 0; i-- {
		x := diff.elements[i]
		if o.Has(x) {
			diff.RemoveByIndex(i)
		}
	}

	return diff
}

func (q *Set) IsSubset(o *Set) bool {
	for _, x := range q.elements {
		if !o.Has(x) {
			return false
		}
	}

	return true
}

func Union(arr ...*Set) *Set {
	union := NewSet()
	for _, s := range arr {
		for _, e := range s.elements {
			if !union.Has(e) {
				union.Add(e)
			}
		}
	}

	return union
}

func Intersect(arr ...*Set) *Set {
	intersection := NewSet()
	lastIndex := -1
	for _, e := range arr[0].elements {
		lastIndex++
		intersection.Add(e)
		for i := 1; i < len(arr); i++ {
			s := arr[i]
			if !s.Has(e) {
				intersection.RemoveByIndex(lastIndex)
				lastIndex--
				break
			}
		}
	}

	return intersection
}
