package problem3

import "fmt"

// Set structure
type Set struct {
	elements map[interface{}]struct{}
}

// NewSet creates and returns a new Set
func NewSet() *Set {
	return &Set{
		elements: make(map[interface{}]struct{}),
	}
}

// Add adds a new element to the Set
func (s *Set) Add(element interface{}) {
	s.elements[element] = struct{}{}
}

// Remove removes an existing element from the Set
func (s *Set) Remove(element interface{}) {
	delete(s.elements, element)
}

// IsEmpty checks if the Set is empty
func (s *Set) IsEmpty() bool {
	return len(s.elements) == 0
}

// Size gets the size of the Set
func (s *Set) Size() int {
	return len(s.elements)
}

// List gets the slice out of the Set
func (s *Set) List() []interface{} {
	result := make([]interface{}, 0, len(s.elements))
	for element := range s.elements {
		result = append(result, element)
	}
	return result
}

// Has checks if an element is a member of the Set
func (s *Set) Has(element interface{}) bool {
	_, exists := s.elements[element]
	return exists
}

// Copy creates a copy of the Set
func (s *Set) Copy() *Set {
	newSet := NewSet()
	for element := range s.elements {
		newSet.Add(element)
	}
	return newSet
}

// Difference returns the difference of two sets A and B
func (s *Set) Difference(other *Set) *Set {
	diff := NewSet()
	for element := range s.elements {
		if !other.Has(element) {
			diff.Add(element)
		}
	}
	return diff
}

// IsSubset checks if the current set is a subset of another set
func (s *Set) IsSubset(other *Set) bool {
	// Check if all elements in the current set are in the other set
	for element := range s.elements {
		if !other.Has(element) {
			return false // If any element is not found, return false
		}
	}
	return true // All elements are found, it's a subset
}

// Union returns the union of two or more sets
func Union(sets ...*Set) *Set {
	unionSet := NewSet()
	for _, set := range sets {
		for element := range set.elements {
			unionSet.Add(element)
		}
	}
	return unionSet
}

// Intersect returns the intersection of two or more sets
func Intersect(sets ...*Set) *Set {
	if len(sets) == 0 {
		return NewSet()
	}
	intersectSet := NewSet()
	for element := range sets[0].elements {
		isCommon := true
		for _, set := range sets[1:] {
			if !set.Has(element) {
				isCommon = false
				break
			}
		}
		if isCommon {
			intersectSet.Add(element)
		}
	}
	return intersectSet
}

func main() {
	// Example usage
	setA := NewSet()
	setA.Add(1)
	setA.Add(2)
	setA.Add(3)

	setB := NewSet()
	setB.Add(1)
	setB.Add(2)

	fmt.Println("Set A:", setA.List()) // should show {1, 2, 3}
	fmt.Println("Set B:", setB.List()) // should show {1, 2}

	// Checking subsets
	fmt.Println("B is a subset of A:", setB.IsSubset(setA)) // should be true
	fmt.Println("A is a subset of B:", setA.IsSubset(setB)) // should be false
}
