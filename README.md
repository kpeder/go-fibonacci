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
kpeder@kpeder-dev:~/Projects/devops-examples/golang/fibonacci$ go run fibonacci -iter 13
[0 1 1 2 3 5 8 13 21 34 55 89 144]
```

### Run Tests
```
kpeder@kpeder-dev:~/Projects/devops-examples/golang/fibonacci$ go test -iter 13 -v
=== RUN   TestFibonacciSeq
    fibonacci_test.go:15: [0 1 1 2 3 5 8 13 21 34 55 89 144]
--- PASS: TestFibonacciSeq (0.00s)
PASS
ok      fibonacci       0.002s
```
