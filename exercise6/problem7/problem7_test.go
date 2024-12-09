package problem7

import (
	"testing"
)

func TestTask(t *testing.T) {
	t.Run(
		"should avoid data race", func(t *testing.T) {
			task()
		},
	)
}
