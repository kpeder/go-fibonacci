package main

import (
	"flag"
	"fmt"
)

var (
	iter *int = flag.Int("iter", 14, "length of the fibonacci sequence to generate")
)

func Seq(iter int) (seq []int) {
	// initialize with seed values
	seq = append(seq, 0, 1)
	// iterate until length iter is reached
	for i := 2; i < iter; i++ {
		seq = append(seq, seq[i-1]+seq[i-2])
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
