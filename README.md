## Go Module: fibonacci
Program to generate an array containing the Fibonacci sequence,
and to demonstrate test patterns in Go.

### Requirements
Tested on Go version 1.21.

External Modules:
```
"flag"
"fmt"
"github.com/stretchr/testify/assert"
"testing"
```

### Compilation
```
user@host:~/Projects/go-fibonacci/fibonacci$ go mod init fibonacci
go: creating new go.mod: module fibonacci
go: to add module requirements and sums:
        go mod tidy
user@host:~/Projects/go-fibonacci/fibonacci$ go mod tidy
go: finding module for package github.com/stretchr/testify/assert
go: found github.com/stretchr/testify/assert in github.com/stretchr/testify v1.8.4
```

### Execution
```
user@host:~:~/Projects/go-fibonacci/fibonacci$ go run fibonacci -iter 13
[0 1 1 2 3 5 8 13 21 34 55 89 144]
```

### Run Tests
```
user@host:~/Projects/go-fibonacci/fibonacci$ go test -v -iter 12
=== RUN   TestSeq
    fibonacci_test.go:21: Seq(12) PASSED, expected length 12, got 12
    fibonacci_test.go:28: Seq(12) PASSED, expected sequence [0 1 1 2 3 5 8 13 21 34 55 89], got [0 1 1 2 3 5 8 13 21 34 55 89]
--- PASS: TestSeq (0.00s)
PASS
ok      fibonacci       0.009s
```
