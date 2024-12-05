package main

import (
	"fmt"

	"drjosh.dev/exp"
)

// Advent of Code 2022
// Day 4, part b

func main() {
	count := 0
	for _, line := range exp.MustReadLines("inputs/4.txt") {
		var a, b, c, d int
		exp.Must(fmt.Sscanf(line, "%d-%d,%d-%d", &a, &b, &c, &d))
		if (a >= c && a <= d) || (b >= c && b <= d) || (c >= a && c <= b) || (d >= a && d <= b) {
			count++
		}
	}
	fmt.Println(count)
}
