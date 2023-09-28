package main

import (
	"flag"
	"fmt"
)

var (
	iter *int = flag.Int("iter", 12, "length of the fibonacci sequence to generate")
	x, y int  = 0, 1
)

func Seq(iter int, x int, y int) (seq []int) {

	var (
		i, n int
	)

	seq = append(seq, x, y)

	for i = 2; i < iter; i++ {
		n = x + y
		x = y
		y = n
		seq = append(seq, n)
	}

	// return a fibonacci sequence of length iter
	return
}

func main() {
	//get flags
	flag.Parse()

	//get iterations
	i := *iter

	//print the function's output
	fmt.Println(Seq(i, x, y))
}
