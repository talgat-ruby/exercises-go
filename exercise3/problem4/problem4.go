package problem4

import "fmt"

// Set represents a collection of unique elements.
type Set struct {
	elements map[interface{}]struct{}
}

// NewSet creates a new empty Set.
func NewSet() *Set {
	return &Set{elements: make(map[interface{}]struct{})}
}

// Add adds a new element to the Set.
func (s *Set) Add(value interface{}) {
	s.elements[value] = struct{}{}
}

// Remove removes an existing element from the Set.
func (s *Set) Remove(value interface{}) {
	delete(s.elements, value)
}

// IsEmpty checks if the Set is empty.
func (s *Set) IsEmpty() bool {
	return len(s.elements) == 0
}

// Size returns the number of elements in the Set.
func (s *Set) Size() int {
	return len(s.elements)
}

// List returns a slice of elements in the Set.
func (s *Set) List() []interface{} {
	list := make([]interface{}, 0, len(s.elements))
	for key := range s.elements {
		list = append(list, key)
	}
	return list
}

// Has checks if an element is a member of the Set.
func (s *Set) Has(value interface{}) bool {
	_, exists := s.elements[value]
	return exists
}

// Copy creates a copy of the Set.
func (s *Set) Copy() *Set {
	copySet := NewSet()
	for key := range s.elements {
		copySet.Add(key)
	}
	return copySet
}

// Difference returns the difference between two Sets (A - B).
func (s *Set) Difference(other *Set) *Set {
	diffSet := NewSet()
	for key := range s.elements {
		if !other.Has(key) {
			diffSet.Add(key)
		}
	}
	return diffSet
}

// IsSubset checks if Set B is a subset of Set A.
func (s *Set) IsSubset(other *Set) bool {
	for key := range other.elements {
		if !s.Has(key) {
			return false
		}
	}
	return true
}

// Union returns the union of multiple Sets.
func (s *Set) Union(sets ...*Set) *Set {
	unionSet := s.Copy() // Start with a copy of the current set
	for _, set := range sets {
		for key := range set.elements {
			unionSet.Add(key)
		}
	}
	return unionSet
}

// Intersect returns the intersection of multiple Sets.
func (s *Set) Intersect(sets ...*Set) *Set {
	intersectSet := NewSet()
	for key := range s.elements {
		inAll := true
		for _, set := range sets {
			if !set.Has(key) {
				inAll = false
				break
			}
		}
		if inAll {
			intersectSet.Add(key)
		}
	}
	return intersectSet
}

// Example usage
func main() {
	setA := NewSet()
	setA.Add(1)
	setA.Add(2)
	setA.Add(3)

	setB := NewSet()
	setB.Add(2)
	setB.Add(3)
	setB.Add(4)

	fmt.Println("Set A:", setA.List())                             // Set A: [1 2 3]
	fmt.Println("Set B:", setB.List())                             // Set B: [2 3 4]
	fmt.Println("Difference A - B:", setA.Difference(setB).List()) // Difference A - B: [1]
	fmt.Println("Is B a subset of A?", setA.IsSubset(setB))        // Is B a subset of A? false

	unionSet := setA.Union(setB)
	fmt.Println("Union of A and B:", unionSet.List()) // Union of A and B: [1 2 3 4]

	intersectSet := setA.Intersect(setB)
	fmt.Println("Intersection of A and B:", intersectSet.List()) // Intersection of A and B: [2 3]
}
