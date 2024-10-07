package problem2

import (
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("Push", func(t *testing.T) {
		table := []struct {
			vals []any
		}{
			{[]any{1, 2, 3}},
			{[]any{"1", "2", "3", "4"}},
			{[]any{true, false}},
		}

		for _, r := range table {
			stack := Stack{}
			for _, v := range r.vals {
				stack.Push(v)
			}
			if stack.Size() != len(r.vals) {
				t.Errorf("Push(%v) was incorrect, got len: %d, expected len: %d", r.vals, stack.Size(), len(r.vals))
			}
		}
	})

	t.Run("Pop", func(t *testing.T) {
		table := []struct {
			vals []any
		}{
			{[]any{1, 2, 3}},
			{[]any{"1", "2", "3", "4"}},
			{[]any{true, false}},
		}

		for _, r := range table {
			stack := Stack{}

			for _, v := range r.vals {
				stack.Push(v)
			}
			for i := range r.vals {
				d, _ := stack.Pop()
				if d != r.vals[len(r.vals)-1-i] {
					t.Errorf("Pop() was incorrect, got: %v, expected: %v", d, r.vals[len(r.vals)-1-i])
				}
			}
		}
	})

	t.Run("Pop_empty", func(t *testing.T) {
		stack := Stack{}

		_, err := stack.Pop()
		if err == nil {
			t.Errorf("Pop() empty was incorrect, expected non nil error")
		}
	})

	t.Run("Peek", func(t *testing.T) {
		table := []struct {
			vals []any
		}{
			{[]any{1, 2, 3}},
			{[]any{"1", "2", "3", "4"}},
			{[]any{true, false}},
		}

		for _, r := range table {
			stack := Stack{}

			for _, v := range r.vals {
				stack.Push(v)
			}
			d, _ := stack.Peek()
			if d != r.vals[len(r.vals)-1] {
				t.Errorf("Peek() was incorrect, got: %v, expected: %v", d, r.vals[len(r.vals)-1])
			}
			if stack.Size() != len(r.vals) {
				t.Errorf("Peek() has incorrect size, got: %d, expected: %d", stack.Size(), len(r.vals))
			}
		}
	})

	t.Run("Peek_empty", func(t *testing.T) {
		stack := Stack{}

		_, err := stack.Peek()
		if err == nil {
			t.Errorf("Peek() empty was incorrect, expected non nil error")
		}
	})

	t.Run("Size", func(t *testing.T) {
		table := []struct {
			vals       []any
			removeSize int
		}{
			{[]any{1, 2, 3}, 0},
			{[]any{"1", "2", "3", "4"}, 1},
			{[]any{true, false}, 2},
		}

		for _, r := range table {
			stack := Stack{}

			for _, v := range r.vals {
				stack.Push(v)
			}
			for range r.removeSize {
				_, _ = stack.Pop()
			}
			if stack.Size() != len(r.vals)-r.removeSize {
				t.Errorf("Size() was incorrect, got: %d, expected: %d", stack.Size(), len(r.vals)-r.removeSize)
			}
		}
	})

	t.Run("IsEmpty", func(t *testing.T) {
		stack := Stack{}

		if stack.IsEmpty() != true {
			t.Errorf("IsEmpty() was incorrect, got: %t, expected: %t", stack.IsEmpty(), true)
		}
		stack.Push(1)
		if stack.IsEmpty() != false {
			t.Errorf("IsEmpty() was incorrect, got: %t, expected: %t", stack.IsEmpty(), false)
		}
		stack.Push(2)
		if stack.IsEmpty() != false {
			t.Errorf("IsEmpty() was incorrect, got: %t, expected: %t", stack.IsEmpty(), false)
		}
		_, _ = stack.Pop()
		if stack.IsEmpty() != false {
			t.Errorf("IsEmpty() was incorrect, got: %t, expected: %t", stack.IsEmpty(), false)
		}
		_, _ = stack.Pop()
		if stack.IsEmpty() != true {
			t.Errorf("IsEmpty() was incorrect, got: %t, expected: %t", stack.IsEmpty(), true)
		}
	})
}
