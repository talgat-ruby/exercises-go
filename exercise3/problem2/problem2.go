package problem2

import "fmt"

// Stack представляет структуру стека
type Stack struct {
	elements []interface{} // Срез для хранения элементов стека
}

// Push добавляет элемент на вершину стека
func (s *Stack) Push(value interface{}) {
	s.elements = append(s.elements, value)
}

// Pop удаляет и возвращает элемент с вершины стека
func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false // Возвращаем nil и false, если стек пуст
	}
	// Получаем последний элемент
	index := len(s.elements) - 1
	element := s.elements[index]
	// Удаляем элемент с конца среза
	s.elements = s.elements[:index]
	return element, true
}

// Peek возвращает элемент с вершины стека без его удаления
func (s *Stack) Peek() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false // Возвращаем nil и false, если стек пуст
	}
	return s.elements[len(s.elements)-1], true
}

// Size возвращает количество элементов в стеке
func (s *Stack) Size() int {
	return len(s.elements)
}

// IsEmpty проверяет, пуст ли стек
func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

// Пример использования
func main() {
	stack := &Stack{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println("Size:", stack.Size())           // Size: 3
	fmt.Println("Peek:", stack.Peek())           // Peek: 3
	fmt.Println("Pop:", stack.Pop())             // Pop: 3
	fmt.Println("Size after pop:", stack.Size()) // Size after pop: 2
	fmt.Println("Is Empty:", stack.IsEmpty())    // Is Empty: false

	stack.Pop()                                                 // Удаляет 2
	stack.Pop()                                                 // Удаляет 1
	fmt.Println("Is Empty after popping all:", stack.IsEmpty()) // Is Empty after popping all: true
}
