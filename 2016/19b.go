package main

import (
	"fmt"
	"os"
	"strconv"

	"drjosh.dev/exp"
)

// Advent of Code 2016
// Day 19, part b

func main() {
	n := exp.Must(strconv.Atoi(os.Args[1]))

	// Solution: generate a run of answers for small inputs with a full
	// simulation, and assume the pattern holds for all inputs.

	p, step := 0, 1
	for i := 1; ; i++ {
		p += step
		if i == n {
			fmt.Println(p)
			return
		}
		switch i {
		case p:
			step = 1
			p = 0
		case p * 2:
			step = 2
		}
	}

}
