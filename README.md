## Go Module: fibonacci
Program to generate an array containing the Fibonacci sequence.

### Requirements
Tested on Go version 1.21.

External Modules:
```
"flag"
"fmt"
"github.com/stretchr/testify/assert"
"testing"
```

### Execution
```
kpeder@kpeder-devops:~/Projects/go-fibonacci$ go run fibonacci -iter 13
[0 1 1 2 3 5 8 13 21 34 55 89 144]
```

### Run Tests
```
kpeder@kpeder-devops:~/Projects/go-fibonacci$ go test -v -iter 12
=== RUN   TestSeq
    fibonacci_test.go:21: Seq(12) PASSED, expected length 12, got 12
    fibonacci_test.go:28: Seq(12) PASSED, expected sequence [0 1 1 2 3 5 8 13 21 34 55 89], got [0 1 1 2 3 5 8 13 21 34 55 89]
--- PASS: TestSeq (0.00s)
PASS
ok      fibonacci       0.009s
```
