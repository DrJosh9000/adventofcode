package main

import (
	"fmt"
	"math/bits"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

// Advent of Code 2015
// Day 24, part a

func main() {
	presents := exp.MustReadInts("inputs/24.txt", "\n")
	N := len(presents)
	target := algo.Sum(presents) / 3

	// The remaining items always partition evenly?

	minitems := N
	minqe := algo.Prod(presents)
	for s := 0; s < (1 << N); s++ {
		n := bits.OnesCount(uint(s))
		if n > minitems {
			continue
		}
		sum, prod := 0, 1
		for i := 0; i < N; i++ {
			if s&(1<<i) != 0 {
				sum += presents[i]
				prod *= presents[i]
			}
		}
		if sum == target {
			if n < minitems {
				minitems = n
				minqe = prod
			}
			if n == minitems && prod < minqe {
				minqe = prod
			}
		}
	}

	fmt.Println(minqe)
}
