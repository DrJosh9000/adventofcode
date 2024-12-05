package main

import (
	"fmt"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

// Advent of Code 2015
// Day 2, part a

func main() {
	sum := 0
	for _, line := range exp.MustReadLines("inputs/2.txt") {
		var l, w, h int
		exp.Must(fmt.Sscanf(line, "%dx%dx%d", &l, &w, &h))
		lw, wh, lh := l*w, w*h, l*h
		sum += 2*(lw+wh+lh) + algo.Min(lw, wh, lh)
	}
	fmt.Println(sum)
}
