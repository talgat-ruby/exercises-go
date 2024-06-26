package util

import (
	"os"
	"testing"
)

func SkipTestOptional(t *testing.T) {
	if os.Getenv("SKIPTEST") == "optional" {
		t.Skip("Skipping testing in CI environment")
	}
}
