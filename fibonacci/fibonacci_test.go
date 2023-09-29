package main

import (
	"flag"
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
)

func TestVersion(t *testing.T) {
	//store the Go version is at least 1.21
	goversion := runtime.Version()
	expectedversion := "go1.21"

	//test that the Go version is at least 1.21
	if assert.GreaterOrEqual(t, goversion, expectedversion) {
		t.Logf("Go runtime version check PASSED, expected version >= '%s', got '%s'", expectedversion, goversion)
	} else {
		t.Errorf("Go runtime version check FAILED, expected version >= '%s', got '%s'", expectedversion, goversion)
	}
}

func TestSeq(t *testing.T) {
	flag.Parse()
	iter := *iter

	//set the control sequence
	fibonacci := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89}

	//execute the function and report result
	sequence := Seq(iter, 0, 1)

	//test function generates a sequence of appropriate length and report result
	if assert.Len(t, sequence, iter) {
		t.Logf("Seq(%d) PASSED, expected length %d, got %d", iter, iter, len(sequence))
	} else {
		t.Errorf("Seq(%d) FAILED, expected length %d, got %d", iter, iter, len(sequence))
	}

	//test the function generates the Fibonacci sequence and report result
	if assert.Equal(t, sequence[:min(iter, len(fibonacci))], fibonacci[:min(iter, len(fibonacci))]) {
		t.Logf("Seq(%d) PASSED, expected sequence %v, got %v", iter, fibonacci[:min(iter, len(fibonacci))], sequence[:min(iter, len(fibonacci))])
	} else {
		t.Errorf("Seq(%d) FAILED, expected sequence %v, got %v", iter, fibonacci[:min(iter, len(fibonacci))], sequence[:min(iter, len(fibonacci))])
	}
}
