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
"runtime"
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
user@host:~/Projects/go-fibonacci/fibonacci$ mkdir ../bin
user@host:~/Projects/go-fibonacci/fibonacci$ go build -o ../bin .
```

### Execution
```
user@host:~:~/Projects/go-fibonacci/fibonacci$ go run fibonacci -iter 13
[0 1 1 2 3 5 8 13 21 34 55 89 144]
user@host:~:~/Projects/go-fibonacci/fibonacci$ cd ../bin
user@host:~:~/Projects/go-fibonacci/bin$ ./fibonacci -iter 9
[0 1 1 2 3 5 8 13 21]
```

### Run Tests
```
user@host:~/Projects/go-fibonacci/fibonacci$ go test fibonacci -v -iter 12
=== RUN   TestVersion
    fibonacci_test.go:17: Go runtime version check PASSED, expected version >= 'go1.21', got 'go1.21.1'
--- PASS: TestVersion (0.00s)
=== RUN   TestSeq
    fibonacci_test.go:35: Seq(12) PASSED, expected length 12, got 12
    fibonacci_test.go:42: Seq(12) PASSED, expected sequence [0 1 1 2 3 5 8 13 21 34 55 89], got [0 1 1 2 3 5 8 13 21 34 55 89]
--- PASS: TestSeq (0.00s)
PASS
ok      fibonacci       0.002s
```

### Make Targets
```
user@host:~/Projects/go-fibonacci$ make help
make <target>

Targets:

    help    Show this help

    build   Build the fibonacci package
    clean   Clean up build files
    init    Initialize fibonacci module
    test    Run the fibonacci package tests
```
