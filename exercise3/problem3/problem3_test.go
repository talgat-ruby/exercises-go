package problem3

import (
	"slices"
	"testing"

	"github.com/talgat-ruby/exercises-go/pkg/util"
)

func TestSet(t *testing.T) {
	t.Run("Add", func(t *testing.T) {
		table := []struct {
			vals []any
			size int
		}{
			{[]any{1, 2, 3, 1, 2}, 3},
			{[]any{"1", "2", "3", "4"}, 4},
			{[]any{true, false, true, false}, 2},
		}

		for _, r := range table {
			set := NewSet()
			for _, v := range r.vals {
				set.Add(v)
			}
			if set.Size() != r.size {
				t.Errorf("Add(%v) was incorrect, got len: %d, expected len: %d", r.vals, set.Size(), r.size)
			}
		}
	})

	t.Run("Remove", func(t *testing.T) {
		table := []struct {
			vals []any
		}{
			{[]any{1, 2, 3, 1, 2}},
			{[]any{"1", "2", "3", "4"}},
			{[]any{true, false, true, false}},
		}

		for _, r := range table {
			set := NewSet()

			for _, v := range r.vals {
				set.Add(v)
			}
			for _, v := range r.vals {
				set.Remove(v)
			}
			if set.Size() != 0 {
				t.Errorf("Remove() was incorrect, got size: %d, expected size: %d", set.Size(), 0)
			}
		}
	})

	t.Run("IsEmpty", func(t *testing.T) {
		set := NewSet()

		if set.IsEmpty() != true {
			t.Errorf("IsEmpty() was incorrect, got: %t, expected: %t", set.IsEmpty(), true)
		}
		set.Add(1)
		if set.IsEmpty() != false {
			t.Errorf("IsEmpty() was incorrect, got: %t, expected: %t", set.IsEmpty(), false)
		}
		set.Add(2)
		if set.IsEmpty() != false {
			t.Errorf("IsEmpty() was incorrect, got: %t, expected: %t", set.IsEmpty(), false)
		}
		set.Remove(1)
		if set.IsEmpty() != false {
			t.Errorf("IsEmpty() was incorrect, got: %t, expected: %t", set.IsEmpty(), false)
		}
		set.Remove(2)
		if set.IsEmpty() != true {
			t.Errorf("IsEmpty() was incorrect, got: %t, expected: %t", set.IsEmpty(), true)
		}
	})

	t.Run("Size", func(t *testing.T) {
		table := []struct {
			vals []any
			size int
		}{
			{[]any{1, 2, 3, 1, 2}, 3},
			{[]any{"1", "2", "3", "4"}, 4},
			{[]any{true, false, true, false}, 2},
		}

		for _, r := range table {
			set := NewSet()

			for _, v := range r.vals {
				set.Add(v)
			}
			if set.Size() != r.size {
				t.Errorf("Size() was incorrect, got: %d, expected: %d", set.Size(), r.size)
			}
			for i, v := range r.vals {
				set.Remove(v)
				if i >= r.size {
					break
				}
			}
			if set.Size() != 0 {
				t.Errorf("Size() was incorrect, got: %d, expected: %d", set.Size(), 0)
			}
		}
	})

	t.Run("List", func(t *testing.T) {
		table := []struct {
			vals []any
			size int
		}{
			{[]any{1, 2, 3, 1, 2}, 3},
			{[]any{"1", "2", "3", "4"}, 4},
			{[]any{true, false, true, false}, 2},
		}

		for _, r := range table {
			set := NewSet()

			for _, v := range r.vals {
				set.Add(v)
			}
			l := set.List()
			for _, v := range r.vals {
				if !slices.Contains(l, v) {
					t.Errorf("List() does not contain item: %v", v)
				}
			}
			if len(l) != set.Size() {
				t.Errorf("List() was incorrect, got size: %d, expected size: %d", len(l), set.Size())
			}
		}
	})

	t.Run("Has", func(t *testing.T) {
		table := []struct {
			vals []any
			size int
		}{
			{[]any{1, 2, 3, 1, 2}, 3},
			{[]any{"1", "2", "3", "4"}, 4},
			{[]any{true, false, true, false}, 2},
		}

		for _, r := range table {
			set := NewSet()

			for _, v := range r.vals {
				set.Add(v)
			}
			for _, v := range r.vals {
				if !set.Has(v) {
					t.Errorf("Has(%v) was incorrect, got: %t, expected: %t", v, false, true)
				}
			}
		}
	})

	t.Run("Copy", func(t *testing.T) {
		table := []struct {
			vals []any
			size int
		}{
			{[]any{1, 2, 3, 1, 2}, 3},
			{[]any{"1", "2", "3", "4"}, 4},
			{[]any{true, false, true, false}, 2},
		}

		for _, r := range table {
			set := NewSet()

			for _, v := range r.vals {
				set.Add(v)
			}
			c := set.Copy()

			if &c == &set {
				t.Errorf("Copy() did not copy to a new set, copy address: %p, set address: %d", &c, &set)
			}
			for _, v := range set.List() {
				if !c.Has(v) {
					t.Errorf("Copy() does not contain item: %v", v)
				}
			}
			if c.Size() != set.Size() {
				t.Errorf("Copy() was incorrect, got size: %d, expected size: %d", c.Size(), set.Size())
			}
		}
	})

	t.Run("Difference", func(t *testing.T) {
		util.SkipTestOptional(t)

		table := []struct {
			vals1 []any
			vals2 []any
			exp   []any
		}{
			{[]any{1, 2, 3, 1, 2}, []any{1, 3}, []any{2}},
			{[]any{"1", "2", "3", "4"}, []any{"2", "4"}, []any{"1", "3"}},
			{[]any{true, false, true, false}, []any{}, []any{true, false}},
		}

		for _, r := range table {
			set1 := NewSet()
			for _, v := range r.vals1 {
				set1.Add(v)
			}

			set2 := NewSet()
			for _, v := range r.vals2 {
				set2.Add(v)
			}

			diff := set1.Difference(set2)

			for _, v := range r.exp {
				if !diff.Has(v) {
					t.Errorf("Difference() does not contain item: %v", v)
				}
			}
			if diff.Size() != len(r.exp) {
				t.Errorf("Difference() was incorrect, got size: %d, expected size: %d", diff.Size(), len(r.exp))
			}
		}
	})

	t.Run("IsSubset", func(t *testing.T) {
		util.SkipTestOptional(t)

		table := []struct {
			set    []any
			subset []any
			exp    bool
		}{
			{[]any{1, 2, 3, 1, 2}, []any{1, 3}, true},
			{[]any{"1", "2", "3", "4"}, []any{"0", "2", "4"}, false},
			{[]any{true, false, true, false}, []any{false}, true},
		}

		for _, r := range table {
			set := NewSet()
			for _, v := range r.set {
				set.Add(v)
			}

			subset := NewSet()
			for _, v := range r.subset {
				subset.Add(v)
			}

			isSubset := subset.IsSubset(set)
			if isSubset != r.exp {
				t.Errorf("IsSubset() was incorrect, got: %t, expected: %t", isSubset, r.exp)
			}
		}
	})

	t.Run("Union", func(t *testing.T) {
		util.SkipTestOptional(t)

		table := []struct {
			sets [][]any
			exp  []any
		}{
			{
				[][]any{
					{1, 2, 3},
					{2, 4, 5, 6},
					{9, 10},
				},
				[]any{1, 2, 3, 4, 5, 6, 9, 10},
			},
			{
				[][]any{
					{"1", "2", "3", "4"},
					{1, 2, 3},
				},
				[]any{"1", "2", "3", "4", 1, 2, 3},
			},
			{
				[][]any{
					{true, false, true, false},
					{true, false, true, false},
					{true, false, true, false},
				},
				[]any{true, false},
			},
		}

		for _, r := range table {
			sets := make([]*Set, 0, len(r.sets))
			for _, s := range r.sets {
				newSet := NewSet()
				for _, v := range s {
					newSet.Add(v)
				}
				sets = append(sets, newSet)
			}

			u := Union(sets...)

			for _, v := range r.exp {
				if !u.Has(v) {
					t.Errorf("Union(%v) does not contain item: %v", r.sets, v)
				}
			}
			if u.Size() != len(r.exp) {
				t.Errorf("Union(%v) was incorrect, got size: %d, expected size: %d", r.sets, u.Size(), len(r.exp))
			}
		}
	})

	t.Run("Intersect", func(t *testing.T) {
		util.SkipTestOptional(t)

		table := []struct {
			sets [][]any
			exp  []any
		}{
			{
				[][]any{
					{1, 2, 3, 4, 5},
					{2, 3, 4, 5, 6},
					{3, 4, 5, 6, 7},
				},
				[]any{3, 4, 5},
			},
			{
				[][]any{
					{"1", "2", "3", "4"},
					{"1", "2", "3", "4"},
					{"1", "2", "3", "4"},
				},
				[]any{"1", "2", "3", "4"},
			},
			{
				[][]any{
					{true},
					{true},
					{false},
				},
				[]any{},
			},
		}

		for _, r := range table {
			sets := make([]*Set, 0, len(r.sets))
			for _, s := range r.sets {
				newSet := NewSet()
				for _, v := range s {
					newSet.Add(v)
				}
				sets = append(sets, newSet)
			}

			u := Intersect(sets...)

			for _, v := range r.exp {
				if !u.Has(v) {
					t.Errorf("Intersect(%v) does not contain item: %v", r.sets, v)
				}
			}
			if u.Size() != len(r.exp) {
				t.Errorf("Intersect(%v) was incorrect, got size: %d, expected size: %d", r.sets, u.Size(), len(r.exp))
			}
		}
	})
}
