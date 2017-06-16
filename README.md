# codecov-go-example
This GO repo only exists to illustrate discrepancy in computation of coverage by
Codecov.io and the coverage reported by `go test -covermode=atomic`.

Specifically, this repository contains only a single GO file, with a single test file.
This test file is made specifically to ignore some conditional path of the GO file.

The fact that some lines are partially covered leads to computation differences between
the 2 applications. This means that relying on one tool will confuse the user of the
other tools.

## Obtaining coverage

```
$ go test -race -covermode=atomic -coverprofile=coverage.txt .
ok  	github.com/sbougerel/codecov-go-example	1.016s	coverage: 50.0% of statements
```

The generated coverage file contents are:

```
mode: atomic
github.com/sbougerel/codecov-go-example/main.go:5.44,6.8 1 1
github.com/sbougerel/codecov-go-example/main.go:10.2,10.8 1 1
github.com/sbougerel/codecov-go-example/main.go:14.2,14.13 1 0
github.com/sbougerel/codecov-go-example/main.go:6.8,8.3 1 0
github.com/sbougerel/codecov-go-example/main.go:10.8,12.3 1 1
github.com/sbougerel/codecov-go-example/main.go:17.13,19.2 1 0
```

## Coverage % computation for Codecov.io is 37.5%

See [coverage results on Codecov.io](https://codecov.io/gh/sbougerel/codecov-go-example).

Codecov.io looks at the same generated file `coverage.txt` as the GO tool but reports a much lower coverage.

## Coverage % computation for Go tool is 50%

As can be seen above, the GO tool reports that 50% of statements are covered.
For more visual evidance that the coverage is reported differently:

```
go tool cover -html coverage.txt
```
