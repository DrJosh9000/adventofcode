package main

import (
	"fmt"

	"github.com/DrJosh9000/exp"
)

// Advent of Code 2022
// Day 4, part a

func main() {
	count := 0
	for _, line := range exp.MustReadLines("inputs/4.txt") {
		var a, b, c, d int
		exp.Must(fmt.Sscanf(line, "%d-%d,%d-%d", &a, &b, &c, &d))
		if (a <= c && b >= d) || (c <= a && d >= b) {
			count++
		}
	}
	fmt.Println(count)
}
