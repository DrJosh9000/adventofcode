package main

import (
	"fmt"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

// Advent of Code 2024
// Day 22, part b

const inputPath = "2024/inputs/22.txt"

func main() {
	patts := make(map[[4]int8]int)
	for _, seed := range exp.MustReadInts(inputPath, "\n") {
		seen := make(algo.Set[[4]int8])
		var patt [4]int8
		n := seed
		prev := n % 10
		for i := range 2000 {
			n = round(n)
			curr := n % 10
			patt[0], patt[1], patt[2], patt[3] = patt[1], patt[2], patt[3], int8(curr-prev)
			if i >= 3 && !seen.Contains(patt) {
				seen.Insert(patt)
				patts[patt] += curr
				// fmt.Printf("%v: %d\n", patt, curr)
			}
			prev = curr
		}
	}

	best := 0
	for _, b := range patts {
		best = max(best, b)
	}
	fmt.Println(best)
}

func round(n int) int {
	n ^= n << 6
	n &= 0xFFFFFF
	n ^= n >> 5
	n &= 0xFFFFFF
	n ^= n << 11
	n &= 0xFFFFFF
	return n
}
