package problem1

import (
	"testing"
)

func TestQueue(t *testing.T) {
	t.Run("Enqueue", func(t *testing.T) {
		table := []struct {
			vals []any
		}{
			{[]any{1, 2, 3}},
			{[]any{"1", "2", "3", "4"}},
			{[]any{true, false}},
		}

		for _, r := range table {
			queue := Queue{}
			for _, v := range r.vals {
				queue.Enqueue(v)
			}
			if queue.Size() != len(r.vals) {
				t.Errorf("Enqueue(%v) was incorrect, got len: %d, expected len: %d", r.vals, queue.Size(), len(r.vals))
			}
		}
	})

	t.Run("Dequeue", func(t *testing.T) {
		table := []struct {
			vals []any
		}{
			{[]any{1, 2, 3}},
			{[]any{"1", "2", "3", "4"}},
			{[]any{true, false}},
		}

		for _, r := range table {
			queue := Queue{}

			for _, v := range r.vals {
				queue.Enqueue(v)
			}
			for i := range r.vals {
				d, _ := queue.Dequeue()
				if d != r.vals[i] {
					t.Errorf("Dequeue() was incorrect, got: %v, expected: %v", d, r.vals[i])
				}
			}
		}
	})

	t.Run("Dequeue_empty", func(t *testing.T) {
		queue := Queue{}

		_, err := queue.Dequeue()
		if err == nil {
			t.Errorf("Dequeue() empty was incorrect, expected non nil error")
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
			queue := Queue{}

			for _, v := range r.vals {
				queue.Enqueue(v)
			}
			d, _ := queue.Peek()
			if d != r.vals[0] {
				t.Errorf("Peek() was incorrect, got: %v, expected: %v", d, r.vals[0])
			}
			if queue.Size() != len(r.vals) {
				t.Errorf("Peek() has incorrect size, got: %d, expected: %d", queue.Size(), len(r.vals))
			}
		}
	})

	t.Run("Peek_empty", func(t *testing.T) {
		queue := Queue{}

		_, err := queue.Peek()
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
			queue := Queue{}

			for _, v := range r.vals {
				queue.Enqueue(v)
			}
			for range r.removeSize {
				_, _ = queue.Dequeue()
			}
			if queue.Size() != len(r.vals)-r.removeSize {
				t.Errorf("Size() was incorrect, got: %d, expected: %d", queue.Size(), len(r.vals)-r.removeSize)
			}
		}
	})

	t.Run("IsEmpty", func(t *testing.T) {
		queue := Queue{}

		if queue.IsEmpty() != true {
			t.Errorf("IsEmpty() was incorrect, got: %t, expected: %t", queue.IsEmpty(), true)
		}
		queue.Enqueue(1)
		if queue.IsEmpty() != false {
			t.Errorf("IsEmpty() was incorrect, got: %t, expected: %t", queue.IsEmpty(), false)
		}
		queue.Enqueue(2)
		if queue.IsEmpty() != false {
			t.Errorf("IsEmpty() was incorrect, got: %t, expected: %t", queue.IsEmpty(), false)
		}
		_, _ = queue.Dequeue()
		if queue.IsEmpty() != false {
			t.Errorf("IsEmpty() was incorrect, got: %t, expected: %t", queue.IsEmpty(), false)
		}
		_, _ = queue.Dequeue()
		if queue.IsEmpty() != true {
			t.Errorf("IsEmpty() was incorrect, got: %t, expected: %t", queue.IsEmpty(), true)
		}
	})
}
