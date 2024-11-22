package problem2

type Stack struct {
	elements []interface{}
}

func (s *Stack) Push(v interface{}) {
	s.elements = append(s.elements, v)
}
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	topIndex := len(s.elements) - 1
	element := s.elements[topIndex]
	s.elements = s.elements[:topIndex]
	return element
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}
func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.elements[len(s.elements)-1]
}
func (s *Stack) Size() int {
	return len(s.elements)
}
