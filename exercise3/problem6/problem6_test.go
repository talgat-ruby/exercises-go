package problem6

import (
	"testing"
)

func TestLegs(t *testing.T) {
	horse := &Animal{
		name:    "horse",
		legsNum: 4,
	}
	kangaroo := &Animal{
		name:    "kangaroo",
		legsNum: 2,
	}
	ant := &Insect{
		name:    "ant",
		legsNum: 6,
	}
	spider := &Insect{
		name:    "spider",
		legsNum: 8,
	}

	totalLegsNum := sumOfAllLegsNum(horse, kangaroo, ant, spider)

	if totalLegsNum != 20 {
		t.Errorf("sumOfAllLegsNum() was incorrect, got: %d, expected: %d", totalLegsNum, 20)
	}
}
