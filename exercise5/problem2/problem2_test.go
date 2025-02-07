package problem2

import (
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/talgat-ruby/exercises-go/pkg/util"
)

var numbers []int

func TestMain(m *testing.M) {
	numbers = generateNumbers(1e7)
	code := m.Run()
	numbers = []int{}
	os.Exit(code)
}

func generateNumbers(max int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	nums := make([]int, max)
	for i := range nums {
		nums[i] = r.Intn(10)
	}
	return nums
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(numbers)
	}
}

func BenchmarkAddConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addConcurrently(numbers)
	}
}

func TestAdd(t *testing.T) {
	util.SkipTestOptional(t)

	inp := numbers
	exp := add(inp)
	out := addConcurrently(inp)
	if out != exp {
		t.Errorf("addConcurrently() was incorrect, got: %d, expected: %d.", out, exp)
	}
}
