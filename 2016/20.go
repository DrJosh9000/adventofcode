package main

import (
	"fmt"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

// Advent of Code 2016
// Day 20

func main() {
	type rainj struct { // "range" is a keyword
		a, b uint
	}

	allowed := algo.Set[rainj]{
		{0, 0xffffffff}: {},
	}
	for _, line := range exp.MustReadLines("inputs/20.txt") {
		var r rainj
		exp.Must(fmt.Sscanf(line, "%d-%d", &r.a, &r.b))

		for a := range allowed {
			if a.a > r.b || a.b < r.a {
				continue
			}
			delete(allowed, a)
			if r.a > a.a {
				allowed.Insert(rainj{a.a, r.a - 1})
			}
			if r.b < a.b {
				allowed.Insert(rainj{r.b + 1, a.b})
			}
		}
	}

	min := uint(0xffffffff)
	sum := uint(0)
	for a := range allowed {
		if a.a < min {
			min = a.a
		}
		sum += a.b - a.a + 1
	}
	fmt.Println(min)
	fmt.Println(sum)
}
