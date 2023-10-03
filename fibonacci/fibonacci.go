package main

import (
	"flag"
	"fmt"
)

var (
	iter *int = flag.Int("iter", 12, "length of the fibonacci sequence to generate")
)

func Seq(iter int) (seq []int) {
	// define local vars
	var (
		i, n int
		x, y int = 0, 1
	)
	// initialize with seed values
	seq = append(seq, x, y)
	// iterate until length iter is reached
	for i = 2; i < iter; i++ {
		n = x + y
		x = y
		y = n
		seq = append(seq, n)
	}
	// return the sequence
	return
}

func main() {
	//get flags
	flag.Parse()

	//get iterations
	iter := *iter

	//print the function's output
	fmt.Println(Seq(iter))
}
