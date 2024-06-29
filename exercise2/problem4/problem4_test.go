package problem4

import (
	"slices"
	"testing"

	"github.com/talgat-ruby/exercises-go/pkg/util"
)

func TestLinkedList(t *testing.T) {
	util.SkipTestOptional(t)

	t.Run("Add", func(t *testing.T) {
		t.Run("int", func(t *testing.T) {
			row := []int{1, 2, 3}
			ll := &LinkedList[int]{}

			for _, v := range row {
				el := &Element[int]{
					value: v,
				}
				ll.Add(el)
			}

			if ll.Size() != len(row) {
				t.Errorf("Add(%v) was incorrect, got len: %d, expected len: %d", row, ll.Size(), len(row))
			}
		})

		t.Run("string", func(t *testing.T) {
			row := []string{"1", "2", "3", "4"}
			ll := &LinkedList[string]{}

			for _, v := range row {
				el := &Element[string]{
					value: v,
				}
				ll.Add(el)
			}

			if ll.Size() != len(row) {
				t.Errorf("Add(%v) was incorrect, got len: %d, expected len: %d", row, ll.Size(), len(row))
			}
		})

		t.Run("bool", func(t *testing.T) {
			row := []bool{true, false}
			ll := &LinkedList[bool]{}

			for _, v := range row {
				el := &Element[bool]{
					value: v,
				}
				ll.Add(el)
			}

			if ll.Size() != len(row) {
				t.Errorf("Add(%v) was incorrect, got len: %d, expected len: %d", row, ll.Size(), len(row))
			}
		})
	})

	t.Run("Insert", func(t *testing.T) {
		row := []int{1, 2, 3}
		ll := &LinkedList[int]{}

		for _, v := range row {
			el := &Element[int]{
				value: v,
			}
			ll.Add(el)
		}

		el0 := &Element[int]{
			value: 0,
		}
		_ = ll.Insert(el0, 2)

		if ll.Size() != len(row)+1 {
			t.Errorf("Insert(%v) was incorrect, got len: %d, expected len: %d", row, ll.Size(), len(row)+1)
		}
	})

	t.Run("Insert_err", func(t *testing.T) {
		row := []int{1, 2, 3}
		ll := &LinkedList[int]{}

		for _, v := range row {
			el := &Element[int]{
				value: v,
			}
			ll.Add(el)
		}

		el0 := &Element[int]{
			value: 0,
		}
		err := ll.Insert(el0, 5)

		if err == nil {
			t.Errorf("Insert(%v) empty was incorrect, expected non nil error", row)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		t.Run("int", func(t *testing.T) {
			row := []int{1, 2, 3}
			ll := &LinkedList[int]{}

			for _, v := range row {
				el := &Element[int]{
					value: v,
				}
				ll.Add(el)
			}

			d := &Element[int]{
				value: row[1],
			}
			_ = ll.Delete(d)

			if ll.Size() != len(row)-1 {
				t.Errorf("Delete(%v) was incorrect, got len: %d, expected len: %d", row, ll.Size(), len(row)-1)
			}
		})

		t.Run("string", func(t *testing.T) {
			row := []string{"1", "2", "3", "4"}
			ll := &LinkedList[string]{}

			for _, v := range row {
				el := &Element[string]{
					value: v,
				}
				ll.Add(el)
			}

			d := &Element[string]{
				value: row[2],
			}
			_ = ll.Delete(d)

			if ll.Size() != len(row)-1 {
				t.Errorf("Delete(%v) was incorrect, got len: %d, expected len: %d", row, ll.Size(), len(row)-1)
			}
		})

		t.Run("bool", func(t *testing.T) {
			row := []bool{true, false}
			ll := &LinkedList[bool]{}

			for _, v := range row {
				el := &Element[bool]{
					value: v,
				}
				ll.Add(el)
			}
			d := &Element[bool]{
				value: row[0],
			}
			_ = ll.Delete(d)

			if ll.Size() != len(row)-1 {
				t.Errorf("Delete(%v) was incorrect, got len: %d, expected len: %d", row, ll.Size(), len(row)-1)
			}
		})
	})

	t.Run("Delete_err", func(t *testing.T) {
		row := []int{1, 2, 3}
		ll := &LinkedList[int]{}

		for _, v := range row {
			el := &Element[int]{
				value: v,
			}
			ll.Add(el)
		}

		el0 := &Element[int]{
			value: 0,
		}
		err := ll.Insert(el0, 5)

		if err == nil {
			t.Errorf("Delete(%v) empty was incorrect, expected non nil error", row)
		}
	})

	t.Run("Find", func(t *testing.T) {
		t.Run("int", func(t *testing.T) {
			row := []int{1, 2, 3}
			ll := &LinkedList[int]{}

			for _, v := range row {
				el := &Element[int]{
					value: v,
				}
				ll.Add(el)
			}

			f, _ := ll.Find(row[1])

			if f.value != row[1] || f.next.value != row[2] {
				t.Errorf("Find(%v) was incorrect, got: %v", row[1], f)
			}
		})

		t.Run("string", func(t *testing.T) {
			row := []string{"1", "2", "3", "4"}
			ll := &LinkedList[string]{}

			for _, v := range row {
				el := &Element[string]{
					value: v,
				}
				ll.Add(el)
			}

			f, _ := ll.Find(row[2])

			if f.value != row[2] || f.next.value != row[3] {
				t.Errorf("Find(%v) was incorrect, got: %v", row[2], f)
			}
		})

		t.Run("bool", func(t *testing.T) {
			row := []bool{true, false}
			ll := &LinkedList[bool]{}

			for _, v := range row {
				el := &Element[bool]{
					value: v,
				}
				ll.Add(el)
			}

			f, _ := ll.Find(row[0])

			if f.value != row[0] || f.next.value != row[1] {
				t.Errorf("Find(%v) was incorrect, got: %v", row[1], f)
			}
		})
	})

	t.Run("Find_err", func(t *testing.T) {
		row := []int{1, 2, 3}
		ll := &LinkedList[int]{}

		for _, v := range row {
			el := &Element[int]{
				value: v,
			}
			ll.Add(el)
		}

		el0 := &Element[int]{
			value: 0,
		}
		err := ll.Insert(el0, 5)

		if err == nil {
			t.Errorf("Insert(%v) empty was incorrect, expected non nil error", row)
		}
	})

	t.Run("List", func(t *testing.T) {
		t.Run("int", func(t *testing.T) {
			row := []int{1, 2, 3}
			ll := &LinkedList[int]{}

			for _, v := range row {
				el := &Element[int]{
					value: v,
				}
				ll.Add(el)
			}

			l := ll.List()
			for _, v := range l {
				if !slices.Contains(l, v) {
					t.Errorf("List() does not contain item: %v", v)
				}
			}
			if len(l) != ll.Size() {
				t.Errorf("List() was incorrect, got size: %d, expected size: %d", len(l), ll.Size())
			}
		})

		t.Run("string", func(t *testing.T) {
			row := []string{"1", "2", "3", "4"}
			ll := &LinkedList[string]{}

			for _, v := range row {
				el := &Element[string]{
					value: v,
				}
				ll.Add(el)
			}

			l := ll.List()
			for _, v := range l {
				if !slices.Contains(l, v) {
					t.Errorf("List() does not contain item: %v", v)
				}
			}
			if len(l) != ll.Size() {
				t.Errorf("List() was incorrect, got size: %d, expected size: %d", len(l), ll.Size())
			}
		})

		t.Run("bool", func(t *testing.T) {
			row := []bool{true, false}
			ll := &LinkedList[bool]{}

			for _, v := range row {
				el := &Element[bool]{
					value: v,
				}
				ll.Add(el)
			}

			l := ll.List()
			for _, v := range l {
				if !slices.Contains(l, v) {
					t.Errorf("List() does not contain item: %v", v)
				}
			}
			if len(l) != ll.Size() {
				t.Errorf("List() was incorrect, got size: %d, expected size: %d", len(l), ll.Size())
			}
		})
	})

	t.Run("IsEmpty", func(t *testing.T) {
		ll := &LinkedList[int]{}

		if ll.IsEmpty() != true {
			t.Errorf("IsEmpty() was incorrect, got: %t, expected: %t", ll.IsEmpty(), true)
		}
		ll.Add(
			&Element[int]{
				value: 1,
			},
		)
		if ll.IsEmpty() != false {
			t.Errorf("IsEmpty() was incorrect, got: %t, expected: %t", ll.IsEmpty(), false)
		}
		ll.Add(
			&Element[int]{
				value: 2,
			},
		)
		if ll.IsEmpty() != false {
			t.Errorf("IsEmpty() was incorrect, got: %t, expected: %t", ll.IsEmpty(), false)
		}
		_ = ll.Delete(
			&Element[int]{
				value: 1,
			},
		)
		if ll.IsEmpty() != false {
			t.Errorf("IsEmpty() was incorrect, got: %t, expected: %t", ll.IsEmpty(), false)
		}
		_ = ll.Delete(
			&Element[int]{
				value: 2,
			},
		)
		if ll.IsEmpty() != true {
			t.Errorf("IsEmpty() was incorrect, got: %t, expected: %t", ll.IsEmpty(), true)
		}
	})

	t.Run("Size", func(t *testing.T) {
		t.Run("int", func(t *testing.T) {
			row := []int{1, 2, 3}
			ll := &LinkedList[int]{}

			for _, v := range row {
				el := &Element[int]{
					value: v,
				}
				ll.Add(el)
			}

			if ll.Size() != len(row) {
				t.Errorf("Size() was incorrect, got: %d, expected: %d", ll.Size(), len(row))
			}
			for _, v := range row {
				_ = ll.Delete(
					&Element[int]{
						value: v,
					},
				)
			}
			if ll.Size() != 0 {
				t.Errorf("Size() was incorrect, got: %d, expected: %d", ll.Size(), 0)
			}
		})

		t.Run("string", func(t *testing.T) {
			row := []string{"1", "2", "3", "4"}
			ll := &LinkedList[string]{}

			for _, v := range row {
				el := &Element[string]{
					value: v,
				}
				ll.Add(el)
			}

			if ll.Size() != len(row) {
				t.Errorf("Size() was incorrect, got: %d, expected: %d", ll.Size(), len(row))
			}
			for _, v := range row {
				_ = ll.Delete(
					&Element[string]{
						value: v,
					},
				)
			}
			if ll.Size() != 0 {
				t.Errorf("Size() was incorrect, got: %d, expected: %d", ll.Size(), 0)
			}
		})

		t.Run("bool", func(t *testing.T) {
			row := []bool{true, false}
			ll := &LinkedList[bool]{}

			for _, v := range row {
				el := &Element[bool]{
					value: v,
				}
				ll.Add(el)
			}

			if ll.Size() != len(row) {
				t.Errorf("Size() was incorrect, got: %d, expected: %d", ll.Size(), len(row))
			}
			for _, v := range row {
				_ = ll.Delete(
					&Element[bool]{
						value: v,
					},
				)
			}
			if ll.Size() != 0 {
				t.Errorf("Size() was incorrect, got: %d, expected: %d", ll.Size(), 0)
			}
		})
	})
}
