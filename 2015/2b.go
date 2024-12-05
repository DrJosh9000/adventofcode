package main

import (
	"fmt"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

// Advent of Code 2015
// Day 2, part b

func main() {
	sum := 0
	for _, line := range exp.MustReadLines("inputs/2.txt") {
		d := make([]int, 3)
		exp.Must(fmt.Sscanf(line, "%dx%dx%d", &d[0], &d[1], &d[2]))
		algo.SortAsc(d)
		sum += 2*(d[0]+d[1]) + algo.Prod(d)
	}
	fmt.Println(sum)
}
