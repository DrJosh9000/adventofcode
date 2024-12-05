package main

import (
	"fmt"

	"drjosh.dev/exp"
)

// Advent of Code 2015
// Day 17

func main() {
	containers := exp.MustReadInts("inputs/17.txt", "\n")

	const L = 150
	combos, combos2 := 0, 0
	// minc := math.MaxInt
	for s := 0; s < (1 << len(containers)); s++ {
		v, c := 0, 0
		for j := 0; j < len(containers); j++ {
			if s&(1<<j) != 0 {
				v += containers[j]
				c++
			}
		}
		if v == L {
			combos++
			// if c < minc {
			// 	minc = c
			// }
			if c == 4 {
				combos2++
			}
		}
	}

	// fmt.Println(minc)
	fmt.Println(combos, combos2)
}
