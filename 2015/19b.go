package main

import (
	"fmt"
	"strings"

	"drjosh.dev/exp"
)

// Advent of Code 2015
// Day 19, part b

// ...that was anticlimactic.

type rule struct {
	in  string
	out string
}

func main() {
	var rules []rule
	var x string
	for _, line := range exp.MustReadLines("inputs/19.txt") {
		if line == "" {
			continue
		}
		in, out, ok := strings.Cut(line, " => ")
		if !ok {
			x = line
			continue
		}
		rules = append(rules, rule{in, out})
	}

	count := 0
	for x != "e" {
		for _, r := range rules {
			for strings.Contains(x, r.out) {
				x = strings.Replace(x, r.out, r.in, 1)
				count++
			}
		}
	}
	fmt.Println(count)
}
