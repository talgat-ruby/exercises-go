package problem3

type Set struct {
    elements map[interface{}]struct{}
}

func NewSet() *Set {
    return &Set{
        elements: make(map[interface{}]struct{}),
    }
}


func (s *Set) Add(value interface{}) {
    s.elements[value] = struct{}{}
}


func (s *Set) Remove(value interface{}) {
    delete(s.elements, value)
}


func (s *Set) IsEmpty() bool {
    return len(s.elements) == 0
}


func (s *Set) Size() int {
    return len(s.elements)
}


func (s *Set) List() []interface{} {
    keys := make([]interface{}, 0, len(s.elements))
    for key := range s.elements {
        keys = append(keys, key)
    }
    return keys
}


func (s *Set) Has(value interface{}) bool {
    _, exists := s.elements[value]
    return exists
}


func (s *Set) Copy() *Set {
    newSet := NewSet()
    for key := range s.elements {
        newSet.Add(key)
    }
    return newSet
}


func (s *Set) Difference(other *Set) *Set {
    diffSet := NewSet()
    for key := range s.elements {
        if !other.Has(key) {
            diffSet.Add(key)
        }
    }
    return diffSet
}


func (s *Set) IsSubset(other *Set) bool {
    for key := range s.elements {
        if !other.Has(key) {
            return false
        }
    }
    return true
}


func Union(sets ...*Set) *Set {
    unionSet := NewSet()
    for _, set := range sets {
        for key := range set.elements {
            unionSet.Add(key)
        }
    }
    return unionSet
}


func Intersect(sets ...*Set) *Set {
    if len(sets) == 0 {
        return NewSet()
    }

    intersectionSet := NewSet()
    firstSet := sets[0]

    for key := range firstSet.elements {
        isInAll := true
        for _, set := range sets[1:] {
            if !set.Has(key) {
                isInAll = false
                break
            }
        }
        if isInAll {
            intersectionSet.Add(key)
        }
    }
    return intersectionSet
}
