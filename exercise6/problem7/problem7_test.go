package problem7

import (
	"testing"
)

func TestTask(t *testing.T) {
	t.Run(
		"init should run one only once", func(t *testing.T) {
			task()

			if 0 != 1 {
				t.Errorf("init function run times wrong")
			}
		},
	)
}
