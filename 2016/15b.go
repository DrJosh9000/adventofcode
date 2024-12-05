package main

import (
	"fmt"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

// Advent of Code 2016
// Day 15, part b

func main() {
	// See 15a for discussion.
	sol := 0
	prod := 1
	last := 0
	update := func(n, size, pos int) {
		last = n
		_, x, y := algo.XGCD(prod, size)
		sol = sol*y*size + (-n-pos)*x*prod
		prod *= size
		for sol < 0 {
			sol += prod
		}
		sol %= prod
	}

	for _, line := range exp.MustReadLines("inputs/15.txt") {
		var n, size, pos int
		exp.Must(fmt.Sscanf(line, "Disc #%d has %d positions; at time=0, it is at position %d.", &n, &size, &pos))
		if prod == 1 {
			sol = size - n - pos
			prod = size
			continue
		}
		update(n, size, pos)
	}

	update(last+1, 11, 0)
	fmt.Println(sol)
}
