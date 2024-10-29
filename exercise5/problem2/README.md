# Problem 2

`addConcurrently` needs to be implemented concurrently:

1. Utilize all cores on machine (Advanced concept)
2. Divide the input into parts
3. Run computation for each part in separate goroutine.

This is an optional problem which involves `runtime` package and includes benchmark tests.
Because of nature of benchmark tests, I can only share my results:

```shell
$ go test -v -bench=. -race ./exercise4/problem2/...
=== RUN   TestAdd
--- PASS: TestAdd (0.10s)
goos: darwin
goarch: arm64
pkg: github.com/talgat-ruby/exercises-go/exercise4/problem2
BenchmarkAdd
BenchmarkAdd-8             	      20	  55690504 ns/op
BenchmarkAddConcurrent
BenchmarkAddConcurrent-8   	      45	  25124303 ns/op
PASS
ok  	github.com/talgat-ruby/exercises-go/exercise4/problem2	6.019s
```

