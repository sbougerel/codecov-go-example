[![codecov](https://codecov.io/gh/sbougerel/codecov-go-example/branch/master/graph/badge.svg)](https://codecov.io/gh/sbougerel/codecov-go-example)

# codecov-go-example
By default, there is a discrepancy in computation of coverage by
Codecov.io and the coverage reported by `go test -covermode=atomic`.

This is because of the way `go test ...` reports coverage; for `if` statements:
```go
   if Condition() {
     DoSomething()
   }
```

it includes the trailing `{` at end of line containig the `if Condition()` into
the _next_ statement `DoSomething()`. This means that if `Condition()` evaluates
to false, you have a _partial hit_ on this line. Additionally, `go test ...`
still keeps track of statements covered, regardless of the lines hit, while
Codecov tracks lines.

__The output of both tools can differ widely due to that__. At my company, which
uses a large codebase, we have more that 10% in coverage difference reported by
both tools on some packages.

I used this repository to experiment with different solutions.

## Obtaining coverage with GO tools

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

## Prefered way to push coverage to Codecov.

After exprimentation, I found that using CoberturaXML format yields the best
match to runing `go test ...` and most importantly, shows on Codecov what one
would expect of the coverage of its files.

Get external utilities:

```
$ go get github.com/axw/gocov/...
$ go get github.com/AlekSi/gocov-xml
```

After the output of `go test ...` above, run:

```
$ gocov convert coverage.txt | gocov-xml > coverage.xml
```

Then push this new XML file to Codecov.

```
bash <(curl -s https://codecov.io/bash) -f coverage.xml
```

See [coverage results on Codecov.io](https://codecov.io/gh/sbougerel/codecov-go-example).

## Results

As can be seen above, the GO tool reports that 50% of statements are covered,
and same for Codecov.io. For more visual evidance that the coverage is reported
are similar, you can run, on the original `coverage.txt`

```
go tool cover -html coverage.txt
```

# Conclusion

This method is the one currently in use in my company.
