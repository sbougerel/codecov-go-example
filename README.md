# codecov-go-example
This GO repo only exists to illustrate discrepancy in computation of coverage by
Codecov.io and the coverage reported by `go test -covermode=atomic`.

Specifically, this repository contains only a single GO file, with a single test file.
This test file is made specifically to ignore some conditional path of the GO file.

The fact that some lines are partially covered leads to computation differences between
the 2 applications. This means that relying on one tool will confuse the user of the
other tools.

## Coverage % computation for Codecov.io

See [https://codecov.io/gh/sbougerel/codecov-go-example](on Codecov.io)

## Coverage % computation for Go tool

```
$ go test -race -covermode=atomic -coverprofile=coverage.txt .
ok  	github.com/sbougerel/codecov-go-example	1.016s	coverage: 50.0% of statements
```

When runing with `go test -covermode=atomic -coverprofile=coverage.txt`, the output
generated contains the number of statements being covered at each line.

