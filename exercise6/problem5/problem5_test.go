package problem5

import (
	"slices"
	"sort"
	"testing"
	"time"
)

func notifiedWorkers(notifier <-chan int) []int {
	workersId := make([]int, 0, 1)

	for {
		select {
		case workerId := <-notifier:
			workersId = append(workersId, workerId)
		case <-time.After(time.Second):
			return workersId
		}
	}
}

func TestShoppingListAllWorkers(t *testing.T) {
	t.Run(
		"only one worker should notify", func(t *testing.T) {
			exp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

			shoppingList := &[]string{}
			numWorkers := 10
			notifier := notifyOnShopListUpdate(shoppingList, numWorkers)
			out := notifiedWorkers(notifier)
			sort.Ints(out)

			if !slices.Equal(out, exp) {
				t.Errorf("notification channel has wrong job, expected: %v, got: %v.", exp, out)
			}
			if len(*shoppingList) == 0 {
				t.Error("shopping list is empty, you need to wait until it is completed")
			}
		},
	)
}
