package main

import (
	"fmt"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

// Advent of Code 2016
// Day 15, part a

func main() {
	// Position of disc n at time t:
	//   pos(n, t) = (t + pos(n, 0)) % size(n)
	// Want to find t such that for all discs:
	//   pos(n, t+n) ≡ 0  (mod size(n))
	// i.e.
	//   (t + n + pos(n, 0)) ≡ 0  (mod size(n))
	// or:
	//   t ≡ -n - pos(n, 0)  (mod size(n))
	// Using CRT, just because.

	sol := 0
	prod := 1
	for _, line := range exp.MustReadLines("inputs/15.txt") {
		var n, size, pos int
		exp.Must(fmt.Sscanf(line, "Disc #%d has %d positions; at time=0, it is at position %d.", &n, &size, &pos))
		if prod == 1 {
			sol = size - n - pos
			prod = size
			//log.Printf("Solution at disc %d (size %d): %d", n, size, sol)
			continue
		}
		_, x, y := algo.XGCD(prod, size)
		//log.Printf("Bézout's identity: %d*%d + %d*%d == %d", prod, x, size, y, prod*x+size*y)
		sol = sol*y*size + (-n-pos)*x*prod
		prod *= size
		for sol < 0 {
			sol += prod
		}
		sol %= prod
		//log.Printf("Solution at disc %d (size %d): %d", n, size, sol)
	}
	fmt.Println(sol)
}
