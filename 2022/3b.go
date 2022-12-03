package main

import (
	"fmt"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

// Advent of Code 2022
// Day 3, part b

func main() {
	sum := 0
	input := exp.MustReadLines("inputs/3.txt")
	for gi := 0; gi < len(input); gi += 3 {
		g := input[gi:][:3]
		s := make([]algo.Set[rune], 3)
		for i, line := range g {
			s[i] = s[i].Insert([]rune(line)...)
		}

		for c := range algo.Intersection(s...) {
			if c > 'a' {
				sum += int(c) - 'a' + 1
			} else {
				sum += int(c) - 'A' + 27
			}
		}
	}
	fmt.Println(sum)
}
